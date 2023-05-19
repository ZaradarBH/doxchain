package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateClientRegistration = "create_client_registration"

var _ sdk.Msg = &MsgCreateClientRegistrationRequest{}

func NewMsgCreateClientRegistration(clientRegistration ClientRegistration) *MsgCreateClientRegistrationRequest {
	return &MsgCreateClientRegistrationRequest{
		ClientRegistration: clientRegistration,
	}
}

func (msg *MsgCreateClientRegistrationRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistrationRequest) Type() string {
	return TypeMsgCreateClientRegistration
}

func (msg *MsgCreateClientRegistrationRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClientRegistrationRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClientRegistrationRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistration.Id.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
