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
	isAuthorized, err := k.AuthorizeRequest(ctx, msg)
	
	if !isAuthorized {
		return tokenResponse, err
	}

	jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg)
	claims := jwtToken.Claims.(jwt.MapClaims)
	signedToken, err := jwtToken.SignedString([]byte(msg.ClientSecret))

	if err != nil {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
	}
	
	tenantAccessTokens, found := k.GetAccessTokens(ctx, msg.Tenant)
	
	if !found {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens for tenant")
	}

	tenantAccessTokens.Tokens = append(tenantAccessTokens.Tokens, types.AccessToken{
		Creator: msg.Creator,
		Uuid: claims["jti"].(string),
		SignedToken: signedToken,
	})

	k.SetAccessTokens(ctx, tenantAccessTokens)

	tokenResponse.AccessToken = signedToken
	tokenResponse.TokenType = types.Bearer.String()
	tokenResponse.ExpiresIn = claims["exp"].(string)

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
