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

func (k Keeper) GenerateAuthorizationCodeToken(ctx sdk.Context, msg types.MsgTokenRequest) (response types.MsgTokenResponse, err error) {
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, fmt.Sprintf("%T", msg), msg.Creator)

	if err != nil {
		return response, err
	}
	
	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, msg.TenantW3CIdentifier)

	if !found {
		return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	for index, authorizationCodeRegistryEntry := range tenantAuthorizationCodeRegistry.Codes {
		if authorizationCodeRegistryEntry.AuthorizationCode == msg.AuthorizationCode && authorizationCodeRegistryEntry.Owner.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(msg.TenantW3CIdentifier, msg.Creator, msg.ClientRegistrationAppIdW3CIdentifier, time.Minute*3)
			claims := jwtToken.Claims.(jwt.MapClaims)
			signedToken, err := jwtToken.SignedString([]byte(msg.AuthorizationCode))

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

			tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes[:index], tenantAuthorizationCodeRegistry.Codes[index+1:]...)

			k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

			break
		}
	}

	return response, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCode TokenResponse could not be issued")
}
