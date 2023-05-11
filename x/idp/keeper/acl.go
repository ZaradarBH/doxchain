package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/be-heroes/doxchain/x/idp/types"
)

// GetAccessClientList for a given tenant
func (k Keeper) GetAccessClientList(ctx sdk.Context, tenant string) (acl types.AccessClientList, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessClientListKey))
	tenantAclBytes := store.Get([]byte(tenant))

	if tenantAclBytes == nil {
		return types.AccessClientList{}, sdkerrors.Wrap(types.AccessClientListError, "No ACL exists for given tenant")
	}

	k.cdc.MustUnmarshal(tenantAclBytes, &acl)

	return acl, nil
}
