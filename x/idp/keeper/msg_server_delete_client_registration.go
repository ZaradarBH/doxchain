package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteClientRegistration(goCtx context.Context, msg *types.MsgDeleteClientRegistration) (*types.MsgDeleteClientRegistrationResponse, error) {
	k.Keeper.RemoveClientRegistration(sdk.UnwrapSDKContext(goCtx), msg.Creator, msg.Name)

	return &types.MsgDeleteClientRegistrationResponse{}, nil
}
