package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/idp/types"
)

// SetTenantRegistry set a specific TenantRegistry in the store based on its tenant
func (k Keeper) SetTenantRegistry(ctx sdk.Context, TenantRegistry types.TenantRegistry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Set(types.TenantRegistryKey(TenantRegistry.Owner.Creator), k.cdc.MustMarshal(&TenantRegistry))
}

// GetTenantRegistry returns a TenantRegistry from its creator
func (k Keeper) GetTenantRegistry(
	ctx sdk.Context,
	creator string,
) (val types.TenantRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	b := store.Get(types.TenantRegistryKey(creator))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

// RemoveTenantRegistry removes a TenantRegistry from the store
func (k Keeper) RemoveTenantRegistry(
	ctx sdk.Context,
	creator string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	store.Delete(types.TenantRegistryKey(creator))
}

// GetAllTenantRegistry returns all TenantRegistry
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

// GetTenant for a given tenant identifier (fullyQualifiedW3CIdentifier)
func (k Keeper) GetTenant(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (tenant types.TenantRegistryEntry, err error) {
	matched := false

	//TODO: We need to benchmark how well it performs. If its a big deal it might be worth having a "graph" of tenant fullyqualifieddididentifiers (ids) to speed up this lookup
	for _, registry := range k.GetAllTenantRegistry(ctx) {
		for _, tenantRegistryEntry := range registry.Tenants {
			if tenantRegistryEntry.Id.GetW3CIdentifier() == fullyQualifiedW3CIdentifier {
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

// GetAccessClientList for a given tenant identifier (fullyQualifiedW3CIdentifier)
func (k Keeper) GetAccessClientList(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (acl types.AccessClientList, err error) {
	tenant, err := k.GetTenant(ctx, fullyQualifiedW3CIdentifier)

	if err != nil {
		return acl, err
	}

	return tenant.AccessClientList, nil
}

func (k Keeper) GetTenantConfiguration(ctx sdk.Context, fullyQualifiedW3CIdentifier string) (configuration types.TenantConfiguration, err error) {
	tenant, err := k.GetTenant(ctx, fullyQualifiedW3CIdentifier)

	if err != nil {
		return configuration, err
	}

	configuration = tenant.TenantConfiguration

	return configuration, nil
}
