package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) SetDid(ctx sdk.Context, did types.Did, override bool) {
	if k.CanOverrideDid(ctx, did, override) {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
		
		store.Set(GetDidIDBytes(did.GetW3CIdentifier()), k.cdc.MustMarshal(&did))

		k.SetDidCount(ctx, k.GetDidCount(ctx)+1)
	}
}

func (k Keeper) GetDid(ctx sdk.Context, didW3CIdentifier string) (result types.Did, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := store.Get(GetDidIDBytes(didW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveDid(ctx sdk.Context, didW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	
	store.Delete(GetDidIDBytes(didW3CIdentifier))
}

func (k Keeper) GetAllDid(ctx sdk.Context) (result []types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var did types.Did

		k.cdc.MustUnmarshal(iterator.Value(), &did)

		result = append(result, did)
	}

	return
}

func (k Keeper) CanOverrideDid(ctx sdk.Context, did types.Did, override bool) bool {
	if override {
		return true
	}

	match, found := k.GetDid(ctx, did.GetW3CIdentifier())
	
	if found && (!override || did.Creator != match.Creator) {
		return false
	}

	return true
}

func GetDidIDBytes(did string) []byte {
	return []byte(did)
}

func GetDidIDFromBytes(bz []byte) string {
	return string(bz)
}
