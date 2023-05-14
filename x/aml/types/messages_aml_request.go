package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
)

const (
	TypeMsgCreateAMLRequest = "create_aml_request"
	TypeMsgDeleteAMLRequest = "delete_aml_request"
)

var _ sdk.Msg = &MsgCreateAMLRequest{}

func NewMsgCreateAMLRequest(creator string, did didTypes.Did) *MsgCreateAMLRequest {
	return &MsgCreateAMLRequest{
		Creator:   creator,
		Did: did,
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
