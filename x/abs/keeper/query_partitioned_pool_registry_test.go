package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/abs/types"
)

var _ = strconv.IntSize

func TestPartitionedPoolsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AbsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPartitionedPools(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPartitionedPoolsRequest
		response *types.QueryGetPartitionedPoolsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPartitionedPoolsRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPartitionedPoolsResponse{PartitionedPools: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPartitionedPoolsRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPartitionedPoolsResponse{PartitionedPools: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPartitionedPoolsRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PartitionedPools(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPartitionedPoolsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AbsKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPartitionedPools(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPartitionedPoolsRequest {
		return &types.QueryAllPartitionedPoolsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PartitionedPoolsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PartitionedPools), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PartitionedPools),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PartitionedPoolsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PartitionedPools), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PartitionedPools),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PartitionedPoolsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PartitionedPools),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PartitionedPoolsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
