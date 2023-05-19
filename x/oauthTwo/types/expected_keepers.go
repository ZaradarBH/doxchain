package types

import (
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

//TODO: Remove AuthzKeeper from module
type AuthzKeeper interface {
	// Methods imported from authz should be defined here
}

//TODO: Remove EvidenceKeeper from module
type EvidenceKeeper interface {
	// Methods imported from evidence should be defined here
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type IdpKeeper interface {
	AuthorizeCreator(ctx sdk.Context, fullyQualifiedW3CIdentifier string, creator string) (bool, error)
	GetTenantConfiguration(ctx sdk.Context, identifier string) (configuration idpTypes.TenantConfiguration, err error)
	SetDeviceCodeRegistry(ctx sdk.Context, deviceCodeRegistry idpTypes.DeviceCodeRegistry)
	GetDeviceCodeRegistry(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (val idpTypes.DeviceCodeRegistry, found bool)
}
