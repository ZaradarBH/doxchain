package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateClientRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateClientRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateClientRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateClientRegistry{
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

func TestMsgUpdateClientRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateClientRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateClientRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateClientRegistry{
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

func TestMsgDeleteClientRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteClientRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteClientRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteClientRegistry{
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
