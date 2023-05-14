package aml

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	amlsimulation "github.com/be-heroes/doxchain/x/aml/simulation"
	"github.com/be-heroes/doxchain/x/aml/types"
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
	_ = amlsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAMLRequest = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAMLRequest int = 100

	opWeightMsgUpdateAMLRequest = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAMLRequest int = 100

	opWeightMsgDeleteAMLRequest = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAMLRequest int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	amlGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&amlGenesis)
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

	var weightMsgCreateAMLRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAMLRequest, &weightMsgCreateAMLRequest, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAMLRequest = defaultWeightMsgCreateAMLRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAMLRequest,
		amlsimulation.SimulateMsgCreateAMLRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAMLRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAMLRequest, &weightMsgDeleteAMLRequest, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAMLRequest = defaultWeightMsgDeleteAMLRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAMLRequest,
		amlsimulation.SimulateMsgDeleteAMLRequest(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
