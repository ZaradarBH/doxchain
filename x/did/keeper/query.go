package keeper

import (
	"github.com/be-heroes/doxchain/x/did/types"
)

var _ types.QueryServer = Keeper{}
