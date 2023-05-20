package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/aml/keeper"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgApproveRequest(
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgApproveAMLRegistrationRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ApproveAMLRegistrationRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ApproveAMLRegistrationRequest simulation not implemented"), nil, nil
	}
}
