package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) SetAccessTokenRegistry(ctx sdk.Context, AccessTokenRegistry types.AccessTokenRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&AccessTokenRegistry)

	store.Set(utils.GetKeyBytes(AccessTokenRegistry.Owner.GetW3CIdentifier()), b)
}

func (k Keeper) GetAccessTokenRegistry(ctx sdk.Context, accessTokenRegistryW3CIdentifier string) (result types.AccessTokenRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))
	b := store.Get(utils.GetKeyBytes(accessTokenRegistryW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveAccessTokenRegistry(ctx sdk.Context, accessTokenRegistryW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokenRegistryKeyPrefix))

	store.Delete(utils.GetKeyBytes(accessTokenRegistryW3CIdentifier))
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
