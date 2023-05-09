package keeper

import (
	"doxchain/x/did/types"
)

var _ types.QueryServer = Keeper{}
