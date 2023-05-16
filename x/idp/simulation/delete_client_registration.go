package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgDeleteClientRegistration(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDeleteClientRegistration{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DeleteClientRegistration simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DeleteClientRegistration simulation not implemented"), nil, nil
	}
}
