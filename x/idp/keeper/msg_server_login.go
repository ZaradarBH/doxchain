package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgAuthenticationRequest) (*types.MsgAuthenticationResponse, error) {
	authToken, err := k.Keeper.Login(sdk.UnwrapSDKContext(goCtx), *msg)

	if err != nil {
		return nil, err
	}

	return &authToken, nil
}
