package types

import (
	idpTypes "github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IdpKeeper interface {
	AuthorizeUser(ctx sdk.Context, user sdk.AccAddress, tenantW3CIdentifier string) (bool)
	AuthorizeScope(ctx sdk.Context, clientRegistrationRegistryW3CIdentitifer string, clientRegistrationW3CIdentitifer string, scope string) (string)
	GetTenantConfiguration(ctx sdk.Context, tenantW3CIdentifier string) (configuration idpTypes.TenantConfiguration)
	SetDeviceCodeRegistry(ctx sdk.Context, deviceCodeRegistry idpTypes.DeviceCodeRegistry)
	GetDeviceCodeRegistry(ctx sdk.Context, deviceCodeRegistryW3CIdentifier string) (deviceCodeRegistry idpTypes.DeviceCodeRegistry, found bool)
}
