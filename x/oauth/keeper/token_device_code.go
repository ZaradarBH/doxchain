package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/be-heroes/doxchain/utils"
	"github.com/golang-jwt/jwt"
)

//TODO: Implement devicecode message handler / logic to generate device code
func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return tokenResponse, err
	}

	tenantDeviceCodeRegistry, found := k.GetDeviceCodeRegistry(ctx, msg.Tenant)

	if !found {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}

	for index, deviceCodeInfo := range tenantDeviceCodeRegistry.Codes {
		if deviceCodeInfo.DeviceCode == msg.DeviceCode && deviceCodeInfo.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.Tenant, msg.Creator, msg.ClientId, time.Minute * 15)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.DeviceCode))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = types.Bearer.String()			
			tokenResponse.ExpiresIn = claims["exp"].(int64)

			tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, msg.Tenant)

			if !found {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenInfo{
				Creator:     msg.Creator,
				Identifier:  claims["jti"].(string),
				ExpiresAt:   tokenResponse.ExpiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes[:index], tenantDeviceCodeRegistry.Codes[index+1:]...)

			k.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
