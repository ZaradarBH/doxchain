package types

import (
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKYCRequest = "create_kyc_request"
	TypeMsgDeleteKYCRequest = "delete_kyc_request"
)

var _ sdk.Msg = &MsgCreateKYCRequest{}

func NewMsgCreateKYCRequest(creator string, did didTypes.Did) *MsgCreateKYCRequest {
	return &MsgCreateKYCRequest{
		Creator: creator,
		Did:     did,
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
