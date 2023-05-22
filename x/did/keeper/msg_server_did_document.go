package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocumentRequest) (*types.MsgCreateDidDocumentResponse, error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	err := k.Keeper.SetDidDocument(sdk.UnwrapSDKContext(goCtx), msg.DidDocument, false)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateDidDocumentResponse{
		DidDocumentW3CIdentifier: msg.DidDocument.Id.GetW3CIdentifier(),
	}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocumentRequest) (*types.MsgUpdateDidDocumentResponse, error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	err := k.SetDidDocument(sdk.UnwrapSDKContext(goCtx), msg.DidDocument, true)

	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateDidDocumentResponse{
		DidDocumentW3CIdentifier: msg.DidDocument.Id.GetW3CIdentifier(),
	}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	if msg.Creator != valFound.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err := k.Keeper.RemoveDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
