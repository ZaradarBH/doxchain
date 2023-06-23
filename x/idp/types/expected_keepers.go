package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI
}
