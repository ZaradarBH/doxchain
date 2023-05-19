package idp

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	idpsimulation "github.com/be-heroes/doxchain/x/idp/simulation"
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	utils "github.com/be-heroes/doxchain/utils/did"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = idpsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAuthenticationRequest = "op_weight_msg_basic_authentication_request"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuthenticationRequest int = 100

	opWeightMsgCreateClientRegistrationRegistry = "op_weight_msg_client_registration_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClientRegistrationRegistry int = 100

	opWeightMsgUpdateClientRegistrationRegistry = "op_weight_msg_client_registration_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClientRegistrationRegistry int = 100

	opWeightMsgDeleteClientRegistrationRegistry = "op_weight_msg_client_registration_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteClientRegistrationRegistry int = 100

	opWeightMsgCreateClientRegistration = "op_weight_msg_create_client_registration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClientRegistration int = 100

	opWeightMsgUpdateClientRegistration = "op_weight_msg_update_client_registration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClientRegistration int = 100

	opWeightMsgDeleteClientRegistration = "op_weight_msg_delete_client_registration"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteClientRegistration int = 100

	opWeightMsgCreateClientRegistrationRelationship = "op_weight_msg_create_client_registration_relationship"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClientRegistrationRelationship int = 100
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	idpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ClientRegistrationRegistries: []types.ClientRegistrationRegistry{
			{
				Owner: *utils.NewDidTokenFactory().Create(sample.AccAddress(), ""),
			},
			{
				Owner: *utils.NewDidTokenFactory().Create(sample.AccAddress(), ""),
			},
		},
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&idpGenesis)
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

	var weightMsgAuthenticationRequest int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAuthenticationRequest, &weightMsgAuthenticationRequest, nil,
		func(_ *rand.Rand) {
			weightMsgAuthenticationRequest = defaultWeightMsgAuthenticationRequest
		},
	)

	var weightMsgCreateClientRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClientRegistration, &weightMsgCreateClientRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClientRegistration = defaultWeightMsgCreateClientRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClientRegistration,
		idpsimulation.SimulateMsgCreateClientRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateClientRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateClientRegistration, &weightMsgUpdateClientRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClientRegistration = defaultWeightMsgUpdateClientRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClientRegistration,
		idpsimulation.SimulateMsgUpdateClientRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteClientRegistration int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteClientRegistration, &weightMsgDeleteClientRegistration, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteClientRegistration = defaultWeightMsgDeleteClientRegistration
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteClientRegistration,
		idpsimulation.SimulateMsgDeleteClientRegistration(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateClientRegistrationRelationship int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateClientRegistrationRelationship, &weightMsgCreateClientRegistrationRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClientRegistrationRelationship = defaultWeightMsgCreateClientRegistrationRelationship
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClientRegistrationRelationship,
		idpsimulation.SimulateMsgCreateClientRegistrationRelationship(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	return operations
}
