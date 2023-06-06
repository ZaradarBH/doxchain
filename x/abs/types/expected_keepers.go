package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
}
