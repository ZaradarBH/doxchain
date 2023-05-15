package keeper

import (
	"encoding/binary"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAMLRequestCount get the total number of requests
func (k Keeper) GetAMLRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRequestCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAMLRequestCount set the total number of requests
func (k Keeper) SetAMLRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRequestCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAMLRequest appends a AMLRequest to store with a new id and update the count
func (k Keeper) AppendAMLRequest(
	ctx sdk.Context,
	request types.AMLRequest,
) string {
	// Get the request count
	count := k.GetAMLRequestCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	appendedValue := k.cdc.MustMarshal(&request)

	store.Set(GetAMLRequestIDBytes(request.Did.Creator), appendedValue)

	// Update AMLRequest count
	k.SetAMLRequestCount(ctx, count+1)

	return request.Did.Creator
}

// SetAMLRequest set a specific AMLRequest in the store
func (k Keeper) SetAMLRequest(ctx sdk.Context, request types.AMLRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	b := k.cdc.MustMarshal(&request)
	store.Set(GetAMLRequestIDBytes(request.Did.Creator), b)
}

// GetAMLRequest returns a AMLRequest object from its creator
func (k Keeper) GetAMLRequest(ctx sdk.Context, creator string) (val types.AMLRequest, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	b := store.Get(GetAMLRequestIDBytes(creator))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveAMLRequest(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	store.Delete(GetAMLRequestIDBytes(creator))
}

// GetAllAMLRequest returns all requests
func (k Keeper) GetAllAMLRequest(ctx sdk.Context) (list []types.AMLRequest) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AMLRequest
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAMLRequestIDBytes returns the byte representation of the Did
func GetAMLRequestIDBytes(did string) []byte {
	return []byte(did)
}

// GetAMLRequestIDFromBytes returns ID in uint64 format from a byte array
func GetAMLRequestIDFromBytes(bz []byte) string {
	return string(bz)
}
