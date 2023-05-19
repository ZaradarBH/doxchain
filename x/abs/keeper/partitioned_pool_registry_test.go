package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/abs/keeper"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

var _ = strconv.IntSize

func createNPartitionedPools(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PartitionedPools {
	items := make([]types.PartitionedPools, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPartitionedPools(ctx, items[i])
	}
	return items
}

func TestPartitionedPoolsGet(t *testing.T) {
	keeper, ctx := keepertest.AbsKeeper(t)
	items := createNPartitionedPools(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPartitionedPools(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPartitionedPoolsRemove(t *testing.T) {
	keeper, ctx := keepertest.AbsKeeper(t)
	items := createNPartitionedPools(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePartitionedPools(ctx,
			item.Index,
		)
		_, found := keeper.GetPartitionedPools(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPartitionedPoolsGetAll(t *testing.T) {
	keeper, ctx := keepertest.AbsKeeper(t)
	items := createNPartitionedPools(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPartitionedPools(ctx)),
	)
}
