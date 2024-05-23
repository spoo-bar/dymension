package delayedack

import (
	"errors"
	"fmt"

	rollappkeeper "github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	commontypes "github.com/dymensionxyz/dymension/v3/x/common/types"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks
type IBCMiddleware struct {
	porttypes.IBCModule
	keeper   keeper.Keeper
	raKeeper rollappkeeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, keeper keeper.Keeper, raK rollappkeeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		IBCModule: app,
		keeper:    keeper,
		raKeeper:  raK,
	}
}

// OnRecvPacket handles the receipt of a packet and puts it into a pending queue
// until its state is finalized
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	if !im.keeper.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source_port", packet.SourcePort,
		"packet_destination_port", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	rollappPortOnHub, rollappChannelOnHub := packet.DestinationPort, packet.DestinationChannel

	rollappID, transferPacketData, err := im.ExtractRollappIDAndTransferPacket(ctx, packet, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Extract rollapp id from packet.", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if transferPacketData.GetMemo() == "special" {
		/*
			What are the steps that need to happen?
			1. Make sure that this packet actually originated from the rollapp transfer channel (how)
			2. Mark the genesis as occurred
			3. Pass on the ibc packet
			(TODO: Order between 2-3)
			(TODO: impact of one failure, how to properly error ack?)
		*/

		logger.Info("got the special memo!")
		p := rollapptypes.TriggerGenesisArgs{
			ChannelID: "channel-0",
			RollappID: "rollappevm_1234-1",
		}

		ra, ok := im.raKeeper.GetRollapp(ctx, p.RollappID)
		if !ok {
			panic(errors.New("expect to find rollapp"))
		}
		metadata := ra.TokenMetadata
		accs := ra.GenesisState.GenesisAccounts
		_ = ra
		_ = metadata
		_ = accs
		err = im.raKeeper.MarkGenesisAsHappened(ctx, p)
		if err != nil {
			err = fmt.Errorf("trigger genesis func: %w", err)
			logger.Error("OnRecvPacket", "err", err)
			panic(err)
		}
		logger.Info("Triggered genesis func due to special memo, now passing on packet.")
		ack := im.IBCModule.OnRecvPacket(ctx, packet, relayer)
		if !ack.Success() {
			logger.Error("OnRecvPacket, forwarded special tx, but got an err ack", "ack", string(ack.Acknowledgement()))
		}
		return ack
	}

	if rollappID == "" {
		logger.Debug("Skipping eIBC transfer OnRecvPacket for non-rollapp chain") // TODO: this should say eIBC right?
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	err = im.keeper.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_RECV, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if finalized {
		logger.Debug("Skipping eIBC transfer OnRecvPacket as the packet proof height is already finalized")
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:   rollappID,
		Packet:      &packet,
		Status:      commontypes.Status_PENDING,
		Relayer:     relayer,
		ProofHeight: proofHeight,
		Type:        commontypes.RollappPacket_ON_RECV,
	}
	im.keeper.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return nil
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	if !im.keeper.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source", packet.SourcePort,
		"packet_destination", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	rollappPortOnHub, rollappChannelOnHub := packet.SourcePort, packet.SourceChannel

	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		logger.Error("Unmarshal acknowledgement", "err", err)
		return errorsmod.Wrapf(types.ErrUnknownRequest, "unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	rollappID, transferPacketData, err := im.ExtractRollappIDAndTransferPacket(ctx, packet, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to extract rollapp id from channel", "err", err)
		return err
	}

	if rollappID == "" {
		logger.Debug("Skipping IBC transfer OnAcknowledgementPacket for non-rollapp chain")
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}
	err = im.keeper.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return err
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_ACK, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return err
	}

	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return err
	}

	if finalized {
		logger.Debug("Skipping eIBC transfer OnAcknowledgementPacket as the packet proof height is already finalized")
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}
	// Run the underlying app's OnAcknowledgementPacket callback
	// with cache context to avoid state changes and report the acknowledgement result.
	// Only save the packet if the underlying app's callback succeeds.
	cacheCtx, _ := ctx.CacheContext()
	err = im.IBCModule.OnAcknowledgementPacket(cacheCtx, packet, acknowledgement, relayer)
	if err != nil {
		return err
	}
	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:       rollappID,
		Packet:          &packet,
		Acknowledgement: acknowledgement,
		Status:          commontypes.Status_PENDING,
		Relayer:         relayer,
		ProofHeight:     proofHeight,
		Type:            commontypes.RollappPacket_ON_ACK,
	}
	im.keeper.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	switch ack.Response.(type) {
	// Only if the acknowledgement is an error, we want to create an order
	case *channeltypes.Acknowledgement_Error:
		err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
		if err != nil {
			return err
		}
	}

	return nil
}

// OnTimeoutPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	if !im.keeper.IsRollappsEnabled(ctx) {
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}
	logger := ctx.Logger().With(
		"module", types.ModuleName,
		"packet_source", packet.SourcePort,
		"packet_destination", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	rollappPortOnHub, rollappChannelOnHub := packet.SourcePort, packet.SourceChannel

	rollappID, transferPacketData, err := im.ExtractRollappIDAndTransferPacket(ctx, packet, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to extract rollapp id from channel", "err", err)
		return err
	}

	if rollappID == "" {
		logger.Debug("Skipping IBC transfer OnTimeoutPacket for non-rollapp chain")
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	err = im.keeper.ValidateRollappId(ctx, rollappID, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		logger.Error("Failed to validate rollappID", "rollappID", rollappID, "err", err)
		return err
	}

	proofHeight, err := im.GetProofHeight(ctx, commontypes.RollappPacket_ON_TIMEOUT, rollappPortOnHub, rollappChannelOnHub, packet.Sequence)
	if err != nil {
		logger.Error("Failed to get proof height from packet", "err", err)
		return err
	}
	finalized, err := im.CheckIfFinalized(ctx, rollappID, proofHeight)
	if err != nil {
		logger.Error("Failed to check if packet is finalized", "err", err)
		return err
	}

	if finalized {
		logger.Debug("Skipping IBC transfer OnTimeoutPacket as the packet proof height is already finalized")
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	// Run the underlying app's OnTimeoutPacket callback
	// with cache context to avoid state changes and report the timeout result.
	// Only save the packet if the underlying app's callback succeeds.
	cacheCtx, _ := ctx.CacheContext()
	err = im.IBCModule.OnTimeoutPacket(cacheCtx, packet, relayer)
	if err != nil {
		return err
	}
	// Save the packet data to the store for later processing
	rollappPacket := commontypes.RollappPacket{
		RollappId:   rollappID,
		Packet:      &packet,
		Status:      commontypes.Status_PENDING,
		Relayer:     relayer,
		ProofHeight: proofHeight,
		Type:        commontypes.RollappPacket_ON_TIMEOUT,
	}
	im.keeper.SetRollappPacket(ctx, rollappPacket)

	logger.Debug("Set rollapp packet",
		"rollappID", rollappPacket.RollappId,
		"proofHeight", rollappPacket.ProofHeight,
		"type", rollappPacket.Type)

	err = im.eIBCDemandOrderHandler(ctx, rollappPacket, *transferPacketData)
	if err != nil {
		return err
	}

	return nil
}

/* ------------------------------- ICS4Wrapper ------------------------------ */

// SendPacket implements the ICS4 Wrapper interface
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return im.keeper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

// GetAppVersion returns the application version of the underlying application
func (im IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.keeper.GetAppVersion(ctx, portID, channelID)
}

// ExtractRollappIDAndTransferPacket extracts the rollapp ID from the packet
func (im IBCMiddleware) ExtractRollappIDAndTransferPacket(ctx sdk.Context, packet channeltypes.Packet, rollappPortOnHub string, rollappChannelOnHub string) (string, *transfertypes.FungibleTokenPacketData, error) {
	// no-op if the packet is not a fungible token packet
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return "", nil, fmt.Errorf("unmarshal fungible token packet data: %w", err)
	}
	// Check if the packet is destined for a rollapp
	chainID, err := im.keeper.ExtractChainIDFromChannel(ctx, rollappPortOnHub, rollappChannelOnHub)
	if err != nil {
		return "", &data, fmt.Errorf("extract chain id from channel: %w", err)
	}
	rollapp, found := im.keeper.GetRollapp(ctx, chainID)
	if !found {
		return "", &data, nil // TODO: needs an error?
	}
	_ = rollapp
	/*
		TODO: reimpl semantics somewhere else
		if rollapp.ChannelId == "" {
			return "", &data, errorsmod.Wrapf(rollapptypes.ErrGenesisEventNotTriggered, "empty channel id: rollap id: %s", chainID)
		}
		// check if the channelID matches the rollappID's channelID
		if rollapp.ChannelId != rollappChannelOnHub {
			return "", &data, errorsmod.Wrapf(
				rollapptypes.ErrMismatchedChannelID,
				"channel id mismatch: expect: %s: got: %s", rollapp.ChannelId, rollappChannelOnHub,
			)
		}
	*/

	return chainID, &data, nil
}

// GetProofHeight returns the proof height of the packet
func (im IBCMiddleware) GetProofHeight(ctx sdk.Context, packetType commontypes.RollappPacket_Type,
	rollappPortOnHub string, rollappChannelOnHub string, sequence uint64,
) (uint64, error) {
	packetId := commontypes.NewPacketUID(packetType, rollappPortOnHub, rollappChannelOnHub, sequence)
	height, ok := types.FromIBCProofContext(ctx, packetId)
	if ok {
		return height.RevisionHeight, nil
	} else {
		err := errors.New("failed to get proof height from context")
		ctx.Logger().Error(err.Error(), "packetId", packetId)
		return 0, err
	}
}

// CheckIfFinalized checks if the packet is finalized and if so, updates the packet status
func (im IBCMiddleware) CheckIfFinalized(ctx sdk.Context, rollappID string, proofHeight uint64) (bool, error) {
	finalizedHeight, err := im.keeper.GetRollappFinalizedHeight(ctx, rollappID)
	if err != nil {
		if errors.Is(err, rollapptypes.ErrNoFinalizedStateYetForRollapp) {
			return false, nil
		}
		return false, err
	}

	return finalizedHeight >= proofHeight, nil
}
