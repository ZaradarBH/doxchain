package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	didTypes "github.com/be-heroes/doxchain/x/did/types"
)

const TypeMsgCreateClientRegistrationRelationshipRequest = "create_client_registration_relationship"

var _ sdk.Msg = &MsgCreateClientRegistrationRelationshipRequest{}

func NewMsgCreateClientRegistrationRelationshipRequest(ownerId didTypes.Did, destinationId didTypes.Did, accessClientList AccessClientList) *MsgCreateClientRegistrationRelationshipRequest {
	return &MsgCreateClientRegistrationRelationshipRequest{
		OwnerId: ownerId,
		DestinationId: destinationId,
		AccessClientList: accessClientList,
	}
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) Type() string {
	return TypeMsgCreateClientRegistrationRelationshipRequest
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.OwnerId.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerId.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
