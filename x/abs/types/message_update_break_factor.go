package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBreakFactor = "update_break_factor"

var _ sdk.Msg = &MsgUpdateBreakFactorRequest{}

func NewMsgUpdateBreakFactorRequest(creator string, breakFactor sdk.Dec) *MsgUpdateBreakFactorRequest {
	return &MsgUpdateBreakFactorRequest{
		Creator: creator,
		BreakFactor: breakFactor,
	}
}

func (msg *MsgUpdateBreakFactorRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBreakFactorRequest) Type() string {
	return TypeMsgUpdateBreakFactor
}

func (msg *MsgUpdateBreakFactorRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBreakFactorRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBreakFactorRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
