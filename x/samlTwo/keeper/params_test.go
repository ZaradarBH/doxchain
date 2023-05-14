package keeper_test

import (
	"testing"

	testkeeper "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/samltwo/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.samltwoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
