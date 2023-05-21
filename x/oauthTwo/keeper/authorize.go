package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/be-heroes/doxchain/utils"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	didUtils "github.com/be-heroes/doxchain/utils/did"
)

func (k Keeper) Authorize(ctx sdk.Context, msg types.MsgAuthorizeRequest) (response types.MsgAuthorizeResponse, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	
	if err != nil {
		return response, err
	}

	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, fmt.Sprintf("%T", msg), msg.Creator)

	if err != nil {
		return response, err
	}

	isAuthorized, err := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, msg.TenantW3CIdentifier)

	if !isAuthorized {
		return response, err
	}

	var validScopes []string

	for _, requestedScope := range msg.Scope {
		validScope, err := k.idpKeeper.AuthorizeScope(ctx, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, requestedScope)

		if err != nil {
			return response, err
		}

		validScopes = append(validScopes, validScope)
	}

	if len(validScopes) == 0 {
		return response, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	if len(msg.UserCode) > 0 {
		tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, msg.TenantW3CIdentifier)

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
	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, msg.TenantW3CIdentifier)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	authorizationCodeRegistryEntry := types.AuthorizationCodeRegistryEntry{
		Owner: *didUtils.NewDidTokenFactory().Create(msg.Creator, didUrl),
		AuthorizationCode: response.AuthorizationCode,
		ExpiresAt: ctx.BlockTime().Add(time.Minute * 3).Unix(),
	}

	tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes, authorizationCodeRegistryEntry)

	k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

	return response, nil
}
