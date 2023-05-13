package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAMLRequest = "create_aml_request"
	TypeMsgUpdateAMLRequest = "update_aml_request"
	TypeMsgDeleteAMLRequest = "delete_aml_request"
)

var _ sdk.Msg = &MsgCreateAMLRequest{}

func NewMsgCreateAMLRequest(creator string, firstName string, lastName string, approved bool) *MsgCreateAMLRequest {
	return &MsgCreateAMLRequest{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Approved:  approved,
	}
}

func (msg *MsgCreateAMLRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateAMLRequest) Type() string {
	return TypeMsgCreateAMLRequest
}

func (msg *MsgCreateAMLRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAMLRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAMLRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAMLRequest{}

func NewMsgUpdateAMLRequest(creator string, firstName string, lastName string, approved bool) *MsgUpdateAMLRequest {
	return &MsgUpdateAMLRequest{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Approved:  approved,
	}
}

func (msg *MsgUpdateAMLRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAMLRequest) Type() string {
	return TypeMsgUpdateAMLRequest
}

func (msg *MsgUpdateAMLRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAMLRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAMLRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAMLRequest{}

func NewMsgDeleteAMLRequest(creator string) *MsgDeleteAMLRequest {
	return &MsgDeleteAMLRequest{
		Creator: creator,
	}
}
func (msg *MsgDeleteAMLRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAMLRequest) Type() string {
	return TypeMsgDeleteAMLRequest
}

func (msg *MsgDeleteAMLRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAMLRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAMLRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
