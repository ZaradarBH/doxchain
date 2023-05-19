package types

import (
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKYCRegistration = "create_kyc_request"
	TypeMsgDeleteKYCRegistration = "delete_kyc_request"
)

var _ sdk.Msg = &MsgCreateKYCRegistrationRequest{}

func NewMsgCreateKYCRegistration(creator string, owner didTypes.Did) *MsgCreateKYCRegistrationRequest {
	return &MsgCreateKYCRegistrationRequest{
		Creator: creator,
		Owner: owner,
	}
}

func (msg *MsgCreateKYCRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateKYCRegistrationRequest) Type() string {
	return TypeMsgCreateKYCRegistration
}

func (msg *MsgCreateKYCRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKYCRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKYCRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKYCRegistrationRequest{}

func NewMsgDeleteKYCRegistration(creator string) *MsgDeleteKYCRegistrationRequest {
	return &MsgDeleteKYCRegistrationRequest{
		Creator: creator,
	}
}
func (msg *MsgDeleteKYCRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKYCRegistrationRequest) Type() string {
	return TypeMsgDeleteKYCRegistration
}

func (msg *MsgDeleteKYCRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKYCRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKYCRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
