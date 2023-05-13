package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/aml/keeper"
	"github.com/be-heroes/doxchain/x/aml/types"
)

func TestAMLRequestMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AmlKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateAMLRequest{Creator: creator}
	_, err := srv.CreateAMLRequest(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetAMLRequest(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestAMLRequestMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAMLRequest
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAMLRequest{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAMLRequest{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AmlKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateAMLRequest{Creator: creator}
			_, err := srv.CreateAMLRequest(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateAMLRequest(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetAMLRequest(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestAMLRequestMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAMLRequest
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAMLRequest{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAMLRequest{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AmlKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateAMLRequest(wctx, &types.MsgCreateAMLRequest{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAMLRequest(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetAMLRequest(ctx)
				require.False(t, found)
			}
		})
	}
}
