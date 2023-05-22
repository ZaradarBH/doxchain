package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefix(types.DidCountKey))

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.KeyPrefix(types.DidCountKey), bz)
}

func (k Keeper) SetDid(ctx sdk.Context, did types.Did, override bool) error {
	err := k.CanOverrideDid(ctx, did, override)

	if err != nil {
		return err
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	store.Set(GetDidIDBytes(did.GetW3CIdentifier()), k.cdc.MustMarshal(&did))

	k.SetDidCount(ctx, k.GetDidCount(ctx)+1)

	return nil
}

func (k Keeper) GetDid(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (val types.Did, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := store.Get(GetDidIDBytes(fullyQualifiedW3CIdentifier))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveDid(ctx sdk.Context, fullyQualifiedW3CIdentifier string) error {
	match, exists := k.GetDid(ctx, fullyQualifiedW3CIdentifier)

	if exists {
		err := k.CanOverrideDid(ctx, match, true)

		if err != nil {
			return err
		}

		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

		store.Delete(GetDidIDBytes(fullyQualifiedW3CIdentifier))
	}

	return nil
}

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

func (k Keeper) CanOverrideDid(ctx sdk.Context, did types.Did, override bool) error {
	didW3CIdentifier := did.GetW3CIdentifier()
	match, found := k.GetDid(ctx, didW3CIdentifier)

	if found {
		if !override {
			return sdkerrors.Wrap(types.DidIdentifierCollisionError, fmt.Sprintf("Did with identifier: %s already exists in KVStore", didW3CIdentifier))
		}

		if did.Creator != match.Creator {
			return sdkerrors.Wrap(types.DidOwnershipError, fmt.Sprintf("Did owned by creator: %s cannot be overriden by creator: %s", match.Creator, did.Creator))
		}
	}

	return nil
}

func (k Keeper) IsValidDidUrl(ctx sdk.Context, url string) bool {

	return false
}

func GetDidIDBytes(did string) []byte {
	return []byte(did)
}

func GetDidIDFromBytes(bz []byte) string {
	return string(bz)
}
