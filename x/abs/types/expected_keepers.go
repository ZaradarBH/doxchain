package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

type BankKeeper interface {
	bankkeeper.SendKeeper

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
