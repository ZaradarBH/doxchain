package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistrationRelationship(goCtx context.Context, msg *types.MsgCreateClientRegistrationRelationshipRequest) (result *types.MsgCreateClientRegistrationRelationshipResponse, err error) {
	k.Keeper.SetClientRegistrationRelationship(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationRelationshipRegistryEntry)

	return result, nil
}
