package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/oauth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Token(goCtx context.Context, msg *types.MsgTokenRequest) (*types.MsgTokenResponse, error) {
	return k.Token(sdk.UnwrapSDKContext(goCtx), msg)
}
