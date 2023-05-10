package keeper

import (
	"fmt"
	"encoding/binary"

	"doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDidCount get the total number of did
func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDidCount set the total number of did
func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDid appends a did in the store with a new id and update the count
func (k Keeper) AppendDid(
	ctx sdk.Context,
	did types.Did,
) string {
	// Get the did count
	count := k.GetDidCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	appendedValue := k.cdc.MustMarshal(&did)
	fullyQualifiedDidIdentifier := GetFullyQualifiedDidIdentifier(did)

	store.Set(GetDidIDBytes(fullyQualifiedDidIdentifier), appendedValue)

	// Update did count
	k.SetDidCount(ctx, count+1)

	return fullyQualifiedDidIdentifier
}

// SetDid set a specific did in the store
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&did)
	store.Set(GetDidIDBytes(GetFullyQualifiedDidIdentifier(did)), b)
}

// GetDid returns a Did object from its Did identifier
func (k Keeper) GetDid(ctx sdk.Context, fullyQualifiedDidIdentifier string) (val types.Did, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := store.Get(GetDidIDBytes(fullyQualifiedDidIdentifier))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveDid(ctx sdk.Context, fullyQualifiedDidIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	store.Delete(GetDidIDBytes(fullyQualifiedDidIdentifier))
}

// GetAllDid returns all did
func (k Keeper) GetAllDid(ctx sdk.Context) (list []types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Did
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDidIDBytes returns the byte representation of the Did
func GetDidIDBytes(did string) []byte {
	return []byte(did)
}

// GetDidIDFromBytes returns ID in uint64 format from a byte array
func GetDidIDFromBytes(bz []byte) string {
	return string(bz)
}

// GetFullyQualifiedDidIdentifier from a Did instance
func GetFullyQualifiedDidIdentifier(did types.Did) string {
	//TODO: Move sprintf logic to utility function or method override on the pb.go type that is generated
	return fmt.Sprintf("did:%s:%s", did.MethodName, did.MethodId)
}