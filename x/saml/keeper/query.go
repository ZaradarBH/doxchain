package keeper

import (
	"github.com/be-heroes/doxchain/x/saml/types"
)

var _ types.QueryServer = Keeper{}
