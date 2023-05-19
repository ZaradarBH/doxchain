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

func TestKYCRegistrationMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.KycKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateKYCRegistration{Creator: creator}
	_, err := srv.CreateKYCRegistration(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetKYCRegistration(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestKYCRegistrationMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteKYCRegistration
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteKYCRegistration{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteKYCRegistration{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.KycKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKYCRegistration(wctx, &types.MsgCreateKYCRegistration{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteKYCRegistration(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKYCRegistration(ctx)
				require.False(t, found)
			}
		})
	}
}
