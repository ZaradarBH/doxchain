package keeper

import (
	"github.com/be-heroes/doxchain/x/oauthtwo/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) SetAuthorizationCodeRegistry(ctx sdk.Context, authorizationCodeRegistry types.AuthorizationCodeRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&authorizationCodeRegistry)

	store.Set(utils.GetKeyBytes(authorizationCodeRegistry.Owner.GetW3CIdentifier()), b)
}

func (k Keeper) GetAuthorizationCodeRegistry(ctx sdk.Context, authorizationCodeRegistryW3CIdentifier string) (result types.AuthorizationCodeRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))
	b := store.Get(utils.GetKeyBytes(authorizationCodeRegistryW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveAuthorizationCodeRegistry(ctx sdk.Context, authorizationCodeRegistryW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthorizationCodeRegistryKeyPrefix))

	store.Delete(utils.GetKeyBytes(authorizationCodeRegistryW3CIdentifier))
}

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
