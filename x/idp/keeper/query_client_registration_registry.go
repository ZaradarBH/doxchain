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

func (k Keeper) ClientRegistrationRegistryAll(goCtx context.Context, req *types.QueryAllClientRegistrationRegistryRequest) (*types.QueryAllClientRegistrationRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ClientRegistrationRegistrys []types.ClientRegistrationRegistry
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	ClientRegistrationRegistryStore := prefix.NewStore(store, types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))

	pageRes, err := query.Paginate(ClientRegistrationRegistryStore, req.Pagination, func(key []byte, value []byte) error {
		var ClientRegistrationRegistry types.ClientRegistrationRegistry
		if err := k.cdc.Unmarshal(value, &ClientRegistrationRegistry); err != nil {
			return err
		}

		ClientRegistrationRegistrys = append(ClientRegistrationRegistrys, ClientRegistrationRegistry)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClientRegistrationRegistryResponse{ClientRegistrationRegistry: ClientRegistrationRegistrys, Pagination: pageRes}, nil
}

func (k Keeper) ClientRegistrationRegistry(goCtx context.Context, req *types.QueryGetClientRegistrationRegistryRequest) (*types.QueryGetClientRegistrationRegistryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetClientRegistrationRegistry(
		ctx,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetClientRegistrationRegistryResponse{ClientRegistrationRegistry: val}, nil
}
