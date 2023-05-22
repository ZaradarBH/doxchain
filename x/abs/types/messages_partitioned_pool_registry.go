package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePartitionedPoolRegistry = "create_partitioned_pool_registry"
)

var _ sdk.Msg = &MsgCreatePartitionedPoolRegistryRequest{}

func NewMsgCreatePartitionedPoolRegistryRequest(
	creator string,
	denom string,

) *MsgCreatePartitionedPoolRegistryRequest {
	return &MsgCreatePartitionedPoolRegistryRequest{
		Creator: creator,
		Denom:   denom,
	}
}

func (msg *MsgCreatePartitionedPoolRegistryRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreatePartitionedPoolRegistryRequest) Type() string {
	return TypeMsgCreatePartitionedPoolRegistry
}

func (msg *MsgCreatePartitionedPoolRegistryRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePartitionedPoolRegistryRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePartitionedPoolRegistryRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
