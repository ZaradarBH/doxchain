package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Token(goCtx context.Context, msg *types.MsgTokenRequest) (*types.MsgTokenResponse, error) {
	accessToken, tokenType, expiresIn, err := k.Keeper.Token(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, msg.Scope, msg.ClientSecret, msg.AuthorizationCode, msg.DeviceCode, msg.ClientAssertion, msg.ClientAssertionType, msg.GrantType)

	if err != nil {
		return nil, err
	}

	return &types.MsgTokenResponse{
		AccessToken: accessToken,
		TokenType: tokenType,
		ExpiresIn: expiresIn,
	}, nil
}
