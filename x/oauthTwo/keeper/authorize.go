package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/be-heroes/doxchain/utils"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	didUtils "github.com/be-heroes/doxchain/utils/did"
)

func (k Keeper) Authorize(ctx sdk.Context, msg types.MsgAuthorizeRequest) (types.MsgAuthorizeResponse, error) {
	response := types.MsgAuthorizeResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	if len(msg.UserCode) > 0 {
		tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, msg.Tenant)

		if !found {
			return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
		}

		userCodeFound := false

		for _, deviceCodeRegistryEntry := range tenantDeviceCodeRegistry.Codes {
			if deviceCodeRegistryEntry.UserCode == msg.UserCode && deviceCodeRegistryEntry.Owner.Creator == msg.Creator {
				userCodeFound = true
			}
		}

		if !userCodeFound {
			return response, sdkerrors.Wrap(types.TokenServiceError, "UserCode not usable")
		}
	}

	response.AuthorizationCode, _ = utils.GenerateRandomString(32)

	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, msg.Tenant)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	authorizationCodeRegistryEntry := types.AuthorizationCodeRegistryEntry{
		Owner: *didUtils.NewDidTokenFactory().Create(msg.Creator, ""),
		AuthorizationCode: response.AuthorizationCode,
		ExpiresAt: ctx.BlockTime().Add(time.Minute * 3).Unix(),
	}

	tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes, authorizationCodeRegistryEntry)

	k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

	return response, nil
}
