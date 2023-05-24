package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDidDocumentRequest = "create_did_document"
const TypeMsgUpdateDidDocumentRequest = "update_did_document"
const TypeMsgDeleteDidDocumentRequest = "delete_did_document"

var _ sdk.Msg = &MsgCreateDidDocumentRequest{}

func NewMsgCreateDidDocumentRequest(creator string, didDocument DidDocument) *MsgCreateDidDocumentRequest {
	return &MsgCreateDidDocumentRequest{
		Creator:     creator,
		DidDocument: didDocument,
	}
}

func (msg *MsgCreateDidDocumentRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateDidDocumentRequest) Type() string {
	return TypeMsgCreateDidDocumentRequest
}

func (msg *MsgCreateDidDocumentRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDidDocumentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDidDocumentRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDidDocumentRequest{}

func NewMsgUpdateDidDocumentRequest(creator string, didDocument DidDocument) *MsgUpdateDidDocumentRequest {
	return &MsgUpdateDidDocumentRequest{
		Creator:     creator,
		DidDocument: didDocument,
	}
}

func (msg *MsgUpdateDidDocumentRequest) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDidDocumentRequest) Type() string {
	return TypeMsgCreateDidDocumentRequest
}

func (msg *MsgUpdateDidDocumentRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDidDocumentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDidDocumentRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDidDocumentRequest{}

func NewMsgDeleteDidDocumentRequest(creator string, didDocumentW3CIdentifier string) *MsgDeleteDidDocumentRequest {
	return &MsgDeleteDidDocumentRequest{
		Creator:                  creator,
		DidDocumentW3CIdentifier: didDocumentW3CIdentifier,
	}
}

func (msg *MsgDeleteDidDocumentRequest) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDidDocumentRequest) Type() string {
	return TypeMsgDeleteDidDocumentRequest
}

func (msg *MsgDeleteDidDocumentRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDidDocumentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)

	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDidDocumentRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
