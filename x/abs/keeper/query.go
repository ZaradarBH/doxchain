package keeper

import (
	"github.com/be-heroes/doxchain/x/abs/types"
)

var _ types.QueryServer = Keeper{}
