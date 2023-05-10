package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/abs/keeper"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AbsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
