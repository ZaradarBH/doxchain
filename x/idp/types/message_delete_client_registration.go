package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteClientRegistration = "delete_client_registration"

var _ sdk.Msg = &MsgDeleteClientRegistrationRequest{}

func NewMsgDeleteClientRegistration(creator string, clientRegistrationRegistryW3CIdentifier string, clientRegistrationW3CIdentifier string) *MsgDeleteClientRegistrationRequest {
	return &MsgDeleteClientRegistrationRequest{
		Creator: creator,
		ClientRegistrationRegistryW3CIdentifier: clientRegistrationRegistryW3CIdentifier,
		ClientRegistrationW3CIdentifier: clientRegistrationW3CIdentifier,
	}
}

func (msg *MsgDeleteClientRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClientRegistrationRequest) Type() string {
	return TypeMsgDeleteClientRegistration
}

func (msg *MsgDeleteClientRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClientRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClientRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	return nil
}
