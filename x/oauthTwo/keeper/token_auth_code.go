package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/be-heroes/doxchain/utils/jwt"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/golang-jwt/jwt"
)

func (k Keeper) GenerateAuthorizationCodeToken(ctx sdk.Context, creator string, tenantW3CIdentifier string, scope []string, clientRegistrationAppIdW3CIdentifier string, authorizationCode string) (accessToken string, tokenType string, expiresIn int64, err error) {
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "Token", creator)

	if err != nil {
		return accessToken, tokenType, expiresIn, err
	}
	
	tenantAuthorizationCodeRegistry, found := k.GetAuthorizationCodeRegistry(ctx, tenantW3CIdentifier)

	if !found {
		return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCodeRegistry cache could not be found for tenant")
	}

	for index, authorizationCodeRegistryEntry := range tenantAuthorizationCodeRegistry.Codes {
		if authorizationCodeRegistryEntry.AuthorizationCode == authorizationCode && authorizationCodeRegistryEntry.Owner.Creator == creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(tenantW3CIdentifier, creator, clientRegistrationAppIdW3CIdentifier, time.Minute*3)
			claims := jwtToken.Claims.(jwt.MapClaims)
			accessToken, err := jwtToken.SignedString([]byte(authorizationCode))

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
				Owner: *didUtils.NewDidTokenFactory().Create(creator, didUrl),
				Jti: claims["jti"].(string),
				ExpiresAt: expiresIn,
			})

			k.SetAccessTokenRegistry(ctx, tenantAccessTokenRegistry)

			tenantAuthorizationCodeRegistry.Codes = append(tenantAuthorizationCodeRegistry.Codes[:index], tenantAuthorizationCodeRegistry.Codes[index+1:]...)

			k.SetAuthorizationCodeRegistry(ctx, tenantAuthorizationCodeRegistry)

			break
		}
	}

	return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "AuthorizationCode TokenResponse could not be issued")
}
