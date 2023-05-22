package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didUtils "github.com/be-heroes/doxchain/utils/did"
	utils "github.com/be-heroes/doxchain/utils/jwt"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateDeviceCodeToken(ctx sdk.Context, creator string, tenantW3CIdentifier string, scope []string, clientRegistrationAppIdW3CIdentifier string, deviceCode string) (accessToken string, tokenType string, expiresIn int64, err error) {
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "Token", creator)

	if err != nil {
		return accessToken, tokenType, expiresIn, err
	}

	tenantDeviceCodeRegistry, found := k.idpKeeper.GetDeviceCodeRegistry(ctx, tenantW3CIdentifier)

	if !found {
		return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "DeviceCodeRegistry cache could not be found for tenant")
	}

	for index, deviceCodeInfo := range tenantDeviceCodeRegistry.Codes {
		if deviceCodeInfo.DeviceCode == deviceCode && deviceCodeInfo.Owner.Creator == creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(tenantW3CIdentifier, creator, clientRegistrationAppIdW3CIdentifier, time.Minute*3)
			claims := jwtToken.Claims.(jwt.MapClaims)
			accessToken, err := jwtToken.SignedString([]byte(deviceCode))

			if err != nil {
				return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "Failed to create access token")
			}

			tokenType = types.Bearer.String()
			expiresIn = claims["exp"].(int64)

			tenantAccessTokenRegistry, found := k.GetAccessTokenRegistry(ctx, tenantW3CIdentifier)

			if !found {
				return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "Failed to fetch access tokens cache for tenant")
			}

			tenantAccessTokenRegistry.Issued = append(tenantAccessTokenRegistry.Issued, types.AccessTokenRegistryEntry{
				Owner:     *didUtils.NewDidTokenFactory().Create(creator, didUrl),
				Jti:       claims["jti"].(string),
				ExpiresAt: expiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantDeviceCodeRegistry.Codes = append(tenantDeviceCodeRegistry.Codes[:index], tenantDeviceCodeRegistry.Codes[index+1:]...)

			k.idpKeeper.SetDeviceCodeRegistry(ctx, tenantDeviceCodeRegistry)

			break
		}
	}

	return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "DeviceCode TokenResponse could not be issued")
}
