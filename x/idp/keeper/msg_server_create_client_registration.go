package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistration(goCtx context.Context, msg *types.MsgCreateClientRegistrationRequest) (result *types.MsgCreateClientRegistrationResponse, err error) {
	k.Keeper.SetClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistrationRegistryW3CIdentifier, msg.ClientRegistration)

	return result, nil
}
