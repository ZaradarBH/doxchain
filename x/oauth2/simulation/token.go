package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/oauth2/keeper"
	"github.com/be-heroes/doxchain/x/oauth2/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgTokenRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTokenRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Token simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Token simulation not implemented"), nil, nil
	}
}
