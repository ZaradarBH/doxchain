package keeper

import (
	"github.com/be-heroes/doxchain/x/saml2/types"
)

var _ types.QueryServer = Keeper{}
