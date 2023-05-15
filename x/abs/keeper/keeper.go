package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/be-heroes/doxchain/x/abs/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

//TODO: Finish PartitionedPool concept
//TODO: Finish Watchlist concept
//TODO: Consider ideas for break factor guards
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetBreakFactor gets the ABS keepers current breakfactor
func (k Keeper) GetBreakFactor(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(types.BreakFactorKey))
	if b == nil {
		return sdk.ZeroDec()
	}

	dp := sdk.DecProto{}
	k.cdc.MustUnmarshal(b, &dp)
	return dp.Dec
}

// SetBreakFactor sets the ABS keepers current breakfactor
func (k Keeper) SetBreakFactor(ctx sdk.Context, breakFactor sdk.Dec) error {
	if breakFactor.IsNegative() || breakFactor.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(types.ErrBreakFactorOutOfBounds, breakFactor.String())
	}

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&sdk.DecProto{Dec: breakFactor})
	store.Set([]byte(types.BreakFactorKey), b)

	return nil
}

func (k Keeper) SetAddressWatchlist(ctx sdk.Context, addr sdk.AccAddress, watchlistEntry types.WatchlistEntry) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)
	b := k.cdc.MustMarshal(&watchlistEntry)
	store.Set(addr.Bytes(), b)
}

func (k Keeper) DeleteAddressWatchlist(ctx sdk.Context, addr sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WatchlistKey)
	store.Delete(addr.Bytes())
}

// get address watch list will return a new watchlist entry if the address is not found
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
