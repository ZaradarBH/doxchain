package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetDidDocumentCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefix(types.DidDocumentCountKey))

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDidDocumentCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.KeyPrefix(types.DidDocumentCountKey), bz)
}

func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	store.Set(GetDidDocumentIDBytes(didDocument.Id.GetW3CIdentifier()), k.cdc.MustMarshal(&didDocument))
	k.SetDidDocumentCount(ctx, k.GetDidDocumentCount(ctx)+1)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) (result types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	b := store.Get(GetDidDocumentIDBytes(didDocumentW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	store.Delete(GetDidDocumentIDBytes(didDocumentW3CIdentifier))
}

func (k Keeper) GetAllDidDocument(ctx sdk.Context) (result []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var didDocument types.DidDocument

		k.cdc.MustUnmarshal(iterator.Value(), &didDocument)
		
		result = append(result, didDocument)
	}

	return
}

func GetDidDocumentIDBytes(DidDocument string) []byte {
	return []byte(DidDocument)
}

func GetDidDocumentIDFromBytes(bz []byte) string {
	return string(bz)
}
