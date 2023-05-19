package aml_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/aml"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AMLRegistration: &types.AMLRegistration{
			FirstName: "91",
			LastName:  "47",
			Approved:  true,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AmlKeeper(t)
	aml.InitGenesis(ctx, *k, genesisState)
	got := aml.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.AMLRegistration, got.AMLRegistration)
	// this line is used by starport scaffolding # genesis/test/assert
}
