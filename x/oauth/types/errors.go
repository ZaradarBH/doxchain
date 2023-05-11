package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oauth module sentinel errors
var (
	TokenServiceError = sdkerrors.Register(ModuleName, 500, "sample error")
)
