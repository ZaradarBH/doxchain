package keeper

import (
	"github.com/be-heroes/doxchain/x/oauth/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccessTokens set a specific accessTokens in the store from its index
func (k Keeper) SetAccessTokens(ctx sdk.Context, accessTokens types.AccessTokens) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokensKeyPrefix))
	b := k.cdc.MustMarshal(&accessTokens)
	store.Set(types.AccessTokensKey(
		accessTokens.Tenant,
	), b)
}

// GetAccessTokens returns a accessTokens from its index
func (k Keeper) GetAccessTokens(
	ctx sdk.Context,
	tenant string,
) (val types.AccessTokens, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokensKeyPrefix))

	b := store.Get(types.AccessTokensKey(
		tenant,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccessTokens removes a accessTokens from the store
func (k Keeper) RemoveAccessTokens(
	ctx sdk.Context,
	tenant string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokensKeyPrefix))
	store.Delete(types.AccessTokensKey(
		tenant,
	))
}

// GetAllAccessTokens returns all accessTokens
func (k Keeper) GetAllAccessTokens(ctx sdk.Context) (list []types.AccessTokens) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccessTokensKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccessTokens
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
