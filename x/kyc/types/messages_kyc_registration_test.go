package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateKYCRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateKYCRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateKYCRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateKYCRegistration{
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

func TestMsgUpdateKYCRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKYCRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKYCRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKYCRegistration{
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

func TestMsgDeleteKYCRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKYCRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKYCRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKYCRegistration{
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
