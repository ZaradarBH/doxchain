package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDidDocumentCount fetches the DidDocument counter from the KVStore
func (k Keeper) GetDidDocumentCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefix(types.DidDocumentCountKey))

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetDidDocumentCount updates the DidDocument counter in the KVStore
func (k Keeper) SetDidDocumentCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.KeyPrefix(types.DidDocumentCountKey), bz)
}

// SetDidDocument adds a DidDocument to the KVStore and updates the DidDocument counter
func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	b := k.cdc.MustMarshal(&didDocument)

	store.Set(GetDidDocumentIDBytes(didDocument.Id.GetFullyQualifiedDidIdentifier()), b)

	k.SetDidDocumentCount(ctx, k.GetDidDocumentCount(ctx)+1)
}

// GetDidDocument returns a DidDocument from its FullyQualifiedDidDocumentIdentifier (DidDocument:MethodName:MethoDidDocument)
func (k Keeper) GetDidDocument(ctx sdk.Context, fullyQualifiedDidIdentifier string) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	b := store.Get(GetDidDocumentIDBytes(fullyQualifiedDidIdentifier))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveDidDocument removes a DidDocument from the KVStore
func (k Keeper) RemoveDidDocument(ctx sdk.Context, fullyQualifiedDidIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))

	store.Delete(GetDidDocumentIDBytes(fullyQualifiedDidIdentifier))
}

// GetAllDidDocument returns all DidDocuments in the KVStore
func (k Keeper) GetAllDidDocument(ctx sdk.Context) (list []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDidDocumentIDBytes returns the byte representation of the DidDocument
func GetDidDocumentIDBytes(DidDocument string) []byte {
	return []byte(DidDocument)
}

// GetDidDocumentIDFromBytes returns ID in uint64 format from a byte array
func GetDidDocumentIDFromBytes(bz []byte) string {
	return string(bz)
}
