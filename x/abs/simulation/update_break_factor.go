package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/abs/keeper"
	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUpdateBreakFactor(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUpdateBreakFactorRequest{
			Creator: simAccount.Address.String(),
		}

		k.SetBreakFactor(ctx, msg.BreakFactor)

		breakFactorFromKeeper := k.GetBreakFactor(ctx)

		if !msg.BreakFactor.Equal(breakFactorFromKeeper) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Keeper BreakFactor setter input does not match getter output"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
