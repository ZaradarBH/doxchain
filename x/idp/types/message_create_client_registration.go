package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateClientRegistration = "create_client_registration"

var _ sdk.Msg = &MsgCreateClientRegistration{}

func NewMsgCreateClientRegistration(clientRegistration ClientRegistration) *MsgCreateClientRegistration {
	return &MsgCreateClientRegistration{
		ClientRegistration: clientRegistration,
	}
}

func (msg *MsgCreateClientRegistration) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistration) Type() string {
	return TypeMsgCreateClientRegistration
}

func (msg *MsgCreateClientRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClientRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClientRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
