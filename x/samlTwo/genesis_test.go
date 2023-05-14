package samltwo_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/samltwo"
	"github.com/be-heroes/doxchain/x/samltwo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.samltwoKeeper(t)
	saml.InitGenesis(ctx, *k, genesisState)
	got := saml.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
