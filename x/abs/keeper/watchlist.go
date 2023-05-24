package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/abs/types"
)

func (k Keeper) SetAddressWatchlist(ctx sdk.Context, addr sdk.AccAddress, watchlistEntry types.WatchlistEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)
	b := k.cdc.MustMarshal(&watchlistEntry)

	store.Set(addr.Bytes(), b)
}

func (k Keeper) DeleteAddressWatchlist(ctx sdk.Context, addr sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)

	store.Delete(addr.Bytes())
}

func (k Keeper) GetAddressWatchlist(ctx sdk.Context, addr sdk.AccAddress) types.WatchlistEntry {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)
	b := store.Get(addr.Bytes())

	if b == nil {
		return types.WatchlistEntry{
			Address:     addr.String(),
			BlockHeight: uint64(ctx.BlockHeight()),
			Coins:       sdk.NewCoins(),
		}
	}

	var entry types.WatchlistEntry

	k.cdc.MustUnmarshal(b, &entry)

	return entry
}

func (k Keeper) IterateWatchList(ctx sdk.Context, cb func(entry types.WatchlistEntry) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)
	iter := store.Iterator(nil, nil)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var entry types.WatchlistEntry
		k.cdc.MustUnmarshal(iter.Value(), &entry)

		if cb(entry) {
			break
		}
	}
}

func (k Keeper) AddToWatchlist(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	if coins.Empty() || k.accountKeeper.GetAccount(ctx, addr) == nil {
		return nil
	}

	err := k.bankKeeper.IsSendEnabledCoins(ctx, coins...)

	if err != nil {
		return err
	}

	blockHeight := uint64(ctx.BlockHeight())
	watchlistEntry := k.GetAddressWatchlist(ctx, addr)
	watchlistEntry.Coins = watchlistEntry.Coins.Add(coins...)

	for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
		//TODO: Finish TRA (throttled rolling average) logic once oracle is done.
		throttledRollingAverage := sdk.ZeroInt()

		if throttledRollingAverage.GT(watchlistEntryCoinPtr.Amount) {
			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventWatchlist,
					sdk.NewAttribute(types.AttributeKeyAddress, watchlistEntry.Address),
					sdk.NewAttribute(types.AttributeKeyDenom, watchlistEntryCoinPtr.Denom),
				),
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				),
			})

			return sdkerrors.Wrap(types.ErrWatchlistSpendingWindowOverflow, watchlistEntry.GetAddress())
		}
	}

	var blockExpireOffset sdk.Int

	k.paramstore.Get(ctx, types.ParamStoreKeyBlockExpireOffset, &blockExpireOffset)

	if watchlistEntry.GetBlockHeight()+blockExpireOffset.Uint64() <= blockHeight {
		k.DeleteAddressWatchlist(ctx, addr)
	} else {
		k.SetAddressWatchlist(ctx, addr, watchlistEntry)
	}

	return nil
}
