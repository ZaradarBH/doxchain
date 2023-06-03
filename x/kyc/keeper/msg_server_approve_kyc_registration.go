package keeper

import (
	"context"

	didTypes "github.com/be-heroes/doxchain/x/did/types"
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveKYCRegistration(goCtx context.Context, msg *types.MsgApproveKYCRegistrationRequest) (result *types.MsgApproveKYCRegistrationResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var approvers []didTypes.Did

	k.Keeper.paramstore.Get(ctx, types.ParamStoreKeyApprovers, &approvers)

	for _, approverId := range approvers {
		if approverId.Creator == msg.Creator {
			k.Keeper.ApproveKYCRegistration(ctx, msg.Target.Creator)

			break
		}
	}

	result = &types.MsgApproveKYCRegistrationResponse{}

	return result, nil
}
