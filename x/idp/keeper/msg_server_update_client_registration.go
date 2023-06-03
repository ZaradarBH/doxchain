package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateClientRegistration(goCtx context.Context, msg *types.MsgUpdateClientRegistrationRequest) (result *types.MsgUpdateClientRegistrationResponse, err error) {
	k.Keeper.SetClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistration)

	result = &types.MsgUpdateClientRegistrationResponse{}

	return result, nil
}
