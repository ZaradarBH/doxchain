package keeper

import (
	"github.com/be-heroes/doxchain/x/idp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDeviceCodeRegistry set a specific DeviceCodeRegistry in the store from its index
func (k Keeper) SetDeviceCodeRegistry(ctx sdk.Context, deviceCodeRegistry types.DeviceCodeRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodeRegistryKeyPrefix))
	b := k.cdc.MustMarshal(&deviceCodeRegistry)
	store.Set(types.DeviceCodeRegistryKey(
		deviceCodeRegistry.Tenant,
	), b)
}

// GetDeviceCodeRegistry returns a DeviceCodeRegistry from its index
func (k Keeper) GetDeviceCodeRegistry(
	ctx sdk.Context,
	fullyQualifiedDidIdentifier string,
) (val types.DeviceCodeRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodeRegistryKeyPrefix))

	b := store.Get(types.DeviceCodeRegistryKey(
		fullyQualifiedDidIdentifier,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeviceCodeRegistry removes a DeviceCodeRegistry from the store
func (k Keeper) RemoveDeviceCodeRegistry(
	ctx sdk.Context,
	fullyQualifiedDidIdentifier string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodeRegistryKeyPrefix))
	store.Delete(types.DeviceCodeRegistryKey(
		fullyQualifiedDidIdentifier,
	))
}

// GetAllDeviceCodeRegistry returns all DeviceCodeRegistry
func (k Keeper) GetAllDeviceCodeRegistry(ctx sdk.Context) (list []types.DeviceCodeRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodeRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DeviceCodeRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
