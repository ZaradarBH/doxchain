package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistrationRelationship(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRelationshipRequest) (result *types.MsgDeleteClientRegistrationRelationshipResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	relationship, found := k.Keeper.GetClientRegistrationRelationship(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.OwnerClientRegistrationW3CIdentifier, msg.DestinationClientRegistrationW3CIdentifier)

	if found {
		if relationship.Owner.Creator != msg.Creator {
			return nil, types.ErrImpersonation
		}

		k.Keeper.RemoveClientRegistrationRelationship(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.OwnerClientRegistrationW3CIdentifier, msg.DestinationClientRegistrationW3CIdentifier)
	}

	result = &types.MsgDeleteClientRegistrationRelationshipResponse{}

	return result, nil
}
