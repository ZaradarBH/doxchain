package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/kyc/types"
	didUtils "github.com/be-heroes/doxchain/utils/did"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateKYCRegistration(goCtx context.Context, msg *types.MsgCreateKYCRegistrationRequest) (result *types.MsgCreateKYCRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetKYCRegistration(ctx, msg.Owner.GetW3CIdentifier())

	if found {
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

	return result, nil
}

func (k msgServer) DeleteKYCRegistration(goCtx context.Context, msg *types.MsgDeleteKYCRegistrationRequest) (result *types.MsgDeleteKYCRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userDid := didUtils.NewDidTokenFactory().Create(msg.Creator, "")
	match, found := k.GetKYCRegistration(ctx, userDid.GetW3CIdentifier())

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	if msg.Creator != match.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "impersonation not supported")
	}

	k.RemoveKYCRegistration(ctx, userDid.GetW3CIdentifier())

	return result, nil
}
