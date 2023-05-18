package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateClientRegistration(goCtx context.Context, msg *types.MsgCreateClientRegistration) (*types.MsgCreateClientRegistrationResponse, error) {
	k.Keeper.SetClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.ClientRegistration)

	return &types.MsgCreateClientRegistrationResponse{}, nil
}
