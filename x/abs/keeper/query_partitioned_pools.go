package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/be-heroes/doxchain/x/abs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PartitionedPoolsAll(goCtx context.Context, req *types.QueryAllPartitionedPoolsRequest) (*types.QueryAllPartitionedPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var partitionedPoolss []types.PartitionedPools
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	partitionedPoolsStore := prefix.NewStore(store, types.KeyPrefix(types.PartitionedPoolsKeyPrefix))

	pageRes, err := query.Paginate(partitionedPoolsStore, req.Pagination, func(key []byte, value []byte) error {
		var partitionedPools types.PartitionedPools
		if err := k.cdc.Unmarshal(value, &partitionedPools); err != nil {
			return err
		}

		partitionedPoolss = append(partitionedPoolss, partitionedPools)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPartitionedPoolsResponse{PartitionedPools: partitionedPoolss, Pagination: pageRes}, nil
}

func (k Keeper) PartitionedPools(goCtx context.Context, req *types.QueryGetPartitionedPoolsRequest) (*types.QueryGetPartitionedPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPartitionedPools(
	    ctx,
	    req.Index,
        )
	if !found {
	    return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPartitionedPoolsResponse{PartitionedPools: val}, nil
}