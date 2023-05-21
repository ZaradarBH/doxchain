package oauthtwo

import (
	"math/rand"

	"github.com/be-heroes/doxchain/testutil/sample"
	oauthtwosimulation "github.com/be-heroes/doxchain/x/oauthtwo/simulation"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	didUtils "github.com/be-heroes/doxchain/utils/did"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oauthtwosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgTokenRequest = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTokenRequest int = 100

	opWeightMsgCreateDeviceCodeRegistry = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDeviceCodeRegistry int = 100

	opWeightMsgUpdateDeviceCodeRegistry = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeviceCodeRegistry int = 100

	opWeightMsgDeleteDeviceCodeRegistry = "op_weight_msg_device_codes"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDeviceCodeRegistry int = 100

	opWeightMsgCreateAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAccessTokenRegistry int = 100

	opWeightMsgUpdateAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccessTokenRegistry int = 100

	opWeightMsgDeleteAccessTokenRegistry = "op_weight_msg_access_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAccessTokenRegistry int = 100

	opWeightMsgDeviceCode = "op_weight_msg_device_code"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeviceCode int = 100

	opWeightMsgAuthorize = "op_weight_msg_authorize"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuthorize int = 100

	opWeightMsgCreateAuthorizationCodeRegistry = "op_weight_msg_authorization_code_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAuthorizationCodeRegistry int = 100

	opWeightMsgUpdateAuthorizationCodeRegistry = "op_weight_msg_authorization_code_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAuthorizationCodeRegistry int = 100

	opWeightMsgDeleteAuthorizationCodeRegistry = "op_weight_msg_authorization_code_registry"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAuthorizationCodeRegistry int = 100

)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oauthGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AccessTokenRegistries: []types.AccessTokenRegistry{
			{
				Owner: *didUtils.NewDidTokenFactory().Create("1", "did:methodname:methodid"),
			},
			{
				Owner: *didUtils.NewDidTokenFactory().Create("2", "did:methodname:methodid"),
			},
		},
		AuthorizationCodeRegistries: []types.AuthorizationCodeRegistry{
			{
				Owner: *didUtils.NewDidTokenFactory().Create("1", "did:methodname:methodid"),
			},
			{
				Owner: *didUtils.NewDidTokenFactory().Create("2", "did:methodname:methodid"),
			},
		},
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

	return operations
}
