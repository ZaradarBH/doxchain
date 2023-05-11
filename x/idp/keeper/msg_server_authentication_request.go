package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AuthenticationRequest(goCtx context.Context, msg *types.MsgAuthenticationRequest) (*types.MsgAuthenticationResponse, error) {
	response, err := k.Login(sdk.UnwrapSDKContext(goCtx), *msg)

	return &response, err
}
