package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrKYCRegistrationImpersonation = sdkerrors.Register(ModuleName, 6000, "impersonation is not allowed")
	ErrKYCRegistrationExists = sdkerrors.Register(ModuleName, 6001, "registration already exists")
)
