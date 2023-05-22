package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/aml/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveAMLRegistration(goCtx context.Context, msg *types.MsgApproveAMLRegistrationRequest) (*types.MsgApproveAMLRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var approvers []didTypes.Did

	k.Keeper.paramstore.Get(ctx, types.ParamStoreKeyApprovers, &approvers)

	for _, approverId := range approvers {
		if approverId.Creator == msg.Creator {
			k.Keeper.ApproveAMLRegistration(ctx, msg.Target.Creator)

			break
		}
	}

	return &types.MsgApproveAMLRegistrationResponse{}, nil
}
