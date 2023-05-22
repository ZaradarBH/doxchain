package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgCreateClientRegistrationRegistryRequest) (result *types.MsgCreateClientRegistrationRegistryResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry.Owner.GetW3CIdentifier())

	if found {
		return nil, types.ErrClientRegistrationRegistryExists
	}

	k.SetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry)

	return result, nil
}

func (k msgServer) UpdateClientRegistrationRegistry(goCtx context.Context, msg *types.MsgUpdateClientRegistrationRegistryRequest) (result *types.MsgUpdateClientRegistrationRegistryResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry.Owner.GetW3CIdentifier())

	if !found {
		return nil, types.ErrClientRegistrationRegistryNotFound
	}

	if msg.ClientRegistrationRegistry.Owner.Creator != match.Owner.Creator {
		return nil, types.ErrImpersonation
	}

	k.SetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistry)

	return result, nil
}

func (k msgServer) DeleteClientRegistrationRegistry(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRegistryRequest) (result *types.MsgDeleteClientRegistrationRegistryResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.GetClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistryW3CIdentifier)

	if !found {
		return nil, types.ErrClientRegistrationRegistryNotFound
	}

	if msg.Creator != match.Owner.Creator {
		return nil, types.ErrImpersonation
	}

	k.RemoveClientRegistrationRegistry(ctx, msg.ClientRegistrationRegistryW3CIdentifier)

	return result, nil
}
