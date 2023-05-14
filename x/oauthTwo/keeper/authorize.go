package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/be-heroes/doxchain/utils"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

// Authorize method for simple oauth keeper
func (k Keeper) Authorize(ctx sdk.Context, msg types.MsgAuthorizeRequest) (types.MsgAuthorizeResponse, error) {
	response := types.MsgAuthorizeResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	//TODO: Implement authorization code registry
	response.AuthorizationCode, _ = utils.GenerateRandomString(32)

	return response, nil
}
