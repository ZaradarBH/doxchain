package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/oauthTwo/keeper"
	"github.com/be-heroes/doxchain/x/oauthTwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDeviceCodeRegistry(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DeviceCodeRegistry {
	items := make([]types.DeviceCodeRegistry, n)
	for i := range items {
		items[i].Tenant = strconv.Itoa(i)

		keeper.SetDeviceCodeRegistry(ctx, items[i])
	}
	return items
}

func TestDeviceCodeRegistryGet(t *testing.T) {
	keeper, ctx := keepertest.OauthTwoKeeper(t)
	items := createNDeviceCodeRegistry(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeviceCodeRegistry(ctx,
			item.Tenant,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceCodeRegistryRemove(t *testing.T) {
	keeper, ctx := keepertest.OauthTwoKeeper(t)
	items := createNDeviceCodeRegistry(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeviceCodeRegistry(ctx,
			item.Tenant,
		)
		_, found := keeper.GetDeviceCodeRegistry(ctx,
			item.Tenant,
		)
		require.False(t, found)
	}
}

func TestDeviceCodeRegistryGetAll(t *testing.T) {
	keeper, ctx := keepertest.OauthTwoKeeper(t)
	items := createNDeviceCodeRegistry(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeviceCodeRegistry(ctx)),
	)
}
