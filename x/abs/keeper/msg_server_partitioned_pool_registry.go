package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/be-heroes/doxchain/x/abs/types"
	utils "github.com/be-heroes/doxchain/utils/did"
)

func (k msgServer) CreatePartitionedPoolRegistry(goCtx context.Context, msg *types.MsgCreatePartitionedPoolRegistryRequest) (*types.MsgCreatePartitionedPoolRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	partitionedPoolRegistry, isFound := k.GetPartitionedPoolRegistry(ctx, msg.Creator)
	creatorDid := utils.NewDidTokenFactory().Create(msg.Creator, "")

	if !isFound {
		partitionedPoolRegistry = types.PartitionedPoolRegistry{
			Owner: creatorDid,
			Pools: []types.PartitionedPool{},
		}
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
