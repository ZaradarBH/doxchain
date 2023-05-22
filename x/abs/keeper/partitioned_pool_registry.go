package keeper

import (
	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetPartitionedPoolRegistry(ctx sdk.Context, partitionedPoolRegistry types.PartitionedPoolRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&partitionedPoolRegistry)

	store.Set(types.PartitionedPoolRegistryKey(
		partitionedPoolRegistry.Owner.GetW3CIdentifier(),
	), b)
}

func (k Keeper) GetPartitionedPoolRegistry(ctx sdk.Context, partitionedPoolRegistryW3CIdentifier string) (result types.PartitionedPoolRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	b := store.Get(types.PartitionedPoolRegistryKey(
		partitionedPoolRegistryW3CIdentifier,
	))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemovePartitionedPoolRegistry(ctx sdk.Context, partitionedPoolRegistryW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))

	store.Delete(types.PartitionedPoolRegistryKey(partitionedPoolRegistryW3CIdentifier))
}

func (k Keeper) GetAllPartitionedPoolRegistries(ctx sdk.Context) (list []types.PartitionedPoolRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PartitionedPoolRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
