package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
)

func createTestKYCRequest(keeper *keeper.Keeper, ctx sdk.Context) types.KYCRequest {
	item := types.KYCRequest{}
	keeper.SetKYCRequest(ctx, item)
	return item
}

func TestKYCRequestGet(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	item := createTestKYCRequest(keeper, ctx)
	rst, found := keeper.GetKYCRequest(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestKYCRequestRemove(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	createTestKYCRequest(keeper, ctx)
	keeper.RemoveKYCRequest(ctx)
	_, found := keeper.GetKYCRequest(ctx)
	require.False(t, found)
}
