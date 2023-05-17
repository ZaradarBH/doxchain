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
	cdc.RegisterConcrete(&MsgCreateClientRegistration{}, "idp/CreateClientRegistration", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistration{}, "idp/UpdateClientRegistration", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistration{}, "idp/DeleteClientRegistration", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRelationshipRequest{}, "idp/CreateClientRegistrationRelationship", nil)
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
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateClientRegistration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteClientRegistration{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistrationRelationshipRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
