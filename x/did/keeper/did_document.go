package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) GetDidDocumentCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefix(types.DidDocumentCountKeyPrefix))

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDidDocumentCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.KeyPrefix(types.DidDocumentCountKeyPrefix), bz)
}

func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument, override bool) {
	if k.CanOverrideDidDocument(ctx, didDocument, override) {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))

		store.Set(utils.GetKeyBytes(didDocument.Id.GetW3CIdentifier()), k.cdc.MustMarshal(&didDocument))

		k.SetDidDocumentCount(ctx, k.GetDidDocumentCount(ctx)+1)
	}
}

func (k Keeper) GetDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) (result types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	b := store.Get(utils.GetKeyBytes(didDocumentW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))

	store.Delete(utils.GetKeyBytes(didDocumentW3CIdentifier))
}

func (k Keeper) GetAllDidDocument(ctx sdk.Context) (result []types.DidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var didDocument types.DidDocument

		k.cdc.MustUnmarshal(iterator.Value(), &didDocument)
		
		result = append(result, didDocument)
	}

	return
}

func (k Keeper) CanOverrideDidDocument(ctx sdk.Context, didDocument types.DidDocument, override bool) bool {
	if override {
		return true
	}

	match, found := k.GetDidDocument(ctx, didDocument.Id.GetW3CIdentifier())
	
	if found && (!override || didDocument.Id.Creator != match.Id.Creator) {
		return false
	}

	return true
}
