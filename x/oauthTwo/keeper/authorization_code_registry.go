package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuthorizationCodeRegistry set a specific authorizationCodeRegistry in the store from its index
func (k Keeper) SetAuthorizationCodeRegistry(ctx sdk.Context, authorizationCodeRegistry types.AuthorizationCodeRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&authorizationCodeRegistry)
	store.Set(types.AuthorizationCodeRegistryKey(
		authorizationCodeRegistry.Tenant,
	), b)
}

// GetAuthorizationCodeRegistry returns a authorizationCodeRegistry from its tenant
func (k Keeper) GetAuthorizationCodeRegistry(
	ctx sdk.Context,
	tenant string,

) (val types.AuthorizationCodeRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))

	b := store.Get(types.AuthorizationCodeRegistryKey(
		tenant,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuthorizationCodeRegistry removes a authorizationCodeRegistry from the store
func (k Keeper) RemoveAuthorizationCodeRegistry(
	ctx sdk.Context,
	tenant string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))
	store.Delete(types.AuthorizationCodeRegistryKey(
		tenant,
	))
}

// GetAllAuthorizationCodeRegistry returns all authorizationCodeRegistry
func (k Keeper) GetAllAuthorizationCodeRegistry(ctx sdk.Context) (list []types.AuthorizationCodeRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AuthorizationCodeRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
