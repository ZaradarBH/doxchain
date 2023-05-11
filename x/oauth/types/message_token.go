package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTokenRequest = "token"

var _ sdk.Msg = &MsgTokenRequest{}

func NewMsgTokenRequest(creator string, tenant string, clientId string, clientSecret string, scope string, grantType string) *MsgTokenRequest {
	return &MsgTokenRequest{
		Creator:      creator,
		Tenant:       tenant,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Scope:        scope,
		GrantType:    grantType,
	}
}

func (msg *MsgTokenRequest) Route() string {
	return RouterKey
}

func (msg *MsgTokenRequest) Type() string {
	return TypeMsgTokenRequest
}

func (msg *MsgTokenRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTokenRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTokenRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
