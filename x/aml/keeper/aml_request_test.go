package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/aml/keeper"
	"github.com/be-heroes/doxchain/x/aml/types"
)

func createTestAMLRequest(keeper *keeper.Keeper, ctx sdk.Context) types.AMLRequest {
	item := types.AMLRequest{}
	keeper.SetAMLRequest(ctx, item)
	return item
}

func TestAMLRequestGet(t *testing.T) {
	keeper, ctx := keepertest.AmlKeeper(t)
	item := createTestAMLRequest(keeper, ctx)
	rst, found := keeper.GetAMLRequest(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAMLRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.AmlKeeper(t)
	createTestAMLRequest(keeper, ctx)
	keeper.RemoveAMLRequest(ctx)
	_, found := keeper.GetAMLRequest(ctx)
	require.False(t, found)
}
