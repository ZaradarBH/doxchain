package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
)

func (k Keeper) SetTenantRegistry(ctx sdk.Context, tenantRegistry types.TenantRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Set(types.TenantRegistryKey(tenantRegistry.Owner.Creator), k.cdc.MustMarshal(&tenantRegistry))
}

func (k Keeper) GetTenantRegistry(ctx sdk.Context, owner sdk.AccAddress) (val types.TenantRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))
	b := store.Get(types.TenantRegistryKey(owner.String()))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveTenantRegistry(ctx sdk.Context, owner sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Delete(types.TenantRegistryKey(owner.String()))
}

func (k Keeper) GetAllTenantRegistry(ctx sdk.Context) (list []types.TenantRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TenantRegistry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetTenant(ctx sdk.Context, tenantW3CIdentifier string) (tenant types.TenantRegistryEntry, err error) {
	matched := false

	//TODO: Benchmark how well it performs. If its a big deal it might be worth having a "graph" of tenant dids to speed up this logic
	for _, registry := range k.GetAllTenantRegistry(ctx) {
		for _, tenantRegistryEntry := range registry.Tenants {
			if tenantRegistryEntry.Id.GetW3CIdentifier() == tenantW3CIdentifier {
				tenant = tenantRegistryEntry
				matched = true

				break
			}
		}
	}

	if !matched {
		err = sdkerrors.Wrap(types.TenantError, "No tenant found for given identifier")
	}

	return tenant, err
}

func (k Keeper) GetAccessClientList(ctx sdk.Context, tenantW3CIdentifier string) (acl types.AccessClientList, err error) {
	tenant, err := k.GetTenant(ctx, tenantW3CIdentifier)

	if err != nil {
		return acl, err
	}

	return tenant.AccessClientList, nil
}

func (k Keeper) GetTenantConfiguration(ctx sdk.Context, tenantW3CIdentifier string) (configuration types.TenantConfiguration, err error) {
	tenant, err := k.GetTenant(ctx, tenantW3CIdentifier)

	if err != nil {
		return configuration, err
	}

	configuration = tenant.TenantConfiguration

	return configuration, nil
}
