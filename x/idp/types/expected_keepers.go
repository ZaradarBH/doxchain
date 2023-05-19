package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

//TODO: Remove dependency on AuthzKeeper
type AuthzKeeper interface {
}

//TODO: Remove dependency on EvidenceKeeper
type EvidenceKeeper interface {
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
