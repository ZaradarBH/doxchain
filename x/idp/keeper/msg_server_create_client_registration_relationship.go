package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistrationRelationship(goCtx context.Context, msg *types.MsgCreateClientRegistrationRelationshipRequest) (*types.MsgCreateClientRegistrationRelationshipResponse, error) {
	k.Keeper.SetClientRegistrationRelationship(sdk.UnwrapSDKContext(goCtx), msg.OwnerId, msg.DestinationId, msg.AccessClientList)

	return &types.MsgCreateClientRegistrationRelationshipResponse{}, nil
}
