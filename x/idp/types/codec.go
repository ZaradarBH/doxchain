package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAuthenticationRequest{}, "idp/Login", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistry{}, "idp/CreateClientRegistry", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistry{}, "idp/UpdateClientRegistry", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistry{}, "idp/DeleteClientRegistry", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthenticationRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistry{},
		&MsgUpdateClientRegistry{},
		&MsgDeleteClientRegistry{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
