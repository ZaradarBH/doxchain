package keeper

import (
	"github.com/be-heroes/doxchain/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
