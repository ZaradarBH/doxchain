package keeper

import (
	"context"
	"fmt"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDid) (*types.MsgCreateDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fullyQualifiedDidIdentifier := k.AppendDid(
		ctx,
		*msg.Did,
	)

	return &types.MsgCreateDidResponse{
		FullyQualifiedDidIdentifier: fullyQualifiedDidIdentifier,
	}, nil
}

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDid) (*types.MsgUpdateDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	fullyQualifiedDidIdentifier := msg.Did.GetFullyQualifiedDidIdentifier()
	val, found := k.GetDid(ctx, fullyQualifiedDidIdentifier)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", fullyQualifiedDidIdentifier))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Did.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDid(ctx, *msg.Did)

	return &types.MsgUpdateDidResponse{}, nil
}

func (k msgServer) DeleteDid(goCtx context.Context, msg *types.MsgDeleteDid) (*types.MsgDeleteDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetDid(ctx, msg.FullyQualifiedDidIdentifier)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.FullyQualifiedDidIdentifier))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDid(ctx, msg.FullyQualifiedDidIdentifier)

	return &types.MsgDeleteDidResponse{}, nil
}
