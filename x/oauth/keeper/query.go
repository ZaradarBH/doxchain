package keeper

import (
	"github.com/be-heroes/doxchain/x/oauth/types"
)

var _ types.QueryServer = Keeper{}
