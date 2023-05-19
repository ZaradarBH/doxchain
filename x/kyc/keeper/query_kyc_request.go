package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) KYCRegistration(goCtx context.Context, req *types.QueryGetKYCRegistrationRequest) (*types.QueryGetKYCRegistrationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKYCRegistration(ctx, req.Creator)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKYCRegistrationResponse{Request: val}, nil
}
