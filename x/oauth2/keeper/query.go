package keeper

import (
	"github.com/be-heroes/doxchain/x/oauth2/types"
)

var _ types.QueryServer = Keeper{}
