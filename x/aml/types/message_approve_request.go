package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveAMLRequest = "approve_aml_request"

var _ sdk.Msg = &MsgApproveAMLRequest{}

func NewMsgApproveAMLRequest(creator string) *MsgApproveAMLRequest {
  return &MsgApproveAMLRequest{
		Creator: creator,
	}
}

func (msg *MsgApproveAMLRequest) Route() string {
  return RouterKey
}

func (msg *MsgApproveAMLRequest) Type() string {
  return TypeMsgApproveAMLRequest
}

func (msg *MsgApproveAMLRequest) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgApproveAMLRequest) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveAMLRequest) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

