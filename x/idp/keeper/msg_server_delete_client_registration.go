package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistration(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRequest) (*types.MsgDeleteClientRegistrationResponse, error) {
	//TODO: Fetch client registration and check if the creator is allowed to delete it	
	k.Keeper.RemoveClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationW3CIdentifier)

	return &types.MsgDeleteClientRegistrationResponse{}, nil
}
