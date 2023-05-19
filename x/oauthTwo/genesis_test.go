package oauthtwo_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/oauthtwo"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DeviceCodeRegistries: []types.DeviceCodeRegistry{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		AccessTokenRegistries: []types.AccessTokenRegistry{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		AuthorizationCodeRegistries: []types.AuthorizationCodeRegistry{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.oauthtwoKeeper(t)
	oauth.InitGenesis(ctx, *k, genesisState)
	got := oauth.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DeviceCodeRegistries, got.DeviceCodeRegistries)
	require.ElementsMatch(t, genesisState.AccessTokenRegistries, got.AccessTokenRegistries)
	require.ElementsMatch(t, genesisState.AuthorizationCodeRegistries, got.AuthorizationCodeRegistries)
	// this line is used by starport scaffolding # genesis/test/assert
}
