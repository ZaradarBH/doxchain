package doxchain_test

import (
	"testing"

	keepertest "doxchain/testutil/keeper"
	"doxchain/testutil/nullify"
	"doxchain/x/doxchain"
	"doxchain/x/doxchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DoxchainKeeper(t)
	doxchain.InitGenesis(ctx, *k, genesisState)
	got := doxchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
