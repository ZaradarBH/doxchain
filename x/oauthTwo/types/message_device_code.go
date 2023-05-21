package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeviceCode = "device_code"

var _ sdk.Msg = &MsgDeviceCodeRequest{}

func NewMsgDeviceCodeRequest(creator string, clientId string, scope []string) *MsgDeviceCodeRequest {
	return &MsgDeviceCodeRequest{
		Creator:  creator,
		ClientId: clientId,
		Scope:    scope,
	}
}

func (msg *MsgDeviceCodeRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeviceCodeRequest) Type() string {
	return TypeMsgDeviceCode
}

func (msg *MsgDeviceCodeRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDeviceCodeRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeviceCodeRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	return nil
}
