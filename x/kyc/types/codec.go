package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateKYCRegistrationRequest{}, "kyc/CreateKYCRegistration", nil)
	cdc.RegisterConcrete(&MsgDeleteKYCRegistrationRequest{}, "kyc/DeleteKYCRegistration", nil)
	cdc.RegisterConcrete(&MsgApproveKYCRegistrationRequest{}, "kyc/ApproveKYCRegistration", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateKYCRegistrationRequest{},
		&MsgDeleteKYCRegistrationRequest{},
		&MsgApproveKYCRegistrationRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
