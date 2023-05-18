package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKYCRequest(goCtx context.Context, msg *types.MsgCreateKYCRequest) (*types.MsgCreateKYCRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKYCRequest(ctx, msg.Creator)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	if msg.Owner.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	var kYCRequest = types.KYCRequest{
		Owner: msg.Owner,
		Approved: false,
	}

	k.SetKYCRequest(
		ctx,
		kYCRequest,
	)

	return &types.MsgCreateKYCRequestResponse{}, nil
}

func (k msgServer) DeleteKYCRequest(goCtx context.Context, msg *types.MsgDeleteKYCRequest) (*types.MsgDeleteKYCRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKYCRequest(ctx, msg.Creator)
	
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation not supported")
	}

	k.RemoveKYCRequest(ctx, msg.Creator)

	return &types.MsgDeleteKYCRequestResponse{}, nil
}
