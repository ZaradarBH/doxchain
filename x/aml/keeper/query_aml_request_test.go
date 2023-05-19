package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/be-heroes/doxchain/testutil/keeper"
	"github.com/be-heroes/doxchain/testutil/nullify"
	"github.com/be-heroes/doxchain/x/aml/types"
)

func TestAMLRegistrationQuery(t *testing.T) {
	keeper, ctx := keepertest.AmlKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestAMLRegistration(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAMLRegistrationRequest
		response *types.QueryGetAMLRegistrationResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAMLRegistrationRequest{},
			response: &types.QueryGetAMLRegistrationResponse{AMLRegistration: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.AMLRegistration(wctx, tc.request)
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
