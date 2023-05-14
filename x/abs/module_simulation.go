package abs

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	abssimulation "github.com/be-heroes/doxchain/x/abs/simulation"
	"github.com/be-heroes/doxchain/x/abs/types"
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
	_ = abssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUpdateBreakFactor = "op_weight_msg_update_break_factor"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBreakFactor int = 100

	opWeightMsgCreatePartitionedPools = "op_weight_msg_partitioned_pools"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePartitionedPools int = 100

	opWeightMsgUpdatePartitionedPools = "op_weight_msg_partitioned_pools"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePartitionedPools int = 100

	opWeightMsgDeletePartitionedPools = "op_weight_msg_partitioned_pools"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePartitionedPools int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	absGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PartitionedPoolsList: []types.PartitionedPools{
		{
			Creator: sample.AccAddress(),
Index: "0",
},
		{
			Creator: sample.AccAddress(),
Index: "1",
},
	},
	// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&absGenesis)
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

	var weightMsgUpdateBreakFactor int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBreakFactor, &weightMsgUpdateBreakFactor, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBreakFactor = defaultWeightMsgUpdateBreakFactor
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBreakFactor,
		abssimulation.SimulateMsgUpdateBreakFactor(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePartitionedPools int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePartitionedPools, &weightMsgCreatePartitionedPools, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePartitionedPools = defaultWeightMsgCreatePartitionedPools
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePartitionedPools,
		abssimulation.SimulateMsgCreatePartitionedPools(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePartitionedPools int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePartitionedPools, &weightMsgUpdatePartitionedPools, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePartitionedPools = defaultWeightMsgUpdatePartitionedPools
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePartitionedPools,
		abssimulation.SimulateMsgUpdatePartitionedPools(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePartitionedPools int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePartitionedPools, &weightMsgDeletePartitionedPools, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePartitionedPools = defaultWeightMsgDeletePartitionedPools
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePartitionedPools,
		abssimulation.SimulateMsgDeletePartitionedPools(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
