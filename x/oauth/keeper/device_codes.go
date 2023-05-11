package keeper

import (
	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDeviceCodes set a specific deviceCodes in the store from its index
func (k Keeper) SetDeviceCodes(ctx sdk.Context, deviceCodes types.DeviceCodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodesKeyPrefix))
	b := k.cdc.MustMarshal(&deviceCodes)
	store.Set(types.DeviceCodesKey(
		deviceCodes.Tenant,
	), b)
}

// GetDeviceCodes returns a deviceCodes from its index
func (k Keeper) GetDeviceCodes(
	ctx sdk.Context,
	tenant string,
) (val types.DeviceCodes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodesKeyPrefix))

	b := store.Get(types.DeviceCodesKey(
		tenant,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeviceCodes removes a deviceCodes from the store
func (k Keeper) RemoveDeviceCodes(
	ctx sdk.Context,
	tenant string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodesKeyPrefix))
	store.Delete(types.DeviceCodesKey(
		tenant,
	))
}

// GetAllDeviceCodes returns all deviceCodes
func (k Keeper) GetAllDeviceCodes(ctx sdk.Context) (list []types.DeviceCodes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeviceCodesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DeviceCodes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
