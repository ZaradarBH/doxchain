package keeper

import (
	"cosmossdk.io/math"
	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)

	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetThrottledRollingAverage(ctx sdk.Context) (res math.Int) {
	k.paramstore.Get(ctx, types.ParamStoreKeyThrottledRollingAverageKey, &res)
	return
}

func (k Keeper) GetBlockExpireOffset(ctx sdk.Context) (res math.Int) {
	k.paramstore.Get(ctx, types.ParamStoreKeyBlockExpireOffset, &res)
	return
}

func (k Keeper) SetBlockExpireOffset(ctx sdk.Context, offset math.Int) {
	k.paramstore.Set(ctx, types.ParamStoreKeyBlockExpireOffset, offset)
}
