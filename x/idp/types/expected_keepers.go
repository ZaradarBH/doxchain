package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

//TODO: Remove dependency on AuthzKeeper
type AuthzKeeper interface {
	// Methods imported from authz should be defined here
}

//TODO: Remove dependency on EvidenceKeeper
type EvidenceKeeper interface {
	// Methods imported from evidence should be defined here
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
