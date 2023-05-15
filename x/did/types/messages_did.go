package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDidRequest = "create_did"
	TypeMsgUpdateDidRequest = "update_did"
	TypeMsgDeleteDidRequest = "delete_did"
)

var _ sdk.Msg = &MsgCreateDidRequest{}

func NewMsgCreateDidRequest(did Did) *MsgCreateDidRequest {
	return &MsgCreateDidRequest{
		Did: did,
	}
}

func (msg *MsgCreateDidRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateDidRequest) Type() string {
	return TypeMsgCreateDidRequest
}

func (msg *MsgCreateDidRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDidRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDidRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDidRequest{}

func NewMsgUpdateDidRequest(did Did) *MsgUpdateDidRequest {
	return &MsgUpdateDidRequest{
		Did: did,
	}
}

func (msg *MsgUpdateDidRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDidRequest) Type() string {
	return TypeMsgUpdateDidRequest
}

func (msg *MsgUpdateDidRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDidRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDidRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDidRequest{}

func NewMsgDeleteDidRequest(creator string, fullyQualifiedDidIdentifier string) *MsgDeleteDidRequest {
	return &MsgDeleteDidRequest{
		Creator:                     creator,
		FullyQualifiedDidIdentifier: fullyQualifiedDidIdentifier,
	}
}
func (msg *MsgDeleteDidRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDidRequest) Type() string {
	return TypeMsgDeleteDidRequest
}

func (msg *MsgDeleteDidRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDidRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDidRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
