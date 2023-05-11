package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgAuthenticationRequest) (*types.MsgAuthenticationResponse, error) {
	return k.Login(sdk.UnwrapSDKContext(goCtx), msg)
}
