package kyc

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	kycsimulation "github.com/be-heroes/doxchain/x/kyc/simulation"
	"github.com/be-heroes/doxchain/x/kyc/types"
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
	_ = kycsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateKYCRegistration = "op_weight_msg_kyc_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKYCRegistration int = 100

	opWeightMsgUpdateKYCRegistration = "op_weight_msg_kyc_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKYCRegistration int = 100

	opWeightMsgDeleteKYCRegistration = "op_weight_msg_kyc_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKYCRegistration int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kycGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kycGenesis)
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

	var weightMsgCreateKYCRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKYCRegistration, &weightMsgCreateKYCRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKYCRegistration = defaultWeightMsgCreateKYCRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKYCRegistration,
		kycsimulation.SimulateMsgCreateKYCRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKYCRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKYCRegistration, &weightMsgDeleteKYCRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKYCRegistration = defaultWeightMsgDeleteKYCRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKYCRegistration,
		kycsimulation.SimulateMsgDeleteKYCRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
