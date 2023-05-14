package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/utils"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	response := types.MsgTokenResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return response, err
	}

	switch msg.ClientAssertionType {
	case "urn:ietf:params:oauth:client-assertion-type:jwt-bearer":
		//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
		return response, sdkerrors.Wrap(types.TokenServiceError, "Assertion is not supported: urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	default:
		jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.Tenant, msg.Creator, msg.ClientId, time.Minute*3)
		claims := jwtToken.Claims.(jwt.MapClaims)
		signedToken, err := jwtToken.SignedString([]byte(msg.ClientSecret))

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

		tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenInfo{
			Creator:    msg.Creator,
			Identifier: claims["jti"].(string),
			ExpiresAt:  response.ExpiresIn,
		})

		k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

	}

	return response, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
