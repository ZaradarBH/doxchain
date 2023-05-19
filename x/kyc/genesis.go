package kyc

import (
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.RegistrationList != nil {
		for _, elem := range genState.RegistrationList {
			k.SetKYCRegistration(ctx, elem)
		}
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	registrations := k.GetAllKYCRegistration(ctx)

	if registrations != nil {
		genesis.RegistrationList = registrations
	}

	return genesis
}
