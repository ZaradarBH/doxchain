package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistrationRelationship(goCtx context.Context, msg *types.MsgCreateClientRegistrationRelationshipRequest) (*types.MsgCreateClientRegistrationRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateClientRegistrationRelationshipResponse{}, nil
}
