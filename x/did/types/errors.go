package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrImpersonation            = sdkerrors.Register(ModuleName, 3000, "Impersonation is not allowed")
	ErrDidExists                = sdkerrors.Register(ModuleName, 3001, "Did already exists")
	ErrDidNotFound              = sdkerrors.Register(ModuleName, 3002, "Did could not be found")
	ErrDidDocumentExists        = sdkerrors.Register(ModuleName, 3003, "DidDocument already exists")
	ErrDidDocumentNotFound      = sdkerrors.Register(ModuleName, 3004, "DidDocument could not be found")
)