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

var _ = strconv.IntSize

func createNClientRegistrationRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ClientRegistrationRegistry {
	items := make([]types.ClientRegistrationRegistry, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClientRegistrationRegistry(ctx, items[i])
	}
	return items
}

func TestClientRegistrationRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistrationRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClientRegistrationRegistry(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClientRegistrationRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistrationRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClientRegistrationRegistry(ctx,
			item.Index,
		)
		_, found := keeper.GetClientRegistrationRegistry(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClientRegistrationRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.IdpKeeper(t)
	items := createNClientRegistrationRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClientRegistrationRegistry(ctx)),
	)
}