package abs

import (
	"github.com/be-heroes/doxchain/x/abs/keeper"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.PartitionedPoolRegistries {
		k.SetPartitionedPoolRegistry(ctx, elem)
	}
	
	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.PartitionedPoolRegistries = k.GetAllPartitionedPoolRegistries(ctx)

	return genesis
}
