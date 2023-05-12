package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/be-heroes/doxchain/x/oauth/utils"
)

func (k Keeper) GenerateClientCredentialToken(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	tokenResponse := types.MsgTokenResponse{}
	acl, err := k.idpKeeper.GetAccessClientList(ctx, msg.Tenant)

	if err != nil {
		return tokenResponse, err
	}

	for _, aclEntry := range acl.Entries {
		if aclEntry.Creator == msg.Creator {
			jwtToken := utils.NewJwtTokenFactory(utils.WithContext(&ctx)).Create(&msg)
			signedToken, err := jwtToken.SignedString([]byte(msg.ClientSecret))

			if err != nil {
				return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "Failed to create token")
			}

			//TODO: Save signed token to store until it is removed, if we even want to do that?
			tokenResponse.AccessToken = signedToken
			tokenResponse.TokenType = types.Bearer.String()

			//TODO: Make expire time configurable!
			tokenResponse.ExpiresIn = 1800

			break
		}
	}

	return tokenResponse, sdkerrors.Wrap(types.TokenServiceError, "ClientCredential TokenResponse could not be issued")
}
