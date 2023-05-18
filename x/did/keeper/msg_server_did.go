package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDidRequest) (*types.MsgCreateDidResponse, error) {
	if msg.Creator != msg.Did.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	err := k.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, false)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateDidResponse{
		FullyQualifiedDidIdentifier: msg.Did.GetW3CIdentifier(),
	}, nil
}

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDidRequest) (*types.MsgUpdateDidResponse, error) {
	if msg.Creator != msg.Did.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	err := k.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, true)

	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateDidResponse{}, nil
}

func (k msgServer) DeleteDid(goCtx context.Context, msg *types.MsgDeleteDidRequest) (*types.MsgDeleteDidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Check if the value exists
	valFound, isFound := k.GetDid(ctx, msg.FullyQualifiedDidIdentifier)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err := k.RemoveDid(ctx, msg.FullyQualifiedDidIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteDidResponse{}, nil
}
