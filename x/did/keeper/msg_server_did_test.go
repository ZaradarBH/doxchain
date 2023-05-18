package keeper_test

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/be-heroes/doxchain/x/did/types"
)

func TestDidMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateDid(ctx, &types.MsgCreateDid{
			Did: &types.Did{
				Creator:    creator,
				MethodName: "method",
				MethodId:   string(i),
			}})
		require.NoError(t, err)
		require.Equal(t, fmt.Sprintf("did:method:%d", i), resp.FullyQualifiedW3CIdentifier)
	}
}

func TestDidMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDid
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDid{Did: &types.Did{
				Creator:    creator,
				MethodName: "method",
				MethodId:   "id",
			}},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDid{Did: &types.Did{
				Creator:    "B",
				MethodName: "method",
				MethodId:   "id",
			}},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDid{Did: &types.Did{
				Creator:    creator,
				MethodName: "method",
				MethodId:   "id",
			}},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateDid(ctx, &types.MsgCreateDid{Did: &types.Did{
				Creator:    creator,
				MethodName: "method",
				MethodId:   "id",
			}})
			require.NoError(t, err)

			_, err = srv.UpdateDid(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDidMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDid
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteDid{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteDid{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteDid{Creator: creator, FullyQualifiedW3CIdentifier: "did:method:id"},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateDid(ctx, &types.MsgCreateDid{
				Did: &types.Did{
					Creator:    creator,
					MethodName: "method",
					MethodId:   "id",
				},
			})
			require.NoError(t, err)
			_, err = srv.DeleteDid(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
