package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidAll(goCtx context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var dids []types.Did
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	didStore := prefix.NewStore(store, types.KeyPrefix(types.DidKey))

	pageRes, err := query.Paginate(didStore, req.Pagination, func(key []byte, value []byte) error {
		var did types.Did
		if err := k.cdc.Unmarshal(value, &did); err != nil {
			return err
		}

		dids = append(dids, did)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidResponse{DidList: dids, Pagination: pageRes}, nil
}

func (k Keeper) Did(goCtx context.Context, req *types.QueryGetDidRequest) (*types.QueryGetDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	did, found := k.GetDid(ctx, req.FullyQualifiedDidIdentifier)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDidResponse{Did: did}, nil
}
