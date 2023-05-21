package keeper

import (
	"fmt"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
	"github.com/be-heroes/doxchain/x/abs/types"
	didUtils "github.com/be-heroes/doxchain/utils/did"
)

func (k msgServer) CreatePartitionedPoolRegistry(goCtx context.Context, msg *types.MsgCreatePartitionedPoolRegistryRequest) (*types.MsgCreatePartitionedPoolRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	partitionedPoolRegistry, isFound := k.GetPartitionedPoolRegistry(ctx, msg.Creator)
	didUrl, err := didUtils.CreateModuleDidUrl(types.ModuleName, fmt.Sprintf("%T", msg), msg.Creator)

	if err != nil {
		return nil, err
	}

	ownerDid := didUtils.NewDidTokenFactory().Create(msg.Creator, didUrl)

	if !isFound {
		partitionedPoolRegistry = types.PartitionedPoolRegistry{
			Owner: *ownerDid,
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

	return &types.MsgCreatePartitionedPoolRegistryResponse{
		PartitionedPoolW3CIdentifier: partitionedPoolRegistry.Owner.GetW3CIdentifier(),
	}, nil
}
