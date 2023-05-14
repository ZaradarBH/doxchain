package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/oauthtwo/keeper"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.oauthtwoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
