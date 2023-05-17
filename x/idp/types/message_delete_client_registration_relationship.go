package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteClientRegistrationRelationshipRequest = "delete_client_registration_relationship"

var _ sdk.Msg = &MsgDeleteClientRegistrationRelationshipRequest{}

func NewMsgDeleteClientRegistrationRelationshipRequest(clientRegistrationRelationshipRegistryEntry ClientRegistrationRelationshipRegistryEntry) *MsgDeleteClientRegistrationRelationshipRequest {
	return &MsgDeleteClientRegistrationRelationshipRequest{
		ClientRegistrationRelationshipRegistryEntry: clientRegistrationRelationshipRegistryEntry,
	}
}

func (msg *MsgDeleteClientRegistrationRelationshipRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteClientRegistrationRelationshipRequest) Type() string {
	return TypeMsgDeleteClientRegistrationRelationshipRequest
}

func (msg *MsgDeleteClientRegistrationRelationshipRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRelationshipRegistryEntry.OwnerId.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteClientRegistrationRelationshipRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteClientRegistrationRelationshipRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ClientRegistrationRelationshipRegistryEntry.OwnerId.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
