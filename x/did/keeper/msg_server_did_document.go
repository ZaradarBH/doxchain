package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocumentRequest) (result *types.MsgCreateDidDocumentResponse, err error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.Keeper.GetDidDocument(ctx, msg.DidDocument.Id.GetW3CIdentifier())

	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DidDocument already exists in store")
	}

	k.Keeper.SetDidDocument(ctx, msg.DidDocument)

	result.DidDocumentW3CIdentifier = msg.DidDocument.Id.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocumentRequest) (result *types.MsgUpdateDidDocumentResponse, err error) {
	if msg.Creator != msg.DidDocument.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "impersonation is not allowed")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDidDocument(ctx, msg.DidDocument.Id.GetW3CIdentifier())

	if found && msg.Creator != match.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "illigal update attempt")
	}

	k.Keeper.SetDidDocument(ctx, msg.DidDocument)

	result.DidDocumentW3CIdentifier = msg.DidDocument.Id.GetW3CIdentifier()
	
	return result, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (result *types.MsgDeleteDidDocumentResponse, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	match, found := k.Keeper.GetDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "DidDocument does not exist in store")
	}

	if msg.Creator != match.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "illigal delete attempt")
	}

	k.Keeper.RemoveDidDocument(ctx, msg.DidDocumentW3CIdentifier)

	return result, nil
}
