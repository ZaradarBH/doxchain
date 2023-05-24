package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeviceCode(goCtx context.Context, msg *types.MsgDeviceCodeRequest) (*types.MsgDeviceCodeResponse, error) {
	deviceCode, userCode, verificationUri, err := k.Keeper.DeviceCode(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.TenantW3CIdentifier, msg.ClientRegistrationAppIdW3CIdentifier, msg.Scope)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeviceCodeResponse{
		DeviceCode:      deviceCode,
		UserCode:        userCode,
		VerificationUri: verificationUri,
	}, nil
}
