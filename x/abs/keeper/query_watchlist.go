package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/abs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryWatchlist(goCtx context.Context, req *types.QueryWatchlistRequest) (*types.QueryWatchlistResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	watchlist := types.Watchlist{Entries: []types.WatchlistEntry{}}

	k.IterateWatchList(sdk.UnwrapSDKContext(goCtx), func(entry types.WatchlistEntry) bool {
		watchlist.Entries = append(watchlist.Entries, entry)

		return false
	})

	return &types.QueryWatchlistResponse{
		Watchlist: watchlist,
	}, nil
}
