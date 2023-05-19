package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateClientRegistrationRegistry = "create_client_registrations"
	TypeMsgUpdateClientRegistrationRegistry = "update_client_registrations"
	TypeMsgDeleteClientRegistrationRegistry = "delete_client_registrations"
)

var _ sdk.Msg = &MsgCreateClientRegistrationRegistryRequest{}

func NewMsgCreateClientRegistrationRegistry(clientRegistry ClientRegistrationRegistry) *MsgCreateClientRegistrationRegistryRequest {
	return &MsgCreateClientRegistrationRegistryRequest{
		ClientRegistrationRegistry: clientRegistry,
	}
}

func (msg *MsgCreateClientRegistrationRegistryRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistrationRegistryRequest) Type() string {
	return TypeMsgCreateClientRegistrationRegistry
}

func (msg *MsgCreateClientRegistrationRegistryRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRegistry.Owner.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClientRegistrationRegistryRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClientRegistrationRegistryRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRegistry.Owner.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClientRegistrationRegistryRequest{}

func NewMsgUpdateClientRegistrationRegistry(clientRegistry ClientRegistrationRegistry) *MsgUpdateClientRegistrationRegistryRequest {
	return &MsgUpdateClientRegistrationRegistryRequest{
		ClientRegistrationRegistry: clientRegistry,
	}
}

func (msg *MsgUpdateClientRegistrationRegistryRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClientRegistrationRegistryRequest) Type() string {
	return TypeMsgUpdateClientRegistrationRegistry
}

func (msg *MsgUpdateClientRegistrationRegistryRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRegistry.Owner.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClientRegistrationRegistryRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClientRegistrationRegistryRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRegistry.Owner.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClientRegistrationRegistryRequest{}

func NewMsgDeleteClientRegistrationRegistry(
	creator string,
) *MsgDeleteClientRegistrationRegistryRequest {
	return &MsgDeleteClientRegistrationRegistryRequest{
		Creator: creator,
	}
}
func (msg *MsgDeleteClientRegistrationRegistryRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClientRegistrationRegistryRequest) Type() string {
	return TypeMsgDeleteClientRegistrationRegistry
}

func (msg *MsgDeleteClientRegistrationRegistryRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClientRegistrationRegistryRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClientRegistrationRegistryRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
