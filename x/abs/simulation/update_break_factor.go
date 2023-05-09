package simulation

import (
	"math/rand"

	"doxchain/x/abs/keeper"
	"doxchain/x/abs/types"
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
		msg := &types.MsgUpdateBreakFactor{
			Creator: simAccount.Address.String(),
		}

		breakFactor, err := sdk.NewDecFromStr(msg.Value)

		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "unable to parse breakFactor sdk.Dec from string"), nil, err
		}

		k.SetBreakFactor(ctx, breakFactor)

		breakFactorFromKeeper := k.GetBreakFactor(ctx)

		if !breakFactor.Equal(breakFactorFromKeeper) {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Keeper BreakFactor setter input does not match getter output"), nil, nil
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}
