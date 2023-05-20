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
	opWeightMsgCreateAMLRegistration = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAMLRegistration int = 100

	opWeightMsgUpdateAMLRegistration = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAMLRegistration int = 100

	opWeightMsgDeleteAMLRegistration = "op_weight_msg_aml_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAMLRegistration int = 100

	opWeightMsgApproveRequest = "op_weight_msg_approve_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveRequest int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	amlGenesis := types.GenesisState{
		Params: types.DefaultParams(),
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

	var weightMsgCreateAMLRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAMLRegistration, &weightMsgCreateAMLRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAMLRegistration = defaultWeightMsgCreateAMLRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAMLRegistration,
		amlsimulation.SimulateMsgCreateAMLRegistration(am.keeper),
	))

	var weightMsgDeleteAMLRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAMLRegistration, &weightMsgDeleteAMLRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAMLRegistration = defaultWeightMsgDeleteAMLRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAMLRegistration,
		amlsimulation.SimulateMsgDeleteAMLRegistration(am.keeper),
	))

	var weightMsgApproveRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveRequest, &weightMsgApproveRequest, nil,
		func(_ *rand.Rand) {
			weightMsgApproveRequest = defaultWeightMsgApproveRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveRequest,
		amlsimulation.SimulateMsgApproveRequest(am.keeper),
	))

	return operations
}
