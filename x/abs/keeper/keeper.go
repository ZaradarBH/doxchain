package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"doxchain/x/abs/types"
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

// GetWatchlist fetches the ABS watchlist from the KVStore
func (k Keeper) GetWatchlist(ctx sdk.Context) types.Watchlist {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(types.WatchlistKey))

	if b == nil {
		return types.Watchlist{}
	}

	dp := types.Watchlist{}
	k.cdc.MustUnmarshal(b, &dp)
	return dp
}
