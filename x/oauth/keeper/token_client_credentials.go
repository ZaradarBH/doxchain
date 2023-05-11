package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/dgrijalva/jwt-go"
)

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	acl, err := k.idpKeeper.GetAccessClientList(ctx, msg.Tenant)

	if err != nil {
		return tokenResponse, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.ClientId == msg.ClientId && aclEntry.ClientSecret == msg.ClientSecret {
			jwtToken := jwt.New(jwt.SigningMethodHS256)
			claims := jwtToken.Claims.(jwt.MapClaims)

			claims["iss"] = types.ModuleName
			claims["sub"] = msg.ClientId
			//TODO: Implement oracle logic for adding unix timestamps to each block so we can use those to improve precision when issuing claims, assertions, etc
			claims["exp"] = time.Unix(int64(ctx.BlockHeight()), 0)

			signedToken, err := jwtToken.SignedString([]byte(aclEntry.ClientSecret))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
			}

			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = "Bearer"
			tokenResponse.ExpiresIn = 1800

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
