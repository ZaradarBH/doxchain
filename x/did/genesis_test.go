package did_test

import (
	"testing"

	keepertest "doxchain/testutil/keeper"
	"doxchain/testutil/nullify"
	"doxchain/x/did"
	"doxchain/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DidList: []types.Did{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DidCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DidKeeper(t)
	did.InitGenesis(ctx, *k, genesisState)
	got := did.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DidList, got.DidList)
	require.Equal(t, genesisState.DidCount, got.DidCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
