package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendIbcPrice{}, "clprice/SendIbcPrice", nil)

	cdc.RegisterConcrete(&MsgCreateSentPrice{}, "clprice/CreateSentPrice", nil)
	cdc.RegisterConcrete(&MsgUpdateSentPrice{}, "clprice/UpdateSentPrice", nil)
	cdc.RegisterConcrete(&MsgDeleteSentPrice{}, "clprice/DeleteSentPrice", nil)

	cdc.RegisterConcrete(&MsgCreatePrice{}, "clprice/CreatePrice", nil)
	cdc.RegisterConcrete(&MsgUpdatePrice{}, "clprice/UpdatePrice", nil)
	cdc.RegisterConcrete(&MsgDeletePrice{}, "clprice/DeletePrice", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendIbcPrice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentPrice{},
		&MsgUpdateSentPrice{},
		&MsgDeleteSentPrice{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePrice{},
		&MsgUpdatePrice{},
		&MsgDeletePrice{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
