package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistrationRelationship(goCtx context.Context, msg *types.MsgCreateClientRegistrationRelationshipRequest) (*types.MsgCreateClientRegistrationRelationshipResponse, error) {
	err := k.Keeper.SetClientRegistrationRelationship(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRelationshipRegistryEntry)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateClientRegistrationRelationshipResponse{}, nil
}
