package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgCreateClientRegistrationRegistryRequest) (*types.MsgCreateClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, isFound := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry.Owner.GetW3CIdentifier())

	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "registry already exists")
	}

	k.SetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry)

	return &types.MsgCreateClientRegistrationRegistryResponse{}, nil
}

func (k msgServer) UpdateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgUpdateClientRegistrationRegistryRequest) (*types.MsgUpdateClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry.Owner.GetW3CIdentifier())

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "registry not found")
	}

	if msg.ClientRegistrationRegistry.Owner.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner of a registry can update it")
	}

	k.SetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry)

	return &types.MsgUpdateClientRegistrationRegistryResponse{}, nil
}

func (k msgServer) DeleteClientRegistrationRegistry(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRegistryRequest) (*types.MsgDeleteClientRegistrationRegistryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valFound, isFound := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistryW3CIdentifier)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "regitry not found")
	}

	if msg.Creator != valFound.Owner.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner of a registry can delete it")
	}

	k.RemoveClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistryW3CIdentifier)

	return &types.MsgDeleteClientRegistrationRegistryResponse{}, nil
}
