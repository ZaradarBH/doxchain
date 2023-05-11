package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/oauth/types"
)

// Token method for simple oauth keeper
func (k Keeper) Token(ctx sdk.Context, msg types.MsgTokenRequest) (types.MsgTokenResponse, error) {
	switch msg.GrantType {
	case types.ClientCredentials.String():
		if msg.ClientAssertionType == "urn:ietf:params:oauth:client-assertion-type:jwt-bearer" {
			//TODO: Implement support for https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#second-case-access-token-request-with-a-certificate
		} else {
			return k.GenerateClientCredentialToken(ctx, msg)
		}
	case types.DeviceCode.String():
		//TODO: https://learn.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-device-code
	}

	return types.MsgTokenResponse{}, sdkerrors.Wrap(types.TokenServiceError, "Unsupported grant_type")
}
