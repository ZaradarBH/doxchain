package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAMLRequest(goCtx context.Context, msg *types.MsgCreateAMLRequest) (*types.MsgCreateAMLRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetAMLRequest(ctx, msg.Creator)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}
	
	if msg.Creator != msg.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	var aMLRequest = types.AMLRequest{
		Owner:      msg.Owner,
		Approved: false,
	}

	k.SetAMLRequest(
		ctx,
		aMLRequest,
	)

	return &types.MsgCreateAMLRequestResponse{}, nil
}

func (k msgServer) DeleteAMLRequest(goCtx context.Context, msg *types.MsgDeleteAMLRequest) (*types.MsgDeleteAMLRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Check if the value exists
	valFound, isFound := k.GetAMLRequest(ctx, msg.Creator)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation is not supported")
	}

	k.RemoveAMLRequest(ctx, msg.Creator)

	return &types.MsgDeleteAMLRequestResponse{}, nil
}
