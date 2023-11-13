package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Token(goCtx context.Context, msg *types.MsgTokenRequest) (result *types.MsgTokenResponse, err error) {
	accessToken, tokenType, expiresIn, err := k.Keeper.Token(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, msg.Scope, msg.ClientSecret, msg.AuthorizationCode, msg.DeviceCode, msg.ClientAssertion, msg.ClientAssertionType, msg.GrantType)

	if err != nil {
		return nil, err
	}

	result = &types.MsgTokenResponse{
		AccessToken: accessToken,
		TokenType:   tokenType,
		ExpiresIn:   expiresIn,
	}
	
	return result, nil
}

func (k msgServer) DeviceCode(goCtx context.Context, msg *types.MsgDeviceCodeRequest) (result *types.MsgDeviceCodeResponse, err error) {
	deviceCode, userCode, verificationUri, err := k.Keeper.DeviceCode(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, msg.Scope)

	if err != nil {
		return nil, err
	}

	result = &types.MsgDeviceCodeResponse{
		DeviceCode:      deviceCode,
		UserCode:        userCode,
		VerificationUri: verificationUri,
	}
	
	return result, nil
}

func (k msgServer) Authorize(goCtx context.Context, msg *types.MsgAuthorizeRequest) (result *types.MsgAuthorizeResponse, err error) {
	authorizationCode, err := k.Keeper.Authorize(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, msg.Scope, msg.UserCode)

	if err != nil {
		return nil, err
	}

	result = &types.MsgAuthorizeResponse{
		AuthorizationCode: authorizationCode,
	}
	
	return result, nil
}
