package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKYCRegistration(goCtx context.Context, msg *types.MsgCreateKYCRegistrationRequest) (*types.MsgCreateKYCRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, isFound := k.GetKYCRegistration(ctx, msg.Creator)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	if msg.Owner.Creator != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not supported")
	}

	k.SetKYCRegistration(
		ctx,
		types.KYCRegistration{
			Owner:    msg.Owner,
			Approved: false,
		},
	)

	return &types.MsgCreateKYCRegistrationResponse{}, nil
}

func (k msgServer) DeleteKYCRegistration(goCtx context.Context, msg *types.MsgDeleteKYCRegistrationRequest) (*types.MsgDeleteKYCRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetKYCRegistration(ctx, msg.Creator)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation not supported")
	}

	k.RemoveKYCRegistration(ctx, msg.Creator)

	return &types.MsgDeleteKYCRegistrationResponse{}, nil
}
