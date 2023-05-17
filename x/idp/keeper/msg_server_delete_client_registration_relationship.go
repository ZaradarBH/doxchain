package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistrationRelationship(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRelationshipRequest) (*types.MsgDeleteClientRegistrationRelationshipResponse, error) {
	err := k.Keeper.RemoveClientRegistrationRelationship(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRelationshipRegistryEntry)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteClientRegistrationRelationshipResponse{}, nil
}
