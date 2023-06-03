package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistration(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRequest) (result *types.MsgDeleteClientRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	registration, found := k.Keeper.GetClientRegistration(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationW3CIdentifier)

	if found {
		if registration.Id.Creator != msg.Creator {
			return nil, types.ErrImpersonation
		}

		k.Keeper.RemoveClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationW3CIdentifier)
	}

	result = &types.MsgDeleteClientRegistrationResponse{}

	return result, nil
}
