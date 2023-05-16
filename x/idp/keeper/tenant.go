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
	b := k.cdc.MustMarshal(&TenantRegistry)
	store.Set(types.TenantRegistryKey(
		TenantRegistry.Creator,
	), b)
}

// GetTenantRegistry returns a TenantRegistry from its creator
func (k Keeper) GetTenantRegistry(
	ctx sdk.Context,
	creator string,
) (val types.TenantRegistry, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TenantRegistryKeyPrefix))

	b := store.Get(types.TenantRegistryKey(
		creator,
	))

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

	store.Delete(types.TenantRegistryKey(
		creator,
	))
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

// GetTenant for a given tenant identifier
func (k Keeper) GetTenant(ctx sdk.Context, tenantIdentifier string) (tenant types.TenantRegistration, err error) {
	matched := false

	for _, registry := range k.GetAllTenantRegistry(ctx) {
		for _, tenantRegistration := range registry.Tenants {
			if tenantRegistration.Identifier == tenantIdentifier {
				tenant = *tenantRegistration
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

// GetAccessClientList for a given tenant identifier
func (k Keeper) GetAccessClientList(ctx sdk.Context, tenantIdentifier string) (acl types.AccessClientList, err error) {
	tenant, err := k.GetTenant(ctx, tenantIdentifier)

	if err != nil {
		return types.AccessClientList{}, err
	}

	return tenant.AccessClientList, nil
}
