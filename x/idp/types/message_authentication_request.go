package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAuthenticationRequest = "basic_authentication_request"

var _ sdk.Msg = &MsgAuthenticationRequest{}

func NewMsgAuthenticationRequest(creator string) *MsgAuthenticationRequest {
	return &MsgAuthenticationRequest{
		Creator: creator,
	}
}

func (msg *MsgAuthenticationRequest) Route() string {
	return RouterKey
}

func (msg *MsgAuthenticationRequest) Type() string {
	return TypeMsgAuthenticationRequest
}

func (msg *MsgAuthenticationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAuthenticationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuthenticationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
