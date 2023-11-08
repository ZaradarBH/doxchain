package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	TokenServiceError = sdkerrors.Register(ModuleName, 500, "Simple Token Service error")
)
