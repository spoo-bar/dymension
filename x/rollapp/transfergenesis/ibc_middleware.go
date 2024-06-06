package transfergenesis

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dymensionxyz/dymension/v3/utils/gerr"

	"github.com/dymensionxyz/dymension/v3/utils"

	delayedacktypes "github.com/dymensionxyz/dymension/v3/x/delayedack/types"

	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	dymerror "github.com/dymensionxyz/dymension/v3/x/common/errors"
	delayedackkeeper "github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	rollappkeeper "github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"
)

var _ porttypes.Middleware = &IBCMiddleware{}

type DenomMetadataKeeper interface {
	CreateDenomMetadata(ctx sdk.Context, metadata banktypes.Metadata) error
}

type TransferKeeper interface {
	SetDenomTrace(ctx sdk.Context, denomTrace transfertypes.DenomTrace)
}

type IBCMiddleware struct {
	porttypes.Middleware // next one
	delayedackKeeper     delayedackkeeper.Keeper
	rollappKeeper        rollappkeeper.Keeper
	transferKeeper       TransferKeeper
	denomKeeper          DenomMetadataKeeper
}

func NewIBCMiddleware(
	next porttypes.Middleware,
	delayedAckKeeper delayedackkeeper.Keeper,
	rollappKeeper rollappkeeper.Keeper,
	transferKeeper TransferKeeper,
	denomKeeper DenomMetadataKeeper,
) IBCMiddleware {
	return IBCMiddleware{
		Middleware:       next,
		delayedackKeeper: delayedAckKeeper,
		rollappKeeper:    rollappKeeper,
		transferKeeper:   transferKeeper,
		denomKeeper:      denomKeeper,
	}
}

type memo struct {
	Denom banktypes.Metadata `json:"denom"`
	// How many transfers in total will be sent in the transfer genesis period
	TotalNumTransfers uint64 `json:"total_num_transfers"`
	// Which transfer is this? If there are 5 transfers total, they will be numbered 0,1,2,3,4.
	ThisTransferIx uint64 `json:"this_transfer_ix"`
}

// OnRecvPacket will, if the packet is a transfer packet:
// if it's not a genesis transfer: pass on the packet only if transfers are enabled
// else: check it's a valid genesis transfer. If it is, then register the denom, if
// it's the last one, trigger the finalization period to begin.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	if !im.delayedackKeeper.IsRollappsEnabled(ctx) || !isTransferPacket(packet) {
		return im.Middleware.OnRecvPacket(ctx, packet, relayer)
	}

	l := ctx.Logger().With(
		"middleware", "transferGenesis",
		"packet_source_port", packet.SourcePort,
		"packet_destination_port", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	chaID, raID, err := im.getChannelAndRollappID(ctx, packet)
	if err != nil {
		// TODO:
	}

	ra, ok := im.rollappKeeper.GetRollapp(ctx, raID) // TODO: necessary?
	if !ok {
		// TODO:
	}

	m, err := getMemo(packet)
	if errorsmod.IsOf(err, gerr.ErrNotFound) {
		// This is a normal transfer
		if !ra.GenesisState.TransfersEnabled {
			err = errorsmod.Wrapf(gerr.ErrFailedPrecondition, "transfers are disabled: rollapp id: %s", ra.RollappId)
			return channeltypes.NewErrorAcknowledgement(err) // TODO: check with Omri
		}
		return im.Middleware.OnRecvPacket(ctx, packet, relayer)
	}
	if err != nil {
		// TODO:
	}

	nTransfersDone, err := im.rollappKeeper.VerifyAndRecordGenesisTransfer(ctx, raID, m.ThisTransferIx, m.TotalNumTransfers)
	if errorsmod.IsOf(err, dymerror.ErrProtocolViolation) {
		// TODO: emit event or freeze rollapp, or something else?
	}
	if err != nil {
		// TODO:
	}

	// it's a valid genesis transfer

	err = im.registerDenomMetadata(ctx, raID, chaID, m.Denom)
	if err != nil {
		// TODO:
	}

	l.Info("Registered denom meta data from genesis transfer.")

	if nTransfersDone == m.TotalNumTransfers {
		// The transfer window is finished!
		im.rollappKeeper.AddTransferGenesisFinalization(ctx, raID)
		// TODO: emit event
	}

	return im.Middleware.OnRecvPacket(delayedacktypes.SkipContext(ctx), packet, relayer)
}

func isTransferPacket(packet channeltypes.Packet) bool {
	var data transfertypes.FungibleTokenPacketData
	return transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data) == nil
}

func getMemo(packet channeltypes.Packet) (memo, error) {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return memo{}, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, "fungible token packet")
	}

	type t struct {
		Data memo `json:"genesis_transfer"`
	}

	rawMemo := data.GetMemo()
	var m t
	err := json.Unmarshal([]byte(rawMemo), &m)
	if err != nil {
		return memo{}, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, "rawMemo")
	}
	return m.Data, nil
}

func (im IBCMiddleware) getChannelAndRollappID(ctx sdk.Context, packet channeltypes.Packet,
) (string, string, error) {
	chaID := "channel-0"
	raID := "rollappevm_1234-1"
	return chaID, raID, nil
}

func (im IBCMiddleware) ensureRollappExists(ctx sdk.Context, raID string) error {
	ra, ok := im.rollappKeeper.GetRollapp(ctx, raID) // TODO: necessary?
	if !ok {
		panic(errors.New("must find rollapp"))
	}

	_ = ra
	// TODO:
	return nil
}

func (im IBCMiddleware) registerDenomMetadata(ctx sdk.Context, rollappID, channelID string, m banktypes.Metadata) error {
	// TODO: only do it if it hasn't been done before?

	trace := utils.GetForeignDenomTrace(channelID, m.Base)

	im.transferKeeper.SetDenomTrace(ctx, trace)

	ibcDenom := trace.IBCDenom()

	/*
		Change the base to the ibc denom, and add an alias to the original
	*/
	m.Description = fmt.Sprintf("auto-generated ibc denom for rollapp: base: %s: rollapp: %s", ibcDenom, rollappID)
	m.Base = ibcDenom
	for i, u := range m.DenomUnits {
		if u.Exponent == 0 {
			m.DenomUnits[i].Aliases = append(m.DenomUnits[i].Aliases, u.Denom)
			m.DenomUnits[i].Denom = ibcDenom
		}
	}

	if err := m.Validate(); err != nil {
		// TODO: errorsmod with nice wrapping
		return fmt.Errorf("invalid denom metadata on genesis event: %w", err)
	}

	// We go by the denom keeper instead of calling bank directly, as something might happen in-between
	err := im.denomKeeper.CreateDenomMetadata(ctx, m)
	if errorsmod.IsOf(err, gerr.ErrAlreadyExist) {
		return nil
	}
	if err != nil {
		return errorsmod.Wrap(err, "create denom metadata")
	}

	return nil
}
