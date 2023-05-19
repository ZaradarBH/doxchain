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

func TestAMLRegistrationMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.AmlKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateAMLRegistration{Creator: creator}
	_, err := srv.CreateAMLRegistration(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetAMLRegistration(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestAMLRegistrationMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateAMLRegistration
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateAMLRegistration{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateAMLRegistration{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AmlKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateAMLRegistration{Creator: creator}
			_, err := srv.CreateAMLRegistration(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateAMLRegistration(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetAMLRegistration(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestAMLRegistrationMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteAMLRegistration
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteAMLRegistration{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteAMLRegistration{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.AmlKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateAMLRegistration(wctx, &types.MsgCreateAMLRegistration{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteAMLRegistration(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetAMLRegistration(ctx)
				require.False(t, found)
			}
		})
	}
}
