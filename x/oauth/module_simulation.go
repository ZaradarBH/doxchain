package oauth

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	oauthsimulation "github.com/be-heroes/doxchain/x/oauth/simulation"
	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oauthsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgTokenRequest = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTokenRequest int = 100

	opWeightMsgCreateDeviceCodes = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeviceCodes int = 100

	opWeightMsgUpdateDeviceCodes = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeviceCodes int = 100

	opWeightMsgDeleteDeviceCodes = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeviceCodes int = 100

	opWeightMsgCreateAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccessTokenRegistry int = 100

	opWeightMsgUpdateAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccessTokenRegistry int = 100

	opWeightMsgDeleteAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAccessTokenRegistry int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oauthGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DeviceCodesList: []types.DeviceCodes{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		AccessTokenRegistryList: []types.AccessTokenRegistry{
			{
				Tenant: "0",
			},
			{
				Tenant: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oauthGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgTokenRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTokenRequest, &weightMsgTokenRequest, nil,
		func(_ *rand.Rand) {
			weightMsgTokenRequest = defaultWeightMsgTokenRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTokenRequest,
		oauthsimulation.SimulateMsgTokenRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
