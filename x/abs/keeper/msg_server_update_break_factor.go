package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateBreakFactor(goCtx context.Context, msg *types.MsgUpdateBreakFactor) (*types.MsgUpdateBreakFactorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateBreakFactorResponse{}, nil
}
