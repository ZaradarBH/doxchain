package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrBreakFactorOutOfBounds          = sdkerrors.Register(ModuleName, 1000, "abs breakfactor out-of-bounds error")
	ErrWatchlistSpendingWindowOverflow = sdkerrors.Register(ModuleName, 1001, "abs watchlist spending window overflow")
)
