package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTokenRequest{}, "oauthtwo/token", nil)
	cdc.RegisterConcrete(&MsgDeviceCodeRequest{}, "oauthtwo/devicecode", nil)
	cdc.RegisterConcrete(&MsgAuthorizeRequest{}, "oauthtwo/Authorize", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTokenRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeviceCodeRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthorizeRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)