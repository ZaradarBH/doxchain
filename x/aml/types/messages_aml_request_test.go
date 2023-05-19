package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateAMLRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAMLRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAMLRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAMLRegistration{
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

func TestMsgUpdateAMLRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAMLRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAMLRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAMLRegistration{
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

func TestMsgDeleteAMLRegistration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAMLRegistration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAMLRegistration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAMLRegistration{
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
