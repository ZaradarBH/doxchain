package types_test

import (
	"testing"

	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated DeviceCodeRegistry",
			genState: &types.GenesisState{
				DeviceCodeRegistries: []types.DeviceCodeRegistry{
					{
						Tenant: "0",
					},
					{
						Tenant: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated AccessTokenRegistry",
			genState: &types.GenesisState{
				AccessTokenRegistries: []types.AccessTokenRegistry{
					{
						Tenant: "0",
					},
					{
						Tenant: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated authorizationCodeRegistry",
			genState: &types.GenesisState{
				AuthorizationCodeRegistries: []types.AuthorizationCodeRegistry{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
