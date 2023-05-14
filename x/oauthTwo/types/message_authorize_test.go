package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAuthorize_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAuthorize
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAuthorize{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAuthorize{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
