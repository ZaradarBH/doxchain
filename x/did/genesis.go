package did

import (
	"github.com/be-heroes/doxchain/x/did/keeper"
	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.DidList {
		k.SetDid(ctx, elem, true)
	}

	k.SetDidCount(ctx, genState.DidCount)
	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DidList = k.GetAllDid(ctx)
	genesis.DidCount = k.GetDidCount(ctx)

	return genesis
}
