package keeper

import (
	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKYCRequest set kYCRequest in the store
func (k Keeper) SetKYCRequest(ctx sdk.Context, kYCRequest types.KYCRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	b := k.cdc.MustMarshal(&kYCRequest)
	store.Set([]byte{0}, b)
}

// GetKYCRequest returns kYCRequest
func (k Keeper) GetKYCRequest(ctx sdk.Context) (val types.KYCRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKYCRequest removes kYCRequest from the store
func (k Keeper) RemoveKYCRequest(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	store.Delete([]byte{0})
}
