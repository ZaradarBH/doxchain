package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
)

type IdpKeeper interface {
	AuthorizeCreator(ctx sdk.Context, fullyQualifiedW3CIdentifier string, creator string) (bool, error)
	GetTenantConfiguration(ctx sdk.Context, identifier string) (configuration idpTypes.TenantConfiguration, err error)
}
