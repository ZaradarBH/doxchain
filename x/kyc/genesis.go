package kyc

import (
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.KYCRequest != nil {
		k.SetKYCRequest(ctx, *genState.KYCRequest)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all kYCRequest
	kYCRequest, found := k.GetKYCRequest(ctx)
	if found {
		genesis.KYCRequest = &kYCRequest
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
