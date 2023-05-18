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
		FullyQualifiedDidIdentifier: msg.DidDocument.Id.GetW3CIdentifier(),
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

	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Check if the value exists
	valFound, isFound := k.GetDidDocument(ctx, msg.FullyQualifiedDidIdentifier)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	
	err := k.Keeper.RemoveDidDocument(ctx, msg.FullyQualifiedDidIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
