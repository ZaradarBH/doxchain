package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Authorize(goCtx context.Context, msg *types.MsgAuthorizeRequest) (*types.MsgAuthorizeResponse, error) {
	authorizationCode, err := k.Keeper.Authorize(sdk.UnwrapSDKContext(goCtx), *msg)

	if err != nil {
		return nil, err
	}

	return &authorizationCode, nil
}
