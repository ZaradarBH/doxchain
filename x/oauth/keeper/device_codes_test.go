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

func createNDeviceCodes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DeviceCodes {
	items := make([]types.DeviceCodes, n)
	for i := range items {
		items[i].Tenant = strconv.Itoa(i)

		keeper.SetDeviceCodes(ctx, items[i])
	}
	return items
}

func TestDeviceCodesGet(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNDeviceCodes(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeviceCodes(ctx,
			item.Tenant,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeviceCodesRemove(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNDeviceCodes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeviceCodes(ctx,
			item.Tenant,
		)
		_, found := keeper.GetDeviceCodes(ctx,
			item.Tenant,
		)
		require.False(t, found)
	}
}

func TestDeviceCodesGetAll(t *testing.T) {
	keeper, ctx := keepertest.OauthKeeper(t)
	items := createNDeviceCodes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeviceCodes(ctx)),
	)
}
