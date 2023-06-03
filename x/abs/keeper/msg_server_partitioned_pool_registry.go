package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	didUtils "github.com/be-heroes/doxchain/utils/did"
	"github.com/be-heroes/doxchain/x/abs/types"
)

func (k msgServer) CreatePartitionedPoolRegistry(goCtx context.Context, msg *types.MsgCreatePartitionedPoolRegistryRequest) (result *types.MsgCreatePartitionedPoolRegistryResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	partitionedPoolRegistry, found := k.Keeper.GetPartitionedPoolRegistry(ctx, msg.Creator)
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, "PartitionedPoolRegistry", msg.Creator)

	if err != nil {
		return nil, err
	}

	ownerDid := didUtils.NewDidTokenFactory().Create(msg.Creator, didUrl)

	if !found {
		partitionedPoolRegistry = types.PartitionedPoolRegistry{
			Owner: *ownerDid,
			Pools: make([]types.PartitionedPool, 0),
		}
	}

	for _, partitionedPool := range partitionedPoolRegistry.Pools {
		if partitionedPool.Denom == msg.Denom {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Denom already exists")
		}
	}

	partitionedPoolRegistry.Pools = append(partitionedPoolRegistry.Pools, types.PartitionedPool{Denom: msg.Denom})

	k.Keeper.SetPartitionedPoolRegistry(
		ctx,
		partitionedPoolRegistry,
	)

	result = &types.MsgCreatePartitionedPoolRegistryResponse{
		PartitionedPoolW3CIdentifier: partitionedPoolRegistry.Owner.GetW3CIdentifier(),
	}

	return result, nil
}
