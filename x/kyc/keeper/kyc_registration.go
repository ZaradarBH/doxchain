package keeper

import (
	"encoding/binary"

	"github.com/be-heroes/doxchain/x/kyc/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetKYCRegistrationCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetKYCRegistrationCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.KYCRegistrationCountKey)
	bz := make([]byte, 8)

	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AppendKYCRegistration(
	ctx sdk.Context,
	request types.KYCRegistration,
) string {
	count := k.GetKYCRegistrationCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	appendedValue := k.cdc.MustMarshal(&request)

	store.Set(GetKYCRegistrationIDBytes(request.Owner.Creator), appendedValue)

	k.SetKYCRegistrationCount(ctx, count+1)

	return request.Owner.Creator
}

func (k Keeper) SetKYCRegistration(ctx sdk.Context, request types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := k.cdc.MustMarshal(&request)

	store.Set(GetKYCRegistrationIDBytes(request.Owner.Creator), b)
}

func (k Keeper) GetKYCRegistration(ctx sdk.Context, creator string) (val types.KYCRegistration, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	b := store.Get(GetKYCRegistrationIDBytes(creator))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)

	return val, true
}

func (k Keeper) RemoveKYCRegistration(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))

	store.Delete(GetKYCRegistrationIDBytes(creator))
}

func (k Keeper) GetAllKYCRegistration(ctx sdk.Context) (list []types.KYCRegistration) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KYCRegistrationKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KYCRegistration
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) ApproveKYCRegistration(ctx sdk.Context, creator string) {
	request, found := k.GetKYCRegistration(ctx, creator)

	if found && creator == request.Owner.Creator {
		request.Approved = true

		k.SetKYCRegistration(ctx, request)
	}
}

func GetKYCRegistrationIDBytes(did string) []byte {
	return []byte(did)
}

func GetKYCRegistrationIDFromBytes(bz []byte) string {
	return string(bz)
}
