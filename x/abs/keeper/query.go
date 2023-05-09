package keeper

import (
	"doxchain/x/abs/types"
)

var _ types.QueryServer = Keeper{}
