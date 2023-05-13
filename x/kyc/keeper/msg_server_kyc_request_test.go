package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/x/kyc/keeper"
	"github.com/be-heroes/doxchain/x/kyc/types"
)

func TestKYCRequestMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.KycKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateKYCRequest{Creator: creator}
	_, err := srv.CreateKYCRequest(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetKYCRequest(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestKYCRequestMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateKYCRequest
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateKYCRequest{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateKYCRequest{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.KycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateKYCRequest{Creator: creator}
			_, err := srv.CreateKYCRequest(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateKYCRequest(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetKYCRequest(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestKYCRequestMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteKYCRequest
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteKYCRequest{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteKYCRequest{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.KycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKYCRequest(wctx, &types.MsgCreateKYCRequest{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteKYCRequest(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKYCRequest(ctx)
				require.False(t, found)
			}
		})
	}
}
