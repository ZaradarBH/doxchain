package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgCreateClientRegistrationRegistryRequest) (*types.MsgCreateClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, isFound := k.GetClientRegistrationRegistry(
		ctx,
		msg.ClientRegistrationRegistry.Owner.Creator,
	)

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	k.SetClientRegistrationRegistry(
		ctx,
		msg.ClientRegistrationRegistry,
	)

	return &types.MsgCreateClientRegistrationRegistryResponse{}, nil
}

func (k msgServer) UpdateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgUpdateClientRegistrationRegistryRequest) (*types.MsgUpdateClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetClientRegistrationRegistry(
		ctx,
		msg.ClientRegistrationRegistry.Owner.Creator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "creator not set")
	}

	if msg.ClientRegistrationRegistry.Owner.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry)

	return &types.MsgUpdateClientRegistrationRegistryResponse{}, nil
}

func (k msgServer) DeleteClientRegistrationRegistry(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRegistryRequest) (*types.MsgDeleteClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetClientRegistrationRegistry(
		ctx,
		msg.Creator,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveClientRegistrationRegistry(
		ctx,
		msg.Creator,
	)

	return &types.MsgDeleteClientRegistrationRegistryResponse{}, nil
}
