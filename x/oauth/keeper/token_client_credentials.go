package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/be-heroes/doxchain/x/oauth/utils"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return tokenResponse, err
	}

	switch msg.ClientAssertionType {
		case "urn:ietf:params:oauth:client-assertion-type:jwt-bearer":
			//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
			return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Assertion is not supported: urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
		default:
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.ClientSecret))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			tenantAccessTokens, found := k.GetAccessTokens(ctx, msg.Tenant)

			if !found {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokens.Tokens = append(tenantAccessTokens.Tokens, types.AccessToken{
				Creator:     msg.Creator,
				Uuid:        claims["jti"].(string),
				SignedToken: signedToken,
			})

			k.SetAccessTokens(ctx, tenantAccessTokens)

			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = types.Bearer.String()
			tokenResponse.ExpiresIn = claims["exp"].(string)
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
