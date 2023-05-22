package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) SetDidDocument(ctx sdk.Context, didDocument types.DidDocument, override bool) error {
	err := k.CanOverrideDidDocument(ctx, didDocument, override)

	if err != nil {
		return err
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))

	store.Set(GetDidDocumentIDBytes(didDocument.Id.GetW3CIdentifier()), k.cdc.MustMarshal(&didDocument))

	k.SetDidDocumentCount(ctx, k.GetDidDocumentCount(ctx)+1)

	return nil
}

func (k Keeper) GetDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) (val types.DidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))
	b := store.Get(GetDidDocumentIDBytes(didDocumentW3CIdentifier))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveDidDocument(ctx sdk.Context, didDocumentW3CIdentifier string) error {
	match, exists := k.GetDidDocument(ctx, didDocumentW3CIdentifier)

	if exists {
		err := k.CanOverrideDidDocument(ctx, match, true)

		if err != nil {
			return err
		}

		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidDocumentKey))

		store.Delete(GetDidDocumentIDBytes(fullyQualifdidDocumentW3CIdentifieriedW3CIdentifier))
	}

	return nil
}

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

func (k Keeper) CanOverrideDidDocument(ctx sdk.Context, document types.DidDocument, override bool) error {
	didDocumentW3CIdentifier := document.Id.GetW3CIdentifier()
	match, found := k.GetDidDocument(ctx, didDocumentW3CIdentifier)

	if found {
		if !override {
			return sdkerrors.Wrap(types.DidIdentifierCollisionError, fmt.Sprintf("DidDocument with identifier: %s already exists in KVStore", didDocumentW3CIdentifier))
		}

		if document.Id.Creator != match.Id.Creator {
			return sdkerrors.Wrap(types.DidOwnershipError, fmt.Sprintf("DidDocument owned by creator: %s cannot be overriden by creator: %s", match.Id.Creator, document.Id.Creator))
		}
	}

	return nil
}

func GetDidDocumentIDBytes(DidDocument string) []byte {
	return []byte(DidDocument)
}

func GetDidDocumentIDFromBytes(bz []byte) string {
	return string(bz)
}
