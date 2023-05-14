package keeper_test

import (
    "strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

    keepertest "github.com/be-heroes/doxchain/testutil/keeper"
    "github.com/be-heroes/doxchain/x/abs/keeper"
    "github.com/be-heroes/doxchain/x/abs/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPartitionedPoolsMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AbsKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePartitionedPools{Creator: creator,
		    Index: strconv.Itoa(i),
            
		}
		_, err := srv.CreatePartitionedPools(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPartitionedPools(ctx,
		    expected.Index,
            
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPartitionedPoolsMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdatePartitionedPools
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePartitionedPools{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgUpdatePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AbsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreatePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(0),
                
			}
			_, err := srv.CreatePartitionedPools(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePartitionedPools(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPartitionedPools(ctx,
				    expected.Index,
                    
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPartitionedPoolsMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeletePartitionedPools
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(0),
                
			},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePartitionedPools{Creator: "B",
			    Index: strconv.Itoa(0),
                
			},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(100000),
                
			},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AbsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreatePartitionedPools(wctx, &types.MsgCreatePartitionedPools{Creator: creator,
			    Index: strconv.Itoa(0),
                
			})
			require.NoError(t, err)
			_, err = srv.DeletePartitionedPools(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPartitionedPools(ctx,
				    tc.request.Index,
                    
				)
				require.False(t, found)
			}
		})
	}
}
