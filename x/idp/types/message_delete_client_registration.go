package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteClientRegistration = "delete_client_registration"

var _ sdk.Msg = &MsgDeleteClientRegistration{}

func NewMsgDeleteClientRegistration(creator string, name string) *MsgDeleteClientRegistration {
	return &MsgDeleteClientRegistration{
		Creator: creator,
		Name: name,
	}
}

func (msg *MsgDeleteClientRegistration) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClientRegistration) Type() string {
	return TypeMsgDeleteClientRegistration
}

func (msg *MsgDeleteClientRegistration) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClientRegistration) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClientRegistration) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
