package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetKYCRegistrationCount get the total number of requests
func (k Keeper) GetKYCRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetKYCRegistrationCount set the total number of requests
func (k Keeper) SetKYCRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendKYCRegistration appends a KYCRegistration to store with a new id and update the count
func (k Keeper) AppendKYCRegistration(
	ctx sdk.Context,
	request types.KYCRegistration,
) string {
	// Get the request count
	count := k.GetKYCRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	appendedValue := k.cdc.MustMarshal(&request)

	store.Set(GetKYCRegistrationIDBytes(request.Owner.Creator), appendedValue)

	// Update KYCRegistration count
	k.SetKYCRegistrationCount(ctx, count+1)

	return request.Owner.Creator
}

// SetKYCRegistration set a specific KYCRegistration in the store
func (k Keeper) SetKYCRegistration(ctx sdk.Context, request types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := k.cdc.MustMarshal(&request)
	store.Set(GetKYCRegistrationIDBytes(request.Owner.Creator), b)
}

// GetKYCRegistration returns a KYCRegistration object from its creator
func (k Keeper) GetKYCRegistration(ctx sdk.Context, creator string) (val types.KYCRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := store.Get(GetKYCRegistrationIDBytes(creator))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveKYCRegistration(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	store.Delete(GetKYCRegistrationIDBytes(creator))
}

// GetAllKYCRegistration returns all requests
func (k Keeper) GetAllKYCRegistration(ctx sdk.Context) (list []types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KYCRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Approves an KYCRegistration in the store
func (k Keeper) ApproveKYCRegistration(ctx sdk.Context, creator string) {
	request, found := k.GetKYCRegistration(ctx, creator)

	if found && creator == request.Owner.Creator {
		request.Approved = true

		k.SetKYCRegistration(ctx, request)
	}
}

// GetKYCRegistrationIDBytes returns the byte representation of the Did
func GetKYCRegistrationIDBytes(did string) []byte {
	return []byte(did)
}

// GetKYCRegistrationIDFromBytes returns ID in uint64 format from a byte array
func GetKYCRegistrationIDFromBytes(bz []byte) string {
	return string(bz)
}
