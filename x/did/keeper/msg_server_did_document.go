package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDidDocument(goCtx context.Context, msg *types.MsgCreateDidDocumentRequest) (*types.MsgCreateDidDocumentResponse, error) {
	err := k.Keeper.SetDidDocument(sdk.UnwrapSDKContext(goCtx), msg.DidDocument, false)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateDidDocumentResponse{
		FullyQualifiedDidIdentifier: msg.DidDocument.Id.GetFullyQualifiedDidIdentifier(),
	}, nil
}

func (k msgServer) UpdateDidDocument(goCtx context.Context, msg *types.MsgUpdateDidDocumentRequest) (*types.MsgUpdateDidDocumentResponse, error) {
	err := k.SetDidDocument(sdk.UnwrapSDKContext(goCtx), msg.DidDocument, true)

	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateDidDocumentResponse{}, nil
}

func (k msgServer) DeleteDidDocument(goCtx context.Context, msg *types.MsgDeleteDidDocumentRequest) (*types.MsgDeleteDidDocumentResponse, error) {
	err := k.Keeper.RemoveDidDocument(sdk.UnwrapSDKContext(goCtx), msg.FullyQualifiedDidIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteDidDocumentResponse{}, nil
}
