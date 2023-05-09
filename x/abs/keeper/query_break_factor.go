package keeper

import (
	"context"

	"doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryBreakFactor(goCtx context.Context, req *types.QueryBreakFactorRequest) (*types.QueryBreakFactorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	value := k.GetBreakFactor(ctx)

	return &types.QueryBreakFactorResponse{
		Value: value.String(),
	}, nil
}
