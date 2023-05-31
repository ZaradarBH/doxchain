package keeper

import (
	types "github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) SetClientRegistrationRegistry(ctx sdk.Context, clientRegistrationRegistry types.ClientRegistrationRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&clientRegistrationRegistry)

	store.Set(utils.GetKeyBytes(clientRegistrationRegistry.Owner.GetW3CIdentifier()), b)
}

func (k Keeper) GetClientRegistrationRegistry(ctx sdk.Context, clientRegistryW3CIdentifier string) (val types.ClientRegistrationRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	b := store.Get(utils.GetKeyBytes(clientRegistryW3CIdentifier))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveClientRegistrationRegistry(ctx sdk.Context, clientRegistrationRegistryW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))

	store.Delete(utils.GetKeyBytes(clientRegistrationRegistryW3CIdentifier))
}

func (k Keeper) GetAllClientRegistrationRegistry(ctx sdk.Context) (list []types.ClientRegistrationRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClientRegistrationRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ClientRegistrationRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
