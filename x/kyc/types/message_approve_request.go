package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveKYCRequest = "approve_kyc_request"

var _ sdk.Msg = &MsgApproveKYCRequest{}

func NewMsgApproveKYCRequest(creator string) *MsgApproveKYCRequest {
  return &MsgApproveKYCRequest{
		Creator: creator,
	}
}

func (msg *MsgApproveKYCRequest) Route() string {
  return RouterKey
}

func (msg *MsgApproveKYCRequest) Type() string {
  return TypeMsgApproveKYCRequest
}

func (msg *MsgApproveKYCRequest) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgApproveKYCRequest) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveKYCRequest) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

