package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetAccessTokenRegistry(ctx sdk.Context, AccessTokenRegistry types.AccessTokenRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&AccessTokenRegistry)

	store.Set(types.AccessTokenRegistryKey(
		AccessTokenRegistry.Owner.Creator,
	), b)
}

func (k Keeper) GetAccessTokenRegistry(
	ctx sdk.Context,
	fullyQualifiedW3CIdentifier string,
) (val types.AccessTokenRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))

	b := store.Get(types.AccessTokenRegistryKey(
		fullyQualifiedW3CIdentifier,
	))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveAccessTokenRegistry(
	ctx sdk.Context,
	fullyQualifiedW3CIdentifier string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	
	store.Delete(types.AccessTokenRegistryKey(
		fullyQualifiedW3CIdentifier,
	))
}

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
