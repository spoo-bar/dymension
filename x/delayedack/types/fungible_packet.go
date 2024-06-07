package types

import (
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

type TransferData struct {
	transfertypes.FungibleTokenPacketData
	RollappID string
	// Proof height is only be populated if and only if the rollappID is not empty
	ProofHeight uint64
	// Finalized is only be populated if and only if the rollappID is not empty
	Finalized bool
}

// MustAmountInt returns the int amount. Should call validateBasic first!
func (d TransferData) MustAmountInt() math.Int {
	x, ok := sdk.NewIntFromString(d.Amount)
	if !ok {
		panic(fmt.Sprintf("parse transfer amount to Int: %s", d.Amount))
	}
	return x
}
