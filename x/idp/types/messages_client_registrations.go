package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateClientRegistry = "create_client_registrations"
	TypeMsgUpdateClientRegistry = "update_client_registrations"
	TypeMsgDeleteClientRegistry = "delete_client_registrations"
)

var _ sdk.Msg = &MsgCreateClientRegistry{}

func NewMsgCreateClientRegistry(clientRegistry ClientRegistry) *MsgCreateClientRegistry {
	return &MsgCreateClientRegistry{
		ClientRegistry: &clientRegistry,
	}
}

func (msg *MsgCreateClientRegistry) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistry) Type() string {
	return TypeMsgCreateClientRegistry
}

func (msg *MsgCreateClientRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistry.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClientRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClientRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistry.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClientRegistry{}

func NewMsgUpdateClientRegistry(clientRegistry ClientRegistry) *MsgUpdateClientRegistry {
	return &MsgUpdateClientRegistry{
		ClientRegistry: &clientRegistry,
	}
}

func (msg *MsgUpdateClientRegistry) Route() string {
	return RouterKey
}

func (msg *MsgUpdateClientRegistry) Type() string {
	return TypeMsgUpdateClientRegistry
}

func (msg *MsgUpdateClientRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistry.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateClientRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateClientRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistry.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClientRegistry{}

func NewMsgDeleteClientRegistry(
	creator string,
) *MsgDeleteClientRegistry {
	return &MsgDeleteClientRegistry{
		Creator: creator,
	}
}
func (msg *MsgDeleteClientRegistry) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClientRegistry) Type() string {
	return TypeMsgDeleteClientRegistry
}

func (msg *MsgDeleteClientRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClientRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClientRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
