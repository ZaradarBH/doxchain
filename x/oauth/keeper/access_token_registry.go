package keeper

import (
	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccessTokenRegistry set a specific AccessTokenRegistry in the store from its index
func (k Keeper) SetAccessTokenRegistry(ctx sdk.Context, AccessTokenRegistry types.AccessTokenRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&AccessTokenRegistry)
	store.Set(types.AccessTokenRegistryKey(
		AccessTokenRegistry.Tenant,
	), b)
}

// GetAccessTokenRegistry returns a AccessTokenRegistry from its index
func (k Keeper) GetAccessTokenRegistry(
	ctx sdk.Context,
	tenant string,
) (val types.AccessTokenRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))

	b := store.Get(types.AccessTokenRegistryKey(
		tenant,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccessTokenRegistry removes a AccessTokenRegistry from the store
func (k Keeper) RemoveAccessTokenRegistry(
	ctx sdk.Context,
	tenant string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	store.Delete(types.AccessTokenRegistryKey(
		tenant,
	))
}

// GetAllAccessTokenRegistry returns all AccessTokenRegistry
func (k Keeper) GetAllAccessTokenRegistry(ctx sdk.Context) (list []types.AccessTokenRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccessTokenRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
