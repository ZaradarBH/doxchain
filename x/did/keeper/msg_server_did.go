package keeper

import (
	"context"

	"github.com/be-heroes/doxchain/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateDid(goCtx context.Context, msg *types.MsgCreateDidRequest) (*types.MsgCreateDidResponse, error) {
	err := k.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, false)

	if err != nil {
		return nil, err
	}

	return &types.MsgCreateDidResponse{
		FullyQualifiedDidIdentifier: msg.Did.GetFullyQualifiedDidIdentifier(),
	}, nil
}

func (k msgServer) UpdateDid(goCtx context.Context, msg *types.MsgUpdateDidRequest) (*types.MsgUpdateDidResponse, error) {
	err := k.SetDid(sdk.UnwrapSDKContext(goCtx), msg.Did, true)

	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateDidResponse{}, nil
}

func (k msgServer) DeleteDid(goCtx context.Context, msg *types.MsgDeleteDidRequest) (*types.MsgDeleteDidResponse, error) {
	err := k.RemoveDid(sdk.UnwrapSDKContext(goCtx), msg.FullyQualifiedDidIdentifier)

	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteDidResponse{}, nil
}
