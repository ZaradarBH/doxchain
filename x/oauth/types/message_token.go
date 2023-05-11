package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToken = "token"

var _ sdk.Msg = &MsgToken{}

func NewMsgToken(creator string, tenant string, clientId string, clientSecret string, scope string, grantType string) *MsgToken {
	return &MsgToken{
		Creator:      creator,
		Tenant:       tenant,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scope:        scope,
		GrantType:    grantType,
	}
}

func (msg *MsgToken) Route() string {
	return RouterKey
}

func (msg *MsgToken) Type() string {
	return TypeMsgToken
}

func (msg *MsgToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
