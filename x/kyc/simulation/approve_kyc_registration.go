package simulation

import (
	"math/rand"

	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
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
		msg := &types.MsgApproveKYCRegistrationRequest{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ApproveRequest simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ApproveRequest simulation not implemented"), nil, nil
	}
}
