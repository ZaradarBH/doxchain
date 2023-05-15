package kyc

import (
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.RequestList != nil {
		for _, elem := range genState.RequestList {
			k.SetKYCRequest(ctx, elem)
		}
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all requests
	requests := k.GetAllKYCRequest(ctx)
	if requests != nil {
		genesis.RequestList = requests
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
