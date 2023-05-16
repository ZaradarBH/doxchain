package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNClientRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClientRegistry {
	items := make([]types.ClientRegistry, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClientRegistry(ctx, items[i])
	}
	return items
}

func TestClientRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClientRegistry(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClientRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClientRegistry(ctx,
			item.Index,
		)
		_, found := keeper.GetClientRegistry(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClientRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClientRegistry(ctx)),
	)
}
