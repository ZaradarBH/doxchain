package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgTokenRequest{}, "oauthtwo/Token", nil)
	cdc.RegisterConcrete(&MsgDeviceCodeRequest{}, "oauthtwo/DeviceCode", nil)
	cdc.RegisterConcrete(&MsgAuthorizeRequest{}, "oauthtwo/Authorize", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTokenRequest{},
		&MsgDeviceCodeRequest{},
		&MsgAuthorizeRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
