package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Token(goCtx context.Context, msg *types.MsgToken) (*types.MsgTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTokenResponse{}, nil
}
