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

func createTestAMLRegistration(keeper *keeper.Keeper, ctx sdk.Context) types.AMLRegistration {
	item := types.AMLRegistration{}
	keeper.SetAMLRegistration(ctx, item)
	return item
}

func TestAMLRegistrationGet(t *testing.T) {
	keeper, ctx := keepertest.AmlKeeper(t)
	item := createTestAMLRegistration(keeper, ctx)
	rst, found := keeper.GetAMLRegistration(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestAMLRegistrationRemove(t *testing.T) {
	keeper, ctx := keepertest.AmlKeeper(t)
	createTestAMLRegistration(keeper, ctx)
	keeper.RemoveAMLRegistration(ctx)
	_, found := keeper.GetAMLRegistration(ctx)
	require.False(t, found)
}
