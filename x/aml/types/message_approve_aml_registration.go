package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveAMLRegistration = "approve_aml_request"

var _ sdk.Msg = &MsgApproveAMLRegistrationRequest{}

func NewMsgApproveAMLRegistration(creator string) *MsgApproveAMLRegistrationRequest {
  return &MsgApproveAMLRegistrationRequest{
		Creator: creator,
	}
}

func (msg *MsgApproveAMLRegistrationRequest) Route() string {
  return RouterKey
}

func (msg *MsgApproveAMLRegistrationRequest) Type() string {
  return TypeMsgApproveAMLRegistration
}

func (msg *MsgApproveAMLRegistrationRequest) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  
  if err != nil {
    panic(err)
  }

  return []sdk.AccAddress{creator}
}

func (msg *MsgApproveAMLRegistrationRequest) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)

  return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveAMLRegistrationRequest) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	
  if err != nil {
  	return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)	
  }

  return nil
}

