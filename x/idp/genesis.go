package idp

import (
	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.DeviceCodeRegistries {
		k.SetDeviceCodeRegistry(ctx, elem)
	}
	
	for _, elem := range genState.ClientRegistrationRegistries {
		k.SetClientRegistrationRegistry(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.DeviceCodeRegistries = k.GetAllDeviceCodeRegistry(ctx)
	genesis.ClientRegistrationRegistries = k.GetAllClientRegistrationRegistry(ctx)

	return genesis
}
