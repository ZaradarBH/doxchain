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

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, creator string, tenantW3CIdentifier string, scope []string, clientRegistrationAppIdW3CIdentifier string, clientSecret string, clientAssertion string, clientAssertionType string) (accessToken string, tokenType string, expiresIn int64, err error) {
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "Token", creator)

	if err != nil {
		return accessToken, tokenType, expiresIn, err
	}

	switch clientAssertionType {
	case "urn:ietf:params:oauth:client-assertion-type:jwt-bearer":
		//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
		return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "Assertion is not supported: urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	default:
		jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(tenantW3CIdentifier, creator, clientRegistrationAppIdW3CIdentifier, time.Minute*3)
		claims := jwtToken.Claims.(jwt.MapClaims)
		accessToken, err := jwtToken.SignedString([]byte(clientSecret))

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
	}

	return accessToken, tokenType, expiresIn, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
