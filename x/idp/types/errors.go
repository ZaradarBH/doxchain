package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/idp module sentinel errors
var (
	LoginError          = sdkerrors.Register(ModuleName, 401, "Error authenticating user")
	IdpMasterKeyMissing = sdkerrors.Register(ModuleName, 1000, "IdpMasterKey is missing from store")
)
