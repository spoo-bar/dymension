package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibctypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	tenderminttypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
	sequencertypes "github.com/dymensionxyz/dymension/v3/x/sequencer/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		hooks      types.MultiDelayedAckHooks
		paramstore paramtypes.Subspace

		rollappKeeper    types.RollappKeeper
		sequencerKeeper  types.SequencerKeeper
		ics4Wrapper      porttypes.ICS4Wrapper
		channelKeeper    types.ChannelKeeper
		connectionKeeper types.ConnectionKeeper
		clientKeeper     types.ClientKeeper
		types.EIBCKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	rollappKeeper types.RollappKeeper,
	sequencerKeeper types.SequencerKeeper,
	ics4Wrapper porttypes.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	connectionKeeper types.ConnectionKeeper,
	clientKeeper types.ClientKeeper,
	eibcKeeper types.EIBCKeeper,
	bankKeeper types.BankKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return &Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		memKey:           memKey,
		paramstore:       ps,
		rollappKeeper:    rollappKeeper,
		sequencerKeeper:  sequencerKeeper,
		ics4Wrapper:      ics4Wrapper,
		channelKeeper:    channelKeeper,
		clientKeeper:     clientKeeper,
		connectionKeeper: connectionKeeper,
		bankKeeper:       bankKeeper,
		EIBCKeeper:       eibcKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ExtractChainIDFromChannel(ctx sdk.Context, portID string, channelID string) (string, error) {
	_, clientState, err := k.channelKeeper.GetChannelClientState(ctx, portID, channelID)
	if err != nil {
		return "", fmt.Errorf("failed to extract clientID from channel: %w", err)
	}

	tmClientState, ok := clientState.(*ibctypes.ClientState)
	if !ok {
		return "", nil
	}

	return tmClientState.ChainId, nil
}

func (k Keeper) IsRollappsEnabled(ctx sdk.Context) bool {
	return k.rollappKeeper.GetParams(ctx).RollappsEnabled
}

func (k Keeper) GetRollapp(ctx sdk.Context, chainID string) (rollapptypes.Rollapp, bool) {
	return k.rollappKeeper.GetRollapp(ctx, chainID)
}

func (k Keeper) GetRollappFinalizedHeight(ctx sdk.Context, chainID string) (uint64, error) {
	res, err := k.rollappKeeper.StateInfo(ctx, &rollapptypes.QueryGetStateInfoRequest{
		RollappId: chainID,
		Finalized: true,
	})
	if err != nil {
		return 0, err
	}

	return (res.StateInfo.StartHeight + res.StateInfo.NumBlocks - 1), nil
}

// Validate that the channel is a rollapp channel and wasn't opened by external actor.
// Currently this is done by comparing:
// 1. the consensus state to the rollapp state for specific height.
// 2. the next validator set of the consensus state for previous height to the rollapp sequencer.
// This is done to protect against external malicious actors but has the assumption that the sequencer is honest.
// In order to remove this assumption we need to
// 1. Validate against the ChanOpenConfirm packet proof height
// 2. In case of a fraud use ibc channel abstraction (WIP).
func (k Keeper) ValidateRollappChannel(ctx sdk.Context, portID string, channelID string) error {
	chainID, err := k.ExtractChainIDFromChannel(ctx, portID, channelID)
	if err != nil {
		return err
	}
	rollapp, found := k.GetRollapp(ctx, chainID)
	if !found {
		return sdkerrors.Wrap(rollapptypes.ErrUnknownRollappID, "channel is not a rollapp channel")
	}
	// Get the rollapp state latest height and compare it to the client state height.
	// As the assumption the sequencer is honest we don't check the packet proof height.
	// Another assumption here is that the clientstate height >= rollapp state height as
	// the client state is updated directly while the rollapp state is updated every batch interval.
	latestStateIndex, found := k.rollappKeeper.GetLatestStateInfoIndex(ctx, rollapp.RollappId)
	if !found {
		return sdkerrors.Wrapf(rollapptypes.ErrUnknownRollappID, "state info not found for the rollapp: %s", rollapp.RollappId)
	}
	stateInfo, found := k.rollappKeeper.GetStateInfo(ctx, rollapp.RollappId, latestStateIndex.Index)
	if !found {
		return sdkerrors.Wrapf(rollapptypes.ErrUnknownRollappID, "state info not found for the rollapp: %s with index: %d", rollapp.RollappId, latestStateIndex.Index)
	}
	// Get the tm consensus state for the channel for the rollapp state height
	tmConsensusState, err := k.getTmConsensusStateForChannelAndHeight(ctx, portID, channelID, stateInfo.StartHeight)
	if err != nil {
		return err
	}
	// Compare the consensus state to the rollapp state. We assume the first BD is for the start height.
	rollappStateRoot := stateInfo.BDs.BD[0].StateRoot
	consensusStateRoot := tmConsensusState.GetRoot().GetHash()
	if !bytes.Equal(consensusStateRoot, rollappStateRoot) {
		errMsg := fmt.Sprintf("consensus state does not match the rollapp state at height %d: client root %x, rollapp root %x",
			stateInfo.StartHeight, consensusStateRoot, rollappStateRoot)
		return sdkerrors.Wrap(types.ErrMismatchedStateRoots, errMsg)
	}
	// Compare the validators set hash of the consensus state to the sequencer hash.
	// We take the previous height as we compare it against the next validator hash of previous block.
	previousRollappStateHeight := stateInfo.StartHeight - 1
	tmConsensusState, err = k.getTmConsensusStateForChannelAndHeight(ctx, portID, channelID, previousRollappStateHeight)
	if err != nil {
		return err
	}
	sequencer, found := k.sequencerKeeper.GetSequencer(ctx, stateInfo.Sequencer)
	if !found {
		return sdkerrors.Wrapf(sequencertypes.ErrUnknownSequencer, "sequencer %s not found for the rollapp %s", stateInfo.Sequencer, rollapp.RollappId)
	}
	seqPubKeyHash, err := sequencer.GetDymintPubKeyHash()
	if err != nil {
		return err
	}
	if !bytes.Equal(tmConsensusState.NextValidatorsHash, seqPubKeyHash) {
		errMsg := fmt.Sprintf("consensus state does not match the rollapp state at height %d: consensus state validators %x, rollapp sequencer %x",
			previousRollappStateHeight, tmConsensusState.NextValidatorsHash, stateInfo.Sequencer)
		return sdkerrors.Wrap(types.ErrMismatchedSequencer, errMsg)
	}
	return nil
}

