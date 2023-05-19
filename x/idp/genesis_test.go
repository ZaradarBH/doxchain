package idp_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/idp"
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClientRegistrationRegistryList: []types.ClientRegistrationRegistry{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdpKeeper(t)
	idp.InitGenesis(ctx, *k, genesisState)
	got := idp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClientRegistrationRegistryList, got.ClientRegistrationRegistryList)
	// this line is used by starport scaffolding # genesis/test/assert
}
