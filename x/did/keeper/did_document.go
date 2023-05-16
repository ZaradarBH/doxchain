package keeper

import (
	"fmt"
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument, override bool) error {
	err := k.CanOverrideDidDocument(ctx, didDocument, override)

	if err != nil {
		return err
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))

	store.Set(GetDidDocumentIDBytes(didDocument.Id.GetFullyQualifiedDidIdentifier()), k.cdc.MustMarshal(&didDocument))

	k.SetDidDocumentCount(ctx, k.GetDidDocumentCount(ctx)+1)
	
	return nil
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
func (k Keeper) RemoveDidDocument(ctx sdk.Context, fullyQualifiedDidIdentifier string) error {
	match, exists := k.GetDidDocument(ctx, fullyQualifiedDidIdentifier)

	if exists {
		err := k.CanOverrideDidDocument(ctx, match, true)

		if err != nil {
			return err
		}

		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))

		store.Delete(GetDidDocumentIDBytes(fullyQualifiedDidIdentifier))
	}

	return nil
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

// CanOverrideDidDocument check if a DidDocument can be safely overwritten without causing and "unapproved identifier collision or ownership error" 
func (k Keeper) CanOverrideDidDocument(ctx sdk.Context, document types.DidDocument, override bool) error {
	fullyQualifiedDidIdentifier := document.Id.GetFullyQualifiedDidIdentifier();
	match, found := k.GetDidDocument(ctx, fullyQualifiedDidIdentifier)

	if found {
		if !override {
			return sdkerrors.Wrap(types.DidIdentifierCollisionError, fmt.Sprintf("DidDocument with identifier: %s already exists in KVStore", fullyQualifiedDidIdentifier))
		}
		
		if document.Id.Creator != match.Id.Creator {
			return sdkerrors.Wrap(types.DidOwnershipError, fmt.Sprintf("DidDocument owned by creator: %s cannot be overriden by creator: %s", match.Id.Creator, document.Id.Creator))
		}
	}

	return nil
}

// GetDidDocumentIDBytes returns the byte representation of the DidDocument
func GetDidDocumentIDBytes(DidDocument string) []byte {
	return []byte(DidDocument)
}

// GetDidDocumentIDFromBytes returns ID in uint64 format from a byte array
func GetDidDocumentIDFromBytes(bz []byte) string {
	return string(bz)
}
