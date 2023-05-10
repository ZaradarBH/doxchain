package keeper

import (
	"github.com/be-heroes/doxchain/x/doxchain/types"
)

var _ types.QueryServer = Keeper{}
