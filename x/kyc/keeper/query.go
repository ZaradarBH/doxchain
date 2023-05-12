package keeper

import (
	"github.com/be-heroes/doxchain/x/kyc/types"
)

var _ types.QueryServer = Keeper{}
