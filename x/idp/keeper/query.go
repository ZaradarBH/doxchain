package keeper

import (
	"github.com/be-heroes/doxchain/x/idp/types"
)

var _ types.QueryServer = Keeper{}
