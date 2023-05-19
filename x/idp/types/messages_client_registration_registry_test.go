package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateClientRegistrationRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateClientRegistrationRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateClientRegistrationRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateClientRegistrationRegistry{
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

func TestMsgUpdateClientRegistrationRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateClientRegistrationRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateClientRegistrationRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateClientRegistrationRegistry{
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

func TestMsgDeleteClientRegistrationRegistry_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteClientRegistrationRegistry
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteClientRegistrationRegistry{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteClientRegistrationRegistry{
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
