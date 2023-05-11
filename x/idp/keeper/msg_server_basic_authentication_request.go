package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BasicAuthenticationRequest(goCtx context.Context, msg *types.MsgBasicAuthenticationRequest) (*types.MsgBasicAuthenticationRequestResponse, error) {
	return k.Login(sdk.UnwrapSDKContext(goCtx), msg);
}
