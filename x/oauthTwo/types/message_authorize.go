package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAuthorize = "authorize"

var _ sdk.Msg = &MsgAuthorizeRequest{}

func NewMsgAuthorizeRequest(creator string, clientId string, scope string) *MsgAuthorizeRequest {
	return &MsgAuthorizeRequest{
		Creator:  creator,
		ClientId: clientId,
		Scope:    scope,
	}
}

func (msg *MsgAuthorizeRequest) Route() string {
	return RouterKey
}

func (msg *MsgAuthorizeRequest) Type() string {
	return TypeMsgAuthorize
}

func (msg *MsgAuthorizeRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAuthorizeRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuthorizeRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
