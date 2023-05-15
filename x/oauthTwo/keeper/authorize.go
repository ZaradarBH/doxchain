package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

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

	response.AuthorizationCode, _ = utils.GenerateRandomString(32)

	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, msg.Tenant)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	authorizationCodeInfo := types.AuthorizationCodeInfo{
		Creator:           msg.Creator,
		AuthorizationCode: response.AuthorizationCode,
		ExpiresAt:         ctx.BlockTime().Add(time.Minute * 3).Unix(),
	}

	tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes, authorizationCodeInfo)

	k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

	return response, nil
}
