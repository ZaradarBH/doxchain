package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

var _ types.QueryServer = Keeper{}
