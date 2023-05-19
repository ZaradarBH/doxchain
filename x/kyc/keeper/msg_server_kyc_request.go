package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKYCRegistration(goCtx context.Context, msg *types.MsgCreateKYCRegistrationRequest) (*types.MsgCreateKYCRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKYCRegistration(ctx, msg.Creator)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	if msg.Owner.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	var kYCRequest = types.KYCRegistration{
		Owner: msg.Owner,
		Approved: false,
	}

	k.SetKYCRegistration(
		ctx,
		kYCRequest,
	)

	return &types.MsgCreateKYCRegistrationResponse{}, nil
}

func (k msgServer) DeleteKYCRegistration(goCtx context.Context, msg *types.MsgDeleteKYCRegistrationRequest) (*types.MsgDeleteKYCRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKYCRegistration(ctx, msg.Creator)
	
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation not supported")
	}

	k.RemoveKYCRegistration(ctx, msg.Creator)

	return &types.MsgDeleteKYCRegistrationResponse{}, nil
}
