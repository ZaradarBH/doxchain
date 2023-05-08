package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "doxchain/testutil/keeper"
	"doxchain/x/doxchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DoxchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
