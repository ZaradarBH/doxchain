package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePartitionedPools = "create_partitioned_pools"
	TypeMsgUpdatePartitionedPools = "update_partitioned_pools"
	TypeMsgDeletePartitionedPools = "delete_partitioned_pools"
)

var _ sdk.Msg = &MsgCreatePartitionedPools{}

func NewMsgCreatePartitionedPools(
    creator string,
    index string,
    
) *MsgCreatePartitionedPools {
  return &MsgCreatePartitionedPools{
		Creator : creator,
		Index: index,
		
	}
}

func (msg *MsgCreatePartitionedPools) Route() string {
  return RouterKey
}

func (msg *MsgCreatePartitionedPools) Type() string {
  return TypeMsgCreatePartitionedPools
}

func (msg *MsgCreatePartitionedPools) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePartitionedPools) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePartitionedPools) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

var _ sdk.Msg = &MsgUpdatePartitionedPools{}

func NewMsgUpdatePartitionedPools(
    creator string,
    index string,
    
) *MsgUpdatePartitionedPools {
  return &MsgUpdatePartitionedPools{
		Creator: creator,
        Index: index,
        
	}
}

func (msg *MsgUpdatePartitionedPools) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePartitionedPools) Type() string {
  return TypeMsgUpdatePartitionedPools
}

func (msg *MsgUpdatePartitionedPools) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePartitionedPools) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePartitionedPools) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
   return nil
}

var _ sdk.Msg = &MsgDeletePartitionedPools{}

func NewMsgDeletePartitionedPools(
    creator string,
    index string,
    
) *MsgDeletePartitionedPools {
  return &MsgDeletePartitionedPools{
		Creator: creator,
		Index: index,
        
	}
}
func (msg *MsgDeletePartitionedPools) Route() string {
  return RouterKey
}

func (msg *MsgDeletePartitionedPools) Type() string {
  return TypeMsgDeletePartitionedPools
}

func (msg *MsgDeletePartitionedPools) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePartitionedPools) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePartitionedPools) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  }
  return nil
}