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

func createNAccessTokens(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccessTokens {
	items := make([]types.AccessTokens, n)
	for i := range items {
		items[i].jti = strconv.Itoa(i)

		keeper.SetAccessTokens(ctx, items[i])
	}
	return items
}

func TestAccessTokensGet(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokens(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccessTokens(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccessTokensRemove(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokens(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccessTokens(ctx,
			item.Index,
		)
		_, found := keeper.GetAccessTokens(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAccessTokensGetAll(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNAccessTokens(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccessTokens(ctx)),
	)
}
