package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrLogin                                   = sdkerrors.Register(ModuleName, 401, "Login forbidden")
	ErrImpersonation                           = sdkerrors.Register(ModuleName, 5000, "Impersonation is not allowed")
	ErrTenantExists                            = sdkerrors.Register(ModuleName, 5001, "TenantRegistry already exists")
	ErrTenantNotFound                          = sdkerrors.Register(ModuleName, 5002, "TenantRegistry could not be found")
	ErrTenantRegistryExists                    = sdkerrors.Register(ModuleName, 5003, "TenantRegistry already exists")
	ErrTenantRegistryNotFound                  = sdkerrors.Register(ModuleName, 5004, "TenantRegistry could not be found")
	ErrClientRegistrationRegistryExists        = sdkerrors.Register(ModuleName, 5005, "ClientRegistrationRegistry already exists")
	ErrClientRegistrationRegistryNotFound      = sdkerrors.Register(ModuleName, 5006, "ClientRegistrationRegistry could not be found")
	ErrClientRegistrationRelationshipNotFound  = sdkerrors.Register(ModuleName, 5007, "ClientRegistrationRelationship could not be found")
	ErrClientRegistrationRelationshipInvalid   = sdkerrors.Register(ModuleName, 5008, "ClientRegistrationRelationship is not valid")
)
