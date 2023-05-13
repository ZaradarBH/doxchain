package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/oauth/keeper"
	"github.com/be-heroes/doxchain/x/oauth/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccessTokenRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccessTokenRegistry {
	items := make([]types.AccessTokenRegistry, n)
	for i := range items {
		items[i].jti = strconv.Itoa(i)

		keeper.SetAccessTokenRegistry(ctx, items[i])
	}
	return items
}

func TestAccessTokenRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokenRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccessTokenRegistry(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccessTokenRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokenRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccessTokenRegistry(ctx,
			item.Index,
		)
		_, found := keeper.GetAccessTokenRegistry(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAccessTokenRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokenRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccessTokenRegistry(ctx)),
	)
}
