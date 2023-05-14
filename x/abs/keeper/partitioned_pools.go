package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/be-heroes/doxchain/x/abs/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetPartitionedPools set a specific partitionedPools in the store from its index
func (k Keeper) SetPartitionedPools(ctx sdk.Context, partitionedPools types.PartitionedPools) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolsKeyPrefix))
	b := k.cdc.MustMarshal(&partitionedPools)
	store.Set(types.PartitionedPoolsKey(
        partitionedPools.Index,
    ), b)
}

// GetPartitionedPools returns a partitionedPools from its index
func (k Keeper) GetPartitionedPools(
    ctx sdk.Context,
    index string,
    
) (val types.PartitionedPools, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolsKeyPrefix))

	b := store.Get(types.PartitionedPoolsKey(
        index,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePartitionedPools removes a partitionedPools from the store
func (k Keeper) RemovePartitionedPools(
    ctx sdk.Context,
    index string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolsKeyPrefix))
	store.Delete(types.PartitionedPoolsKey(
	    index,
    ))
}

// GetAllPartitionedPools returns all partitionedPools
func (k Keeper) GetAllPartitionedPools(ctx sdk.Context) (list []types.PartitionedPools) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PartitionedPoolsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PartitionedPools
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
