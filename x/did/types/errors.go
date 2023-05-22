package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrDidImpersonation         = sdkerrors.Register(ModuleName, 3001, "Did impersonation is not allowed")
	ErrDidExists                = sdkerrors.Register(ModuleName, 3002, "Did already exists")
	ErrDidNotFound              = sdkerrors.Register(ModuleName, 3003, "Did could not be found")
	ErrDidDocumentImpersonation = sdkerrors.Register(ModuleName, 3004, "DidDocument impersonation is not allowed")
	ErrDidDocumentExists        = sdkerrors.Register(ModuleName, 3005, "DidDocument already exists")
	ErrDidDocumentNotFound      = sdkerrors.Register(ModuleName, 3006, "Did could not be found")
)