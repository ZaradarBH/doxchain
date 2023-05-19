package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateBreakFactor(goCtx context.Context, msg *types.MsgUpdateBreakFactorRequest) (*types.MsgUpdateBreakFactorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	var operators []didTypes.Did

	k.Keeper.paramstore.Get(ctx, types.ParamStoreKeyOperators, &operators)

	for _, operatorId := range operators {
		if operatorId.Creator == msg.Creator {
			k.Keeper.SetBreakFactor(ctx, msg.BreakFactor)

			break;
		}
	}

	return &types.MsgUpdateBreakFactorResponse{}, nil
}
