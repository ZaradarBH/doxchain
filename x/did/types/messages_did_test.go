package types

import (
	"testing"

	"github.com/be-heroes/doxchain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateDid_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDid{
				&Did{
					Creator:    "invalid_address",
					Url:        "did:method:id",
					MethodName: "method",
					MethodId:   "id",
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDid{
				&Did{
					Creator:    sample.AccAddress(),
					Url:        "did:method:id",
					MethodName: "method",
					MethodId:   "id",
				},
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

func TestMsgUpdateDid_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDid{
				&Did{
					Creator:    "invalid_address",
					Url:        "did:method:id",
					MethodName: "method",
					MethodId:   "id",
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDid{
				&Did{
					Creator:    sample.AccAddress(),
					Url:        "did:method:id",
					MethodName: "method",
					MethodId:   "id",
				},
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

func TestMsgDeleteDid_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDid
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDid{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDid{
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
