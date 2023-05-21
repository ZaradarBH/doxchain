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

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))

	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}

	idpGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ClientRegistrationRegistries: []types.ClientRegistrationRegistry{
			{
				Owner: *utils.NewDidTokenFactory().Create(sample.AccAddress(), "did:methodname:methodid"),
			},
			{
				Owner: *utils.NewDidTokenFactory().Create(sample.AccAddress(), "did:methodname:methodid"),
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

	return operations
}
