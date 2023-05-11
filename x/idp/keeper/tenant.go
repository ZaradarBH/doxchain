package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/be-heroes/doxchain/x/idp/types"
)

// GetTenant for a given tenant identifier
func (k Keeper) GetTenant(ctx sdk.Context, tenantIdentifier string) (tenant types.Tenant, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	tenantListBytes := store.Get(types.KeyPrefix(types.TenantListKey))

	if tenantListBytes == nil {
		return types.Tenant{}, sdkerrors.Wrap(types.TenantListError, "No tenant list found")
	}

	tenants := &types.TenantList{}

	k.cdc.MustUnmarshal(tenantListBytes, tenants)
	
	for _, tenantEntry := range tenants.Entries {
		if(tenantEntry.Identifier == tenantIdentifier){
			tenant = *tenantEntry

			break
		}
	}

	if &tenant == nil {
		return types.Tenant{}, sdkerrors.Wrap(types.TenantError, "No tenant found for given identifier")
	}

	return tenant, nil
}

// GetAccessClientList for a given tenant identifier
func (k Keeper) GetAccessClientList(ctx sdk.Context, tenantIdentifier string) (acl types.AccessClientList, err error) {
	tenant, err := k.GetTenant(ctx, tenantIdentifier)

	if err != nil {
		return types.AccessClientList{}, err
	}

	if &tenant.AccessClientList == nil {
		return types.AccessClientList{}, sdkerrors.Wrap(types.AccessClientListError, "No ACL found for tenant")
	}

	return tenant.AccessClientList, nil
}
