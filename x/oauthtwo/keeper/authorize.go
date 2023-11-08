package keeper

import (
	"time"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"	
	"github.com/be-heroes/doxchain/utils"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
)

func (k Keeper) Authorize(ctx sdk.Context, creator string, tenantW3CIdentifier string, clientRegistrationAppIdW3CIdentifier string, scope []string, userCode string) (authorizationCode string, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(creator)

	if err != nil {
		return authorizationCode, err
	}

	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "Authorize", creator)

	if err != nil {
		return authorizationCode, err
	}

	isAuthorized := k.idpKeeper.AuthorizeUser(ctx, creatorAddress, tenantW3CIdentifier)

	if !isAuthorized {
		return authorizationCode, err
	}

	var validScopes []string

	for _, requestedScope := range scope {
		validScopes = append(validScopes, k.idpKeeper.AuthorizeScope(ctx, tenantW3CIdentifier, clientRegistrationAppIdW3CIdentifier, requestedScope))
	}

	if len(validScopes) == 0 {
		return authorizationCode, sdkerrors.Wrap(types.TokenServiceError, "No valid scopes in request")
	}

	if len(userCode) > 0 {
		tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, tenantW3CIdentifier)

		if !found {
			return authorizationCode, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
		}

		userCodeFound := false

		for _, deviceCodeRegistryEntry := range tenantDeviceCodeRegistry.Codes {
			if deviceCodeRegistryEntry.UserCode == userCode && deviceCodeRegistryEntry.Owner.Creator == creator {
				userCodeFound = true
			}
		}

		if !userCodeFound {
			return authorizationCode, sdkerrors.Wrap(types.TokenServiceError, "UserCode not usable")
		}
	}

	authorizationCode = strconv.FormatUint(utils.HashStringToUint64(creator + ctx.BlockTime().String()), 10)
	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, tenantW3CIdentifier)

	if !found {
		return authorizationCode, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	authorizationCodeRegistryEntry := types.AuthorizationCodeRegistryEntry{
		Owner:             *didUtils.NewDidTokenFactory().Create(creator, didUrl),
		AuthorizationCode: authorizationCode,
		ExpiresAt:         ctx.BlockTime().Add(time.Minute * 3).Unix(),
	}

	tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes, authorizationCodeRegistryEntry)

	k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

	return authorizationCode, nil
}
