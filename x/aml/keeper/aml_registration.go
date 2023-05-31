package keeper

import (
	"encoding/binary"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	utils "github.com/be-heroes/doxchain/utils"
)

func (k Keeper) GetAMLRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKeyPrefix)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAMLRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKeyPrefix)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendAMLRegistration(ctx sdk.Context, request types.AMLRegistration) string {
	count := k.GetAMLRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&request)
	w3cIdentifier := request.Owner.GetW3CIdentifier()

	store.Set(utils.GetKeyBytes(w3cIdentifier), appendedValue)

	k.SetAMLRegistrationCount(ctx, count+1)

	return w3cIdentifier
}

func (k Keeper) SetAMLRegistration(ctx sdk.Context, request types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKeyPrefix))
	b := k.cdc.MustMarshal(&request)

	store.Set(utils.GetKeyBytes(request.Owner.GetW3CIdentifier()), b)
}

func (k Keeper) GetAMLRegistration(ctx sdk.Context, registrationW3CIdentifier string) (result types.AMLRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKeyPrefix))
	b := store.Get(utils.GetKeyBytes(registrationW3CIdentifier))

	if b == nil {
		return result, false
	}

	k.cdc.MustUnmarshal(b, &result)

	return result, true
}

func (k Keeper) RemoveAMLRegistration(ctx sdk.Context, registrationW3CIdentifier string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKeyPrefix))

	store.Delete(utils.GetKeyBytes(registrationW3CIdentifier))
}

func (k Keeper) GetAllAMLRegistration(ctx sdk.Context) (list []types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AMLRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ApproveAMLRegistration(ctx sdk.Context, registrationW3CIdentifier string) {
	request, found := k.GetAMLRegistration(ctx, registrationW3CIdentifier)

	if found {
		request.Approved = true

		k.SetAMLRegistration(ctx, request)
	}
}
