package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateClientRegistration = "update_client_registration"

var _ sdk.Msg = &MsgUpdateClientRegistration{}

func NewMsgUpdateClientRegistration(clientRegistration ClientRegistration) *MsgUpdateClientRegistration {
	return &MsgUpdateClientRegistration{
		ClientRegistration: &clientRegistration,
	}
}

func (msg *MsgUpdateClientRegistration) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClientRegistration) Type() string {
	return TypeMsgUpdateClientRegistration
}

func (msg *MsgUpdateClientRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClientRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClientRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
