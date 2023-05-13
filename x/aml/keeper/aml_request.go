package keeper

import (
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAMLRequest set aMLRequest in the store
func (k Keeper) SetAMLRequest(ctx sdk.Context, aMLRequest types.AMLRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	b := k.cdc.MustMarshal(&aMLRequest)
	store.Set([]byte{0}, b)
}

// GetAMLRequest returns aMLRequest
func (k Keeper) GetAMLRequest(ctx sdk.Context) (val types.AMLRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAMLRequest removes aMLRequest from the store
func (k Keeper) RemoveAMLRequest(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	store.Delete([]byte{0})
}
