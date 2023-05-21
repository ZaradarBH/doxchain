package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteClientRegistration(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRequest) (*types.MsgDeleteClientRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	result, found := k.Keeper.GetClientRegistration(ctx, msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationW3CIdentifier)
	
	if found {		
		if result.Id.Creator != msg.Creator {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid creator")
		}

		k.Keeper.RemoveClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistrationW3CIdentifier)
	}

	return &types.MsgDeleteClientRegistrationResponse{}, nil
}
