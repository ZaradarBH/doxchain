package types

import (
	didTypes "github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAMLRegistration = "create_aml_request"
	TypeMsgDeleteAMLRegistration = "delete_aml_request"
)

var _ sdk.Msg = &MsgCreateAMLRegistrationRequest{}

func NewMsgCreateAMLRegistration(creator string, owner didTypes.Did) *MsgCreateAMLRegistrationRequest {
	return &MsgCreateAMLRegistrationRequest{
		Creator: creator,
		Owner:   owner,
	}
}

func (msg *MsgCreateAMLRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateAMLRegistrationRequest) Type() string {
	return TypeMsgCreateAMLRegistration
}

func (msg *MsgCreateAMLRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAMLRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAMLRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteAMLRegistrationRequest{}

func NewMsgDeleteAMLRegistration(creator string) *MsgDeleteAMLRegistrationRequest {
	return &MsgDeleteAMLRegistrationRequest{
		Creator: creator,
	}
}
func (msg *MsgDeleteAMLRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAMLRegistrationRequest) Type() string {
	return TypeMsgDeleteAMLRegistration
}

func (msg *MsgDeleteAMLRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAMLRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAMLRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
