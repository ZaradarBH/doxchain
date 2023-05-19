package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/idp/keeper"
	"github.com/be-heroes/doxchain/x/idp/types"
)

var _ = strconv.IntSize

func TestClientRegistrationRegistryMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.IdpKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateClientRegistrationRegistry{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateClientRegistrationRegistry(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetClientRegistrationRegistry(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestClientRegistrationRegistryMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateClientRegistrationRegistry
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateClientRegistrationRegistry{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdpKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateClientRegistrationRegistry(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateClientRegistrationRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetClientRegistrationRegistry(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestClientRegistrationRegistryMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteClientRegistrationRegistry
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteClientRegistrationRegistry{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.IdpKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateClientRegistrationRegistry(wctx, &types.MsgCreateClientRegistrationRegistry{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteClientRegistrationRegistry(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetClientRegistrationRegistry(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
