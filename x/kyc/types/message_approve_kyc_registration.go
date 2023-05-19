package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveKYCRegistration = "approve_kyc_request"

var _ sdk.Msg = &MsgApproveKYCRegistrationRequest{}

func NewMsgApproveKYCRegistration(creator string) *MsgApproveKYCRegistrationRequest {
  return &MsgApproveKYCRegistrationRequest{
		Creator: creator,
	}
}

func (msg *MsgApproveKYCRegistrationRequest) Route() string {
  return RouterKey
}

func (msg *MsgApproveKYCRegistrationRequest) Type() string {
  return TypeMsgApproveKYCRegistration
}

func (msg *MsgApproveKYCRegistrationRequest) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)

  if err != nil {
    panic(err)
  }
  
  return []sdk.AccAddress{creator}
}

func (msg *MsgApproveKYCRegistrationRequest) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  
  return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveKYCRegistrationRequest) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)

  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }

  return nil
}

