package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistrationRelationship(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRelationshipRequest) (*types.MsgDeleteClientRegistrationRelationshipResponse, error) {
	//TODO: Fetch client registration and check Owner.Creator to see if the creator of the msg is allowed to delete it	
	err := k.Keeper.RemoveClientRegistrationRelationship(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.OwnerClientRegistrationW3CIdentifier, msg.DestinationClientRegistrationW3CIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteClientRegistrationRelationshipResponse{}, nil
}
