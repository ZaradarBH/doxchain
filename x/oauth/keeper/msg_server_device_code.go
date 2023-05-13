package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeviceCode(goCtx context.Context, msg *types.MsgDeviceCodeRequest) (*types.MsgDeviceCodeResponse, error) {
	deviceCode, err := k.Keeper.DeviceCode(sdk.UnwrapSDKContext(goCtx), *msg)

	if err != nil {
		return nil, err
	}

	return &deviceCode, nil
}
