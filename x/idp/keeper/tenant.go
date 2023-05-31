package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/be-heroes/doxchain/x/idp/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) SetTenantRegistry(ctx sdk.Context, tenantRegistry types.TenantRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Set(utils.GetKeyBytes(tenantRegistry.Owner.Creator), k.cdc.MustMarshal(&tenantRegistry))
}

func (k Keeper) GetTenantRegistry(ctx sdk.Context, tenantRegistryW3CIdentifier string) (result types.TenantRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))
	b := store.Get(utils.GetKeyBytes(tenantRegistryW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveTenantRegistry(ctx sdk.Context, tenantRegistryW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Delete(utils.GetKeyBytes(tenantRegistryW3CIdentifier))
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

func (k Keeper) GetTenant(ctx sdk.Context, tenantW3CIdentifier string) (tenant types.TenantRegistryEntry) {
	//TODO: Benchmark how well it performs. If its a big deal it might be worth having a "graph" of tenant dids to speed up this logic
	for _, registry := range k.GetAllTenantRegistry(ctx) {
		for _, tenantRegistryEntry := range registry.Tenants {
			if tenantRegistryEntry.Id.GetW3CIdentifier() == tenantW3CIdentifier {
				tenant = tenantRegistryEntry

				break
			}
		}
	}

	return tenant
}

func (k Keeper) GetAccessClientList(ctx sdk.Context, tenantW3CIdentifier string) (acl types.AccessClientList) {
	tenant := k.GetTenant(ctx, tenantW3CIdentifier)

	return tenant.AccessClientList
}

func (k Keeper) GetTenantConfiguration(ctx sdk.Context, tenantW3CIdentifier string) (configuration types.TenantConfiguration) {
	tenant := k.GetTenant(ctx, tenantW3CIdentifier)

	return tenant.TenantConfiguration
}
