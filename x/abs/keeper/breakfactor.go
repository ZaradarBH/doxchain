package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/abs/types"
)

func (k Keeper) GetBreakFactor(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(types.BreakFactorKey))

	if b == nil {
		return sdk.ZeroDec()
	}

	dp := sdk.DecProto{}

	k.cdc.MustUnmarshal(b, &dp)

	return dp.Dec
}

func (k Keeper) SetBreakFactor(ctx sdk.Context, breakFactor sdk.Dec) error {
	if breakFactor.IsNegative() || breakFactor.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(types.ErrBreakFactorOutOfBounds, breakFactor.String())
	}

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&sdk.DecProto{Dec: breakFactor})

	store.Set([]byte(types.BreakFactorKey), b)

	return nil
}
