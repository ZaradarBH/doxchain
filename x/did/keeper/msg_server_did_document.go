package keeper

import (
	"fmt"
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocumentRequest) (*types.MsgCreateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Keeper.SetDidDocument(ctx, msg.DidDocument)
	
	return &types.MsgCreateDidDocumentResponse{}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocumentRequest) (*types.MsgUpdateDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	fullyQualifiedDidIdentifier := msg.DidDocument.Id.GetFullyQualifiedDidIdentifier()
	val, found := k.GetDidDocument(ctx, fullyQualifiedDidIdentifier)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", fullyQualifiedDidIdentifier))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.DidDocument.Id.Creator != val.Id.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.Keeper.SetDidDocument(ctx, msg.DidDocument)

	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (*types.MsgDeleteDidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Keeper.RemoveDidDocument(ctx, msg.FullyQualifiedDidIdentifier)

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
