package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/oauthtwo/keeper"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAuthorizationCodeRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AuthorizationCodeRegistry {
	items := make([]types.AuthorizationCodeRegistry, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAuthorizationCodeRegistry(ctx, items[i])
	}
	return items
}

func TestAuthorizationCodeRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.OauthtwoKeeper(t)
	items := createNAuthorizationCodeRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAuthorizationCodeRegistry(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAuthorizationCodeRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.OauthtwoKeeper(t)
	items := createNAuthorizationCodeRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAuthorizationCodeRegistry(ctx,
			item.Index,
		)
		_, found := keeper.GetAuthorizationCodeRegistry(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAuthorizationCodeRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.OauthtwoKeeper(t)
	items := createNAuthorizationCodeRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAuthorizationCodeRegistry(ctx)),
	)
}
