package keeper

import (
	"doxchain/x/doxchain/types"
)

var _ types.QueryServer = Keeper{}
