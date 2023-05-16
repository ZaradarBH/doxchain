package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateClientRegistration(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUpdateClientRegistration{}

		// TODO: Handling the UpdateClientRegistration simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UpdateClientRegistration simulation not implemented"), nil, nil
	}
}
