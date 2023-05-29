package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type msgServer struct {
	BaseKeeper
}

func NewMsgServerImpl(keeper BaseKeeper) banktypes.MsgServer {
	return &msgServer{BaseKeeper: keeper}
}

var _ banktypes.MsgServer = msgServer{}

func (k msgServer) Send(goCtx context.Context, msg *banktypes.MsgSend) (*banktypes.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.BaseSendKeeper.IsSendEnabledCoins(ctx, msg.Amount...); err != nil {
		return nil, err
	}

	if err := k.abs.AddToWatchlist(ctx, sdk.AccAddress(msg.FromAddress), msg.Amount); err != nil {
		return nil, err
	}

	return bankkeeper.NewMsgServerImpl(k.BaseKeeper).Send(sdk.WrapSDKContext(ctx), msg)
}

func (k msgServer) MultiSend(goCtx context.Context, msg *banktypes.MsgMultiSend) (*banktypes.MsgMultiSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, input := range msg.Inputs {
		if err := k.BaseSendKeeper.IsSendEnabledCoins(ctx, input.Coins...); err != nil {
			return nil, err
		}
	}

	for _, input := range msg.Inputs {
		if err := k.abs.AddToWatchlist(ctx, sdk.AccAddress(input.Address), input.Coins); err != nil {
			return nil, err
		}
	}

	return bankkeeper.NewMsgServerImpl(k.BaseKeeper).MultiSend(sdk.WrapSDKContext(ctx), msg)
}
