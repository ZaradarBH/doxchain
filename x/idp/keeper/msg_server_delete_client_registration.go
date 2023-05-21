package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteClientRegistration(goCtx context.Context, msg *types.MsgDeleteClientRegistrationRequest) (*types.MsgDeleteClientRegistrationResponse, error) {
	if msg.Creator != msg.ClientRegistration.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only the owner of a client registration can delete it")
	}

	k.Keeper.RemoveClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistration.Id.GetW3CIdentifier())

	return &types.MsgDeleteClientRegistrationResponse{}, nil
}
