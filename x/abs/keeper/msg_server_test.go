package keeper_test

import (
	"context"
	"testing"

	keepertest "doxchain/testutil/keeper"
	"doxchain/x/abs/keeper"
	"doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AbsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
