package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryBreakFactor(goCtx context.Context, req *types.QueryBreakFactorRequest) (*types.QueryBreakFactorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryBreakFactorResponse{
		Value: k.GetBreakFactor(sdk.UnwrapSDKContext(goCtx)).String(),
	}, nil
}
