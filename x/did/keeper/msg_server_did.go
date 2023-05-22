package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDidRequest) (result *types.MsgCreateDidResponse, err error) {
	if msg.Creator != msg.Did.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.Keeper.GetDid(ctx, msg.Did.GetW3CIdentifier())

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Did already exists in store")
	}

	k.Keeper.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did)

	result.DidW3CIdentifier = msg.Did.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDidRequest) (result *types.MsgUpdateDidResponse, err error) {
	if msg.Creator != msg.Did.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDid(ctx, msg.Did.GetW3CIdentifier())

	if found && msg.Creator != match.Creator {		
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "illigal update attempt")
	}

	k.Keeper.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did)

	result.DidW3CIdentifier = msg.Did.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) DeleteDid(goCtx context.Context, msg *types.MsgDeleteDidRequest) (result *types.MsgDeleteDidResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDid(ctx, msg.DidW3CIdentifier)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Did does not exist in store")
	}

	if msg.Creator != match.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "illigal delete attempt")
	}

	k.Keeper.RemoveDid(ctx, msg.DidW3CIdentifier)

	return result, nil
}
