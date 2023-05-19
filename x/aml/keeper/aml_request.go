package keeper

import (
	"encoding/binary"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAMLRegistrationCount get the total number of requests
func (k Keeper) GetAMLRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAMLRegistrationCount set the total number of requests
func (k Keeper) SetAMLRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAMLRegistration appends a AMLRegistration to store with a new id and update the count
func (k Keeper) AppendAMLRegistration(
	ctx sdk.Context,
	request types.AMLRegistration,
) string {
	// Get the request count
	count := k.GetAMLRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	appendedValue := k.cdc.MustMarshal(&request)

	store.Set(GetAMLRegistrationIDBytes(request.Owner.Creator), appendedValue)

	// Update AMLRegistration count
	k.SetAMLRegistrationCount(ctx, count+1)

	return request.Owner.Creator
}

// SetAMLRegistration set a specific AMLRegistration in the store
func (k Keeper) SetAMLRegistration(ctx sdk.Context, request types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	b := k.cdc.MustMarshal(&request)
	store.Set(GetAMLRegistrationIDBytes(request.Owner.Creator), b)
}

// GetAMLRegistration returns a AMLRegistration object from its creator
func (k Keeper) GetAMLRegistration(ctx sdk.Context, creator string) (val types.AMLRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	b := store.Get(GetAMLRegistrationIDBytes(creator))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveAMLRegistration(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	store.Delete(GetAMLRegistrationIDBytes(creator))
}

// GetAllAMLRegistration returns all requests
func (k Keeper) GetAllAMLRegistration(ctx sdk.Context) (list []types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AMLRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Approves an AMLRegistration in the store
func (k Keeper) ApproveAMLRegistration(ctx sdk.Context, creator string) {
	request, found := k.GetAMLRegistration(ctx, creator)

	if found && creator == request.Owner.Creator {
		request.Approved = true

		k.SetAMLRegistration(ctx, request)
	}
}

// GetAMLRegistrationIDBytes returns the byte representation of the Did
func GetAMLRegistrationIDBytes(did string) []byte {
	return []byte(did)
}

// GetAMLRegistrationIDFromBytes returns ID in uint64 format from a byte array
func GetAMLRegistrationIDFromBytes(bz []byte) string {
	return string(bz)
}
