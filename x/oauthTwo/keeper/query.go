package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthTwo/types"
)

var _ types.QueryServer = Keeper{}
