package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/be-heroes/doxchain/x/oauth/utils"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	isAuthorized, err := k.idpKeeper.AuthorizeCreator(ctx, msg.Tenant, msg.Creator)

	if !isAuthorized {
		return tokenResponse, err
	}

	tenantDeviceCodes, found := k.GetDeviceCodes(ctx, msg.Tenant)

	if !found {
		return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodes cache could not be found for tenant")
	}

	for index, deviceCodeEntry := range tenantDeviceCodes.Codes {
		if deviceCodeEntry.DeviceCode == msg.DeviceCode && deviceCodeEntry.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.DeviceCode))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = types.Bearer.String()
			tokenResponse.ExpiresIn = claims["exp"].(string)

			tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, msg.Tenant)

			if !found {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenInfo{
				Creator:     msg.Creator,
				Identifier:  claims["jti"].(string),
				ExpiresIn:   tokenResponse.ExpiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantDeviceCodes.Codes = append(tenantDeviceCodes.Codes[:index], tenantDeviceCodes.Codes[index+1:]...)

			k.SetDeviceCodes(ctx, tenantDeviceCodes)

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
