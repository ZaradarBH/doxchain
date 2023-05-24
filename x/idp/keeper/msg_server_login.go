package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgAuthenticationRequest) (result *types.MsgAuthenticationResponse, err error) {
	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return nil, err
	}

	result.Token = k.Keeper.Login(sdk.UnwrapSDKContext(goCtx), creatorAddress, msg.TenantW3CIdentifier)

	if len(result.Token) == 0 {
		return nil, types.ErrLogin
	}

	return result, nil
}
