package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (response types.MsgTokenResponse, err error) {
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, fmt.Sprintf("%T", msg), msg.Creator)

	if err != nil {
		return response, err
	}

	tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, msg.TenantW3CIdentifier)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}

	for index, deviceCodeInfo := range tenantDeviceCodeRegistry.Codes {
		if deviceCodeInfo.DeviceCode == msg.DeviceCode && deviceCodeInfo.Owner.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.TenantW3CIdentifier, msg.Creator, msg.ClientRegistrationAppIdW3CIdentifier, time.Minute*3)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.DeviceCode))

			if err != nil {
				return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			response.AccessToken = signedToken
			response.TokenType = types.Bearer.String()
			response.ExpiresIn = claims["exp"].(int64)

			tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, msg.TenantW3CIdentifier)

			if !found {
				return response, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenRegistryEntry{
				Owner: *didUtils.NewDidTokenFactory().Create(msg.Creator, didUrl),
				Jti: claims["jti"].(string),
				ExpiresAt: response.ExpiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes[:index], tenantDeviceCodeRegistry.Codes[index+1:]...)

			k.idpKeeper.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

			break
		}
	}

	return response, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
