package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/idp module sentinel errors
var (
	LoginError            = sdkerrors.Register(ModuleName, 401, "Authentication error")
	IdpMasterKeyError     = sdkerrors.Register(ModuleName, 1000, "IdpMasterKey error")
	AccessClientListError = sdkerrors.Register(ModuleName, 1001, "AccessClientList error")
)
