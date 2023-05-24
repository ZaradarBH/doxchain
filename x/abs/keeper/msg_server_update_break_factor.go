package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateBreakFactor(goCtx context.Context, msg *types.MsgUpdateBreakFactorRequest) (result *types.MsgUpdateBreakFactorResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var operators []didTypes.Did

	k.Keeper.paramstore.Get(ctx, types.ParamStoreKeyOperators, &operators)

	for _, operatorId := range operators {
		if operatorId.Creator == msg.Creator {
			if !msg.BreakFactor.IsNegative() && !msg.BreakFactor.GT(sdk.OneDec()) { 
				k.Keeper.SetBreakFactor(ctx, msg.BreakFactor)
			} else {
				return nil, sdkerrors.Wrap(types.ErrBreakFactorOutOfBounds, msg.BreakFactor.String())
			}

			break
		}
	}

	return result, nil
}
