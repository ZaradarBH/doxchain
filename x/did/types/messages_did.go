package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDid = "create_did"
	TypeMsgUpdateDid = "update_did"
	TypeMsgDeleteDid = "delete_did"
)

var _ sdk.Msg = &MsgCreateDid{}

func NewMsgCreateDid(did Did) *MsgCreateDid {
	return &MsgCreateDid{
		Did: &did,
	}
}

func (msg *MsgCreateDid) Route() string {
	return RouterKey
}

func (msg *MsgCreateDid) Type() string {
	return TypeMsgCreateDid
}

func (msg *MsgCreateDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDid{}

func NewMsgUpdateDid(did Did) *MsgUpdateDid {
	return &MsgUpdateDid{
		Did: &did,
	}
}

func (msg *MsgUpdateDid) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDid) Type() string {
	return TypeMsgUpdateDid
}

func (msg *MsgUpdateDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Did.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDid{}

func NewMsgDeleteDid(creator string, fullyQualifiedDidIdentifier string) *MsgDeleteDid {
	return &MsgDeleteDid{
		Creator:                     creator,
		FullyQualifiedDidIdentifier: fullyQualifiedDidIdentifier,
	}
}
func (msg *MsgDeleteDid) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDid) Type() string {
	return TypeMsgDeleteDid
}

func (msg *MsgDeleteDid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
