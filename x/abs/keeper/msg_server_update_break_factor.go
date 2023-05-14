package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateBreakFactor(goCtx context.Context, msg *types.MsgUpdateBreakFactorRequest) (*types.MsgUpdateBreakFactorResponse, error) {
	decValue, err := sdk.NewDecFromStr(msg.Value)

	if err != nil {
		return nil, err
	}

	err = k.Keeper.SetBreakFactor(sdk.UnwrapSDKContext(goCtx), decValue)

	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateBreakFactorResponse{}, nil
}
