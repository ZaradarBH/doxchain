package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryBreakFactor(goCtx context.Context, req *types.QueryBreakFactorRequest) (result *types.QueryBreakFactorResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	result.BreakFactor = k.GetBreakFactor(sdk.UnwrapSDKContext(goCtx)).String()

	return result, nil
}
