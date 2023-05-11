package keeper

import (
	"github.com/be-heroes/doxchain/x/twins/types"
)

var _ types.QueryServer = Keeper{}
