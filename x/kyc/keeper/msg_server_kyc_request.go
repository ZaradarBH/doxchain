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
	_, isFound := k.GetKYCRequest(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var kYCRequest = types.KYCRequest{
		Creator:   msg.Creator,
		FirstName: msg.FirstName,
		LastName:  msg.LastName,
		Approved:  msg.Approved,
	}

	k.SetKYCRequest(
		ctx,
		kYCRequest,
	)
	return &types.MsgCreateKYCRequestResponse{}, nil
}

func (k msgServer) UpdateKYCRequest(goCtx context.Context, msg *types.MsgUpdateKYCRequest) (*types.MsgUpdateKYCRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKYCRequest(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kYCRequest = types.KYCRequest{
		Creator:   msg.Creator,
		FirstName: msg.FirstName,
		LastName:  msg.LastName,
		Approved:  msg.Approved,
	}

	k.SetKYCRequest(ctx, kYCRequest)

	return &types.MsgUpdateKYCRequestResponse{}, nil
}

func (k msgServer) DeleteKYCRequest(goCtx context.Context, msg *types.MsgDeleteKYCRequest) (*types.MsgDeleteKYCRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKYCRequest(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKYCRequest(ctx)

	return &types.MsgDeleteKYCRequestResponse{}, nil
}
