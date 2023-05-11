package did

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	didsimulation "github.com/be-heroes/doxchain/x/did/simulation"
	"github.com/be-heroes/doxchain/x/did/types"
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
	_ = didsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDid = "op_weight_msg_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDid int = 100

	opWeightMsgUpdateDid = "op_weight_msg_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDid int = 100

	opWeightMsgDeleteDid = "op_weight_msg_did"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	didGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DidList: []types.Did{
			{
				Creator:    sample.AccAddress(),
				Url:        "did:example:123/path?service=agent#degree",
				MethodName: "example",
				MethodId:   "123",
				Path:       "path",
				Fragment:   "degree",
				Query:      "service=agent",
			},
			{
				Creator:    sample.AccAddress(),
				Url:        "did:example:456/path?service=agent#degree",
				MethodName: "example",
				MethodId:   "456",
				Path:       "path",
				Fragment:   "degree",
				Query:      "service=agent",
			},
		},
		DidCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&didGenesis)
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

	var weightMsgCreateDid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDid, &weightMsgCreateDid, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDid = defaultWeightMsgCreateDid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDid,
		didsimulation.SimulateMsgCreateDid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDid, &weightMsgUpdateDid, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDid = defaultWeightMsgUpdateDid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDid,
		didsimulation.SimulateMsgUpdateDid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDid, &weightMsgDeleteDid, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDid = defaultWeightMsgDeleteDid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDid,
		didsimulation.SimulateMsgDeleteDid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
