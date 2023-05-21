package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateClientRegistration(goCtx context.Context, msg *types.MsgUpdateClientRegistrationRequest) (*types.MsgUpdateClientRegistrationResponse, error) {
	k.Keeper.SetClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistration)

	return &types.MsgUpdateClientRegistrationResponse{}, nil
}
