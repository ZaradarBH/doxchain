package keeper

import (
	"encoding/binary"
	"github.com/be-heroes/doxchain/x/aml/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetAMLRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAMLRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AMLRegistrationCountKey)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendAMLRegistration(
	ctx sdk.Context,
	request types.AMLRegistration,
) string {
	count := k.GetAMLRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	appendedValue := k.cdc.MustMarshal(&request)

	store.Set(GetAMLRegistrationIDBytes(request.Owner.Creator), appendedValue)

	k.SetAMLRegistrationCount(ctx, count+1)

	return request.Owner.Creator
}

func (k Keeper) SetAMLRegistration(ctx sdk.Context, request types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	b := k.cdc.MustMarshal(&request)

	store.Set(GetAMLRegistrationIDBytes(request.Owner.Creator), b)
}

func (k Keeper) GetAMLRegistration(ctx sdk.Context, creator string) (val types.AMLRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	b := store.Get(GetAMLRegistrationIDBytes(creator))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveAMLRegistration(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))

	store.Delete(GetAMLRegistrationIDBytes(creator))
}

func (k Keeper) GetAllAMLRegistration(ctx sdk.Context) (list []types.AMLRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AMLRegistrationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AMLRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ApproveAMLRegistration(ctx sdk.Context, creator string) {
	request, found := k.GetAMLRegistration(ctx, creator)

	if found && creator == request.Owner.Creator {
		request.Approved = true

		k.SetAMLRegistration(ctx, request)
	}
}

func GetAMLRegistrationIDBytes(did string) []byte {
	return []byte(did)
}

func GetAMLRegistrationIDFromBytes(bz []byte) string {
	return string(bz)
}
