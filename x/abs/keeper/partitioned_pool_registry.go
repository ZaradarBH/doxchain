package keeper

import (
	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPartitionedPoolRegistry set a specific PartitionedPoolRegistry in the store from its index
func (k Keeper) SetPartitionedPoolRegistry(ctx sdk.Context, partitionedPoolRegistry types.PartitionedPoolRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&partitionedPoolRegistry)
	store.Set(types.PartitionedPoolRegistryKey(
		partitionedPoolRegistry.Owner.Creator,
	), b)
}

// GetPartitionedPoolRegistry returns a PartitionedPoolRegistry from its creator
func (k Keeper) GetPartitionedPoolRegistry(
	ctx sdk.Context,
	creator string,
) (val types.PartitionedPoolRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	b := store.Get(types.PartitionedPoolRegistryKey(
		creator,
	))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemovePartitionedPoolRegistry removes a PartitionedPoolRegistry from the store
func (k Keeper) RemovePartitionedPoolRegistry(
	ctx sdk.Context,
	creator string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolRegistryKeyPrefix))
	store.Delete(types.PartitionedPoolRegistryKey(
		creator,
	))
}

// GetAllPartitionedPools returns all partitionedPools
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
