package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateAuthorizationCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	response := types.MsgTokenResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, msg.Tenant)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	for index, authorizationCodeRegistryEntry := range tenantAuthorizationCodeRegistry.Codes {
		if authorizationCodeRegistryEntry.AuthorizationCode == msg.AuthorizationCode && authorizationCodeRegistryEntry.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.Tenant, msg.Creator, msg.ClientId, time.Minute*3)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.AuthorizationCode))

			if err != nil {
				return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			response.AccessToken = signedToken
			response.TokenType = types.Bearer.String()
			response.ExpiresIn = claims["exp"].(int64)

			tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, msg.Tenant)

			if !found {
				return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenRegistryEntry{
				Creator:    msg.Creator,
				Identifier: claims["jti"].(string),
				ExpiresAt:  response.ExpiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes[:index], tenantAuthorizationCodeRegistry.Codes[index+1:]...)

			k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

			break
		}
	}

	return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCode TokenResponse could not be issued")
}
