package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/be-heroes/doxchain/x/abs/types"
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
	watchlistEntry := k.GetAddressWatchlist(ctx, addr)
	watchlistEntry.Coins = watchlistEntry.Coins.Add(coins...)

	for _, watchlistEntryCoinPtr := range watchlistEntry.Coins {
		//TODO: Implement TRA (throttled rolling average) logic
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

	//TODO: Manage blockExpireOffset via consensus params
	blockExpireOffset := uint64(100000)

	if watchlistEntry.GetBlockHeight()+blockExpireOffset <= blockHeight {
		k.DeleteAddressWatchlist(ctx, addr)
	} else {
		k.SetAddressWatchlist(ctx, addr, watchlistEntry)
	}

	return nil
}
