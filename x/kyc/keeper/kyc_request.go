package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKYCRequestCount get the total number of requests
func (k Keeper) GetKYCRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKYCRequestCount set the total number of requests
func (k Keeper) SetKYCRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKYCRequest appends a KYCRequest to store with a new id and update the count
func (k Keeper) AppendKYCRequest(
	ctx sdk.Context,
	request types.KYCRequest,
) string {
	// Get the request count
	count := k.GetKYCRequestCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	appendedValue := k.cdc.MustMarshal(&request)
	
	store.Set(GetKYCRequestIDBytes(request.Did.Creator), appendedValue)

	// Update KYCRequest count
	k.SetKYCRequestCount(ctx, count+1)

	return request.Did.Creator
}

// SetKYCRequest set a specific KYCRequest in the store
func (k Keeper) SetKYCRequest(ctx sdk.Context, request types.KYCRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	b := k.cdc.MustMarshal(&request)
	store.Set(GetKYCRequestIDBytes(request.Did.Creator), b)
}

// GetKYCRequest returns a KYCRequest object from its creator
func (k Keeper) GetKYCRequest(ctx sdk.Context, creator string) (val types.KYCRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	b := store.Get(GetKYCRequestIDBytes(creator))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveKYCRequest(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	store.Delete(GetKYCRequestIDBytes(creator))
}

// GetAllKYCRequest returns all requests
func (k Keeper) GetAllKYCRequest(ctx sdk.Context) (list []types.KYCRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KYCRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetKYCRequestIDBytes returns the byte representation of the Did
func GetKYCRequestIDBytes(did string) []byte {
	return []byte(did)
}

// GetKYCRequestIDFromBytes returns ID in uint64 format from a byte array
func GetKYCRequestIDFromBytes(bz []byte) string {
	return string(bz)
}
