package oauthtwo

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/keeper"
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.AccessTokenRegistries {
		k.SetAccessTokenRegistry(ctx, elem)
	}

	for _, elem := range genState.AuthorizationCodeRegistries {
		k.SetAuthorizationCodeRegistry(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.AccessTokenRegistries = k.GetAllAccessTokenRegistry(ctx)
	genesis.AuthorizationCodeRegistries = k.GetAllAuthorizationCodeRegistry(ctx)

	return genesis
}
