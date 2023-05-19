package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateClientRegistrationRelationshipRequest = "create_client_registration_relationship"

var _ sdk.Msg = &MsgCreateClientRegistrationRelationshipRequest{}

func NewMsgCreateClientRegistrationRelationshipRequest(clientRegistrationRelationshipRegistryEntry ClientRegistrationRelationshipRegistryEntry) *MsgCreateClientRegistrationRelationshipRequest {
	return &MsgCreateClientRegistrationRelationshipRequest{
		ClientRegistrationRelationshipRegistryEntry: clientRegistrationRelationshipRegistryEntry,
	}
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) Type() string {
	return TypeMsgCreateClientRegistrationRelationshipRequest
}

func (msg *MsgCreateClientRegistrationRelationshipRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRelationshipRegistryEntry.Owner.Creator)

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
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRelationshipRegistryEntry.Owner.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	return nil
}
