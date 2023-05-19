package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAuthenticationRequest{}, "idp/Login", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRegistryRequest{}, "idp/CreateClientRegistrationRegistry", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistrationRegistryRequest{}, "idp/UpdateClientRegistrationRegistry", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRegistryRequest{}, "idp/DeleteClientRegistrationRegistry", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRequest{}, "idp/CreateClientRegistration", nil)
	cdc.RegisterConcrete(&MsgUpdateClientRegistrationRequest{}, "idp/UpdateClientRegistration", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRequest{}, "idp/DeleteClientRegistration", nil)
	cdc.RegisterConcrete(&MsgCreateClientRegistrationRelationshipRequest{}, "idp/CreateClientRegistrationRelationship", nil)
	cdc.RegisterConcrete(&MsgDeleteClientRegistrationRelationshipRequest{}, "idp/DeleteClientRegistrationRelationship", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthenticationRequest{},
		&MsgCreateClientRegistrationRegistryRequest{},
		&MsgUpdateClientRegistrationRegistryRequest{},
		&MsgDeleteClientRegistrationRegistryRequest{},
		&MsgCreateClientRegistrationRequest{},
		&MsgUpdateClientRegistrationRequest{},
		&MsgDeleteClientRegistrationRequest{},
		&MsgCreateClientRegistrationRelationshipRequest{},
		&MsgDeleteClientRegistrationRelationshipRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
