package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/dgrijalva/jwt-go"
)

// Token method for simple oauth keeper
func (k Keeper) Token(ctx sdk.Context, msg types.MsgToken) (types.TokenResponse, error) {
	switch msg.GrantType {
		case "client_credentials":
			if msg.ClientAssertionType == "urn:ietf:params:oauth:client-assertion-type:jwt-bearer" {
				//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
			} else {
				return k.GenerateClientCredentialToken(ctx, msg)
			}
		case "device_code":
			//TODO: https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-device-code
	}

	return types.TokenResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgToken) (types.TokenResponse, error) {
	tokenResponse := types.TokenResponse{}
	acl, err := k.GetAccessClientList(ctx, msg.Tenant)

	if err != nil {
		return tokenResponse, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.ClientId == msg.ClientId && aclEntry.ClientSecret == msg.ClientSecret {
			// Create a new token
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			
			claims["iss"] = types.ModuleName
			claims["sub"] = msg.ClientId
			//Figure out best way to manage exp´timestamp.
			//claims["exp"] = jwt.Now().Add(time.Minute * 30).Unix()
		
			// Sign the token with the client secret
			signedToken, err := token.SignedString([]byte(aclEntry.ClientSecret))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
			}
		
			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = "Bearer"
			tokenResponse.ExpiresIn = 1800
			
			break;
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
