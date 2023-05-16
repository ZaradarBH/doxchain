package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	DidKeeperError = sdkerrors.Register(ModuleName, 1100, "Generic error in the Did keeper")	
	DidIdentifierCollisionError = sdkerrors.Register(ModuleName, 1101, "Errors related to DID identifier collisions")
	DidOwnershipError = sdkerrors.Register(ModuleName, 1102, "Errors related to DID ownership")
)