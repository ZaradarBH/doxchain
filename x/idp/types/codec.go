package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAuthenticationRequest{}, "idp/Login", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRegistryRequest{}, "idp/CreateClientRegistrationRegistryRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistrationRegistryRequest{}, "idp/UpdateClientRegistrationRegistryRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRegistryRequest{}, "idp/DeleteClientRegistrationRegistryRequest", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRequest{}, "idp/CreateClientRegistrationRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistrationRequest{}, "idp/UpdateClientRegistrationRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRequest{}, "idp/DeleteClientRegistrationRequest", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRelationshipRequest{}, "idp/CreateClientRegistrationRelationshipRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRelationshipRequest{}, "idp/DeleteClientRegistrationRelationshipRequest", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthenticationRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistrationRegistryRequest{},
		&MsgUpdateClientRegistrationRegistryRequest{},
		&MsgDeleteClientRegistrationRegistryRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistrationRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateClientRegistrationRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteClientRegistrationRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateClientRegistrationRelationshipRequest{},
		&MsgDeleteClientRegistrationRelationshipRequest{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
