package keeper

import (
	"github.com/be-heroes/doxchain/x/aml/types"
)

var _ types.QueryServer = Keeper{}
