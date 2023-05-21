package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteClientRegistrationRelationship(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRelationshipRequest) (*types.MsgDeleteClientRegistrationRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result, found := k.Keeper.GetClientRegistrationRelationship(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.OwnerClientRegistrationW3CIdentifier, msg.DestinationClientRegistrationW3CIdentifier)
	
	if found {		
		if result.Owner.Creator != msg.Creator {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid creator")
		}

		err := k.Keeper.RemoveClientRegistrationRelationship(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.OwnerClientRegistrationW3CIdentifier, msg.DestinationClientRegistrationW3CIdentifier)

		if err != nil {
			return nil, err
		}
	}
	
	return &types.MsgDeleteClientRegistrationRelationshipResponse{}, nil
}
