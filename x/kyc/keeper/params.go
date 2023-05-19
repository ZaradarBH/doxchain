package keeper

import (
	"github.com/be-heroes/doxchain/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
