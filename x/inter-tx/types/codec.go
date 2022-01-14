package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	// this line is used by starport scaffolding # 1
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgRegisterAccount{}, "intertx/MsgRegisterAccount", nil)
	cdc.RegisterConcrete(MsgSend{}, "intertx/MsgSend", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "intertx/MsgDelegate", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSend{},
		&MsgRegisterAccount{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
