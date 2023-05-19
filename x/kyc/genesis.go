package kyc

import (
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if genState.RegistrationList != nil {
		for _, elem := range genState.RegistrationList {
			k.SetKYCRegistration(ctx, elem)
		}
	}
	k.SetParams(ctx, genState.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	registrations := k.GetAllKYCRegistration(ctx)

	if registrations != nil {
		genesis.RegistrationList = registrations
	}

	return genesis
}
