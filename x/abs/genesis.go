package abs

import (
	"github.com/be-heroes/doxchain/x/abs/keeper"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the partitionedPools
	for _, elem := range genState.PartitionedPoolRegistries {
		k.SetPartitionedPoolRegistry(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PartitionedPoolRegistries = k.GetAllPartitionedPoolRegistries(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
