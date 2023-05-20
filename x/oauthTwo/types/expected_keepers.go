package types

import (
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IdpKeeper interface {
	AuthorizeCreator(ctx sdk.Context, fullyQualifiedW3CIdentifier string, creator string) (bool, error)
	GetTenantConfiguration(ctx sdk.Context, identifier string) (configuration idpTypes.TenantConfiguration, err error)
	SetDeviceCodeRegistry(ctx sdk.Context, deviceCodeRegistry idpTypes.DeviceCodeRegistry)
	GetDeviceCodeRegistry(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (val idpTypes.DeviceCodeRegistry, found bool)
}
