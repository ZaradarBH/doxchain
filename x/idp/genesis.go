package idp

import (
	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the DeviceCodeRegistry
	for _, elem := range genState.DeviceCodeRegistries {
		k.SetDeviceCodeRegistry(ctx, elem)
	}
	// Set all the ClientRegistry
	for _, elem := range genState.ClientRegistries {
		k.SetClientRegistry(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	
	genesis.DeviceCodeRegistries = k.GetAllDeviceCodeRegistry(ctx)
	genesis.ClientRegistries = k.GetAllClientRegistry(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
