package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAMLRegistrationRequest{}, "aml/CreateAMLRegistrationRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteAMLRegistrationRequest{}, "aml/DeleteAMLRegistrationRequest", nil)
	cdc.RegisterConcrete(&MsgApproveAMLRegistrationRequest{}, "aml/ApproveAMLRegistrationRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAMLRegistrationRequest{},
		&MsgDeleteAMLRegistrationRequest{},
		&MsgApproveAMLRegistrationRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
