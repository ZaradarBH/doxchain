package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgAuthenticationRequest) (response *types.MsgAuthenticationResponse, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	
	if err != nil {
		return response, err
	}

	tokenString, err := k.Keeper.Login(sdk.UnwrapSDKContext(goCtx), creatorAddress, msg.TenantW3CIdentifier)

	if err != nil {
		return response, err
	}

	return &types.MsgAuthenticationResponse{
		Token: tokenString,
	}, nil
}
