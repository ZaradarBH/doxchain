package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKYCRequest = "create_kyc_request"
	TypeMsgUpdateKYCRequest = "update_kyc_request"
	TypeMsgDeleteKYCRequest = "delete_kyc_request"
)

var _ sdk.Msg = &MsgCreateKYCRequest{}

func NewMsgCreateKYCRequest(creator string, firstName string, lastName string, approved bool) *MsgCreateKYCRequest {
	return &MsgCreateKYCRequest{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Approved:  approved,
	}
}

func (msg *MsgCreateKYCRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateKYCRequest) Type() string {
	return TypeMsgCreateKYCRequest
}

func (msg *MsgCreateKYCRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKYCRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKYCRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKYCRequest{}

func NewMsgUpdateKYCRequest(creator string, firstName string, lastName string, approved bool) *MsgUpdateKYCRequest {
	return &MsgUpdateKYCRequest{
		Creator:   creator,
		FirstName: firstName,
		LastName:  lastName,
		Approved:  approved,
	}
}

func (msg *MsgUpdateKYCRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKYCRequest) Type() string {
	return TypeMsgUpdateKYCRequest
}

func (msg *MsgUpdateKYCRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKYCRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKYCRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKYCRequest{}

func NewMsgDeleteKYCRequest(creator string) *MsgDeleteKYCRequest {
	return &MsgDeleteKYCRequest{
		Creator: creator,
	}
}
func (msg *MsgDeleteKYCRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKYCRequest) Type() string {
	return TypeMsgDeleteKYCRequest
}

func (msg *MsgDeleteKYCRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKYCRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKYCRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
