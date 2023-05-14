package keeper

import (
	"context"

    "github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) CreatePartitionedPools(goCtx context.Context,  msg *types.MsgCreatePartitionedPools) (*types.MsgCreatePartitionedPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value already exists
    _, isFound := k.GetPartitionedPools(
        ctx,
        msg.Index,
        )
    if isFound {
        return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
    }

    var partitionedPools = types.PartitionedPools{
        Creator: msg.Creator,
        Index: msg.Index,
        
    }

   k.SetPartitionedPools(
   		ctx,
   		partitionedPools,
   	)
	return &types.MsgCreatePartitionedPoolsResponse{}, nil
}

func (k msgServer) UpdatePartitionedPools(goCtx context.Context,  msg *types.MsgUpdatePartitionedPools) (*types.MsgUpdatePartitionedPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value exists
    valFound, isFound := k.GetPartitionedPools(
        ctx,
        msg.Index,
    )
    if !isFound {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
    }

    // Checks if the the msg creator is the same as the current owner
    if msg.Creator != valFound.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

    var partitionedPools = types.PartitionedPools{
		Creator: msg.Creator,
		Index: msg.Index,
        
	}

	k.SetPartitionedPools(ctx, partitionedPools)

	return &types.MsgUpdatePartitionedPoolsResponse{}, nil
}

func (k msgServer) DeletePartitionedPools(goCtx context.Context,  msg *types.MsgDeletePartitionedPools) (*types.MsgDeletePartitionedPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Check if the value exists
    valFound, isFound := k.GetPartitionedPools(
        ctx,
        msg.Index,
    )
    if !isFound {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
    }

    // Checks if the the msg creator is the same as the current owner
    if msg.Creator != valFound.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }

	k.RemovePartitionedPools(
	    ctx,
	msg.Index,
    )

	return &types.MsgDeletePartitionedPoolsResponse{}, nil
}