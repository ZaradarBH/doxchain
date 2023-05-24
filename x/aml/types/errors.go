package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrImpersonation =         sdkerrors.Register(ModuleName, 2000, "impersonation is not allowed")
	ErrAMLRegistrationExists = sdkerrors.Register(ModuleName, 2001, "registration already exists")
)