func (k Keeper) GetConnectionEnd(ctx sdk.Context, portID string, channelID string) (connectiontypes.ConnectionEnd, error) {
	channel, found := k.channelKeeper.GetChannel(ctx, portID, channelID)
	if !found {
		return connectiontypes.ConnectionEnd{}, sdkerrors.Wrap(channeltypes.ErrChannelNotFound, channelID)
	}
	connectionEnd, found := k.connectionKeeper.GetConnection(ctx, channel.ConnectionHops[0])
	if !found {
		return connectiontypes.ConnectionEnd{}, sdkerrors.Wrap(connectiontypes.ErrConnectionNotFound, channel.ConnectionHops[0])
	}
	return connectionEnd, nil
}

// getTmConsensusStateForChannelAndHeight returns the tendermint consensus state for the channel for specific height
func (k Keeper) getTmConsensusStateForChannelAndHeight(ctx sdk.Context, portID string, channelID string, height uint64) (*tenderminttypes.ConsensusState, error) {
	// Get the client state for the channel for specific height
	connectionEnd, err := k.GetConnectionEnd(ctx, portID, channelID)
	if err != nil {
		return &tenderminttypes.ConsensusState{}, err
	}
	clientState, err := k.GetClientState(ctx, portID, channelID)
	if err != nil {
		return &tenderminttypes.ConsensusState{}, err
	}
	revisionHeight := clienttypes.NewHeight(clientState.GetLatestHeight().GetRevisionNumber(), height)
	consensusState, found := k.clientKeeper.GetClientConsensusState(ctx, connectionEnd.GetClientID(), revisionHeight)
	if !found {
		return nil, clienttypes.ErrConsensusStateNotFound
	}
	tmConsensusState, ok := consensusState.(*tenderminttypes.ConsensusState)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected tendermint consensus state, got %T", consensusState)
	}
	return tmConsensusState, nil
}

// GetClientState retrieves the client state for a given packet.
func (k Keeper) GetClientState(ctx sdk.Context, portID string, channelID string) (exported.ClientState, error) {
	connectionEnd, err := k.GetConnectionEnd(ctx, portID, channelID)
	if err != nil {
		return nil, err
	}
	clientState, found := k.clientKeeper.GetClientState(ctx, connectionEnd.GetClientID())
	if !found {
		return nil, clienttypes.ErrConsensusStateNotFound
	}
	return clientState, nil
}

func (k Keeper) BlockedAddr(addr string) bool {
	account, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return false
	}
	return k.bankKeeper.BlockedAddr(account)
}

/* -------------------------------------------------------------------------- */
/*                               Hooks handling                               */
/* -------------------------------------------------------------------------- */
func (k *Keeper) SetHooks(hooks types.MultiDelayedAckHooks) {
	if k.hooks != nil {
		panic("DelayedAckHooks already set")
	}
	k.hooks = hooks
}

func (k *Keeper) GetHooks() types.MultiDelayedAckHooks {
	return k.hooks
}

/* -------------------------------------------------------------------------- */
/*                                 ICS4Wrapper                                */
/* -------------------------------------------------------------------------- */

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (k Keeper) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string, sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return k.ics4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// WriteAcknowledgement wraps IBC ICS4Wrapper WriteAcknowledgement function.
// ICS29 WriteAcknowledgement is used for asynchronous acknowledgements.
func (k *Keeper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, acknowledgement exported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// WriteAcknowledgement wraps IBC ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(
	ctx sdk.Context,
	portID,
	channelID string,
) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}

// LookupModuleByChannel wraps ChannelKeeper LookupModuleByChannel function.
func (k *Keeper) LookupModuleByChannel(ctx sdk.Context, portID, channelID string) (string, *capabilitytypes.Capability, error) {
	return k.channelKeeper.LookupModuleByChannel(ctx, portID, channelID)
}
