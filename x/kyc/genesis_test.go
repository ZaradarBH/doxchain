package kyc_test

import (
	"testing"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/kyc"
	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KYCRegistration: &types.KYCRegistration{
			FirstName: "56",
			LastName:  "3",
			Approved:  true,
		},
	}

	k, ctx := keepertest.KycKeeper(t)
	kyc.InitGenesis(ctx, *k, genesisState)
	got := kyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.KYCRegistration, got.KYCRegistration)
}
