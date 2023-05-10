package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"doxchain/x/abs/types"
)

// AddToWatchlist tracks account spendings inside a 24-hour rolling window and returns a ErrWatchlistSpendingWindowOverflow if a given account exceeds the "throttled rolling average"
func (k Keeper) AddToWatchlist(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	if coins.Empty() || k.accountKeeper.GetAccount(ctx, addr) == nil {
		return nil
	}

	err := k.bankKeeper.IsSendEnabledCoins(ctx, coins...)

	if err != nil {
		return err
	}

	blockHeight := uint64(ctx.BlockHeight())
	watchlist := k.GetWatchlist(ctx)
	watchlistEntry := &types.WatchlistEntry{
		Address:     addr.String(),
		BlockHeight: blockHeight,
	}
	watchlistEntryIndex := len(watchlist.Entries)
	newCoinPointers := make([]*sdk.Coin, len(coins))

	for i, wle := range watchlist.Entries {
		if wle.Address == addr.String() {
			watchlistEntry = &wle
			watchlistEntryIndex = i

			break
		}
	}

	for _, coin := range coins {
		newCoinPointers = append(newCoinPointers, &coin)
	}

	if watchlistEntry.Coins == nil {
		watchlistEntry.Coins = newCoinPointers
	} else {
		for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
			for _, newCoinPtr := range newCoinPointers {
				if watchlistEntryCoinPtr.GetDenom() == newCoinPtr.GetDenom() {
					watchlistEntryCoinPtr.Amount = watchlistEntryCoinPtr.Amount.Add(newCoinPtr.Amount)

					break
				}
			}
		}
	}

	for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
		//TODO: Decide on oracle design and implement TRA (throttled rolling average) logic
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

	//TODO: Decide how to implement blockExpireOffset (constant | dynamic | param)
	blockExpireOffset := uint64(100000)

	if watchlistEntry.GetBlockHeight()+blockExpireOffset <= blockHeight {
		watchlist.Entries = append(watchlist.Entries[:watchlistEntryIndex], watchlist.Entries[watchlistEntryIndex+1:]...)
	} else {
		if len(watchlist.Entries) == watchlistEntryIndex {
			watchlist.Entries = append(watchlist.Entries, *watchlistEntry)
		} else {
			watchlist.Entries[watchlistEntryIndex] = *watchlistEntry
		}
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.WatchlistKey), k.cdc.MustMarshal(&watchlist))

	return nil
}
