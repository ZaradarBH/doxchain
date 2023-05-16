package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClientRegistryAll(goCtx context.Context, req *types.QueryAllClientRegistryRequest) (*types.QueryAllClientRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ClientRegistrys []types.ClientRegistry
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ClientRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.ClientRegistryKeyPrefix))

	pageRes, err := query.Paginate(ClientRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var ClientRegistry types.ClientRegistry
		if err := k.cdc.Unmarshal(value, &ClientRegistry); err != nil {
			return err
		}

		ClientRegistrys = append(ClientRegistrys, ClientRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClientRegistryResponse{ClientRegistry: ClientRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) ClientRegistry(goCtx context.Context, req *types.QueryGetClientRegistryRequest) (*types.QueryGetClientRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetClientRegistry(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClientRegistryResponse{ClientRegistry: val}, nil
}
