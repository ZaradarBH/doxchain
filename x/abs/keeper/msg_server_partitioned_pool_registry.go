package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePartitionedPoolRegistry(goCtx context.Context, msg *types.MsgCreatePartitionedPoolRegistryRequest) (*types.MsgCreatePartitionedPoolRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	partitionedPoolRegistry, isFound := k.GetPartitionedPoolRegistry(ctx, msg.Creator)

	if !isFound {
		partitionedPoolRegistry = types.PartitionedPoolRegistry{
			Creator: msg.Creator,
			Pools:   []types.PartitionedPool{},
		}
	}

	if partitionedPoolRegistry.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Only the creator of a partioned pool can interface with it after its inception")
	}

	for _, partitionedPool := range partitionedPoolRegistry.Pools {
		if partitionedPool.Denom == msg.Denom {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Denom already exists")
		}
	}

	partitionedPoolRegistry.Pools = append(partitionedPoolRegistry.Pools, types.PartitionedPool{Denom: msg.Denom})

	k.SetPartitionedPoolRegistry(
		ctx,
		partitionedPoolRegistry,
	)

	return &types.MsgCreatePartitionedPoolRegistryResponse{}, nil
}
