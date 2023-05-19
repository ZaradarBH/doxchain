package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateClientRegistration = "update_client_registration"

var _ sdk.Msg = &MsgUpdateClientRegistrationRequest{}

func NewMsgUpdateClientRegistration(clientRegistration ClientRegistration) *MsgUpdateClientRegistrationRequest {
	return &MsgUpdateClientRegistrationRequest{
		ClientRegistration: clientRegistration,
	}
}

func (msg *MsgUpdateClientRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClientRegistrationRequest) Type() string {
	return TypeMsgUpdateClientRegistration
}

func (msg *MsgUpdateClientRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClientRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClientRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	return nil
}
