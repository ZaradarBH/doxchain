package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBasicAuthenticationRequest = "basic_authentication_request"

var _ sdk.Msg = &MsgBasicAuthenticationRequest{}

func NewMsgBasicAuthenticationRequest(creator string) *MsgBasicAuthenticationRequest {
	return &MsgBasicAuthenticationRequest{
		Creator: creator,
	}
}

func (msg *MsgBasicAuthenticationRequest) Route() string {
	return RouterKey
}

func (msg *MsgBasicAuthenticationRequest) Type() string {
	return TypeMsgBasicAuthenticationRequest
}

func (msg *MsgBasicAuthenticationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBasicAuthenticationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBasicAuthenticationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
