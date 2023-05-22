package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	LoginError                      = sdkerrors.Register(ModuleName, 401, "Authentication error")
	IdpMasterKeyError               = sdkerrors.Register(ModuleName, 1000, "IdpMasterKey error")
	TenantListError                 = sdkerrors.Register(ModuleName, 1001, "TenantListError error")
	TenantError                     = sdkerrors.Register(ModuleName, 1002, "TenantListError error")
	AccessClientListError           = sdkerrors.Register(ModuleName, 1003, "AccessClientList error")
	ClientRegistrationRegistryError = sdkerrors.Register(ModuleName, 1004, "ClientRegistrationRegistry error")
)
