package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
)

type IdpKeeper interface {
	AuthorizeUser(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) (bool, error)
	GetTenantConfiguration(ctx sdk.Context, identifier string) (configuration idpTypes.TenantConfiguration, err error)
}
