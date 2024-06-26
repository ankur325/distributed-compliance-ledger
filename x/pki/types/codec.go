package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposeAddX509RootCert{}, "pki/ProposeAddX509RootCert", nil)
	cdc.RegisterConcrete(&MsgApproveAddX509RootCert{}, "pki/ApproveAddX509RootCert", nil)
	cdc.RegisterConcrete(&MsgAddX509Cert{}, "pki/AddX509Cert", nil)
	cdc.RegisterConcrete(&MsgProposeRevokeX509RootCert{}, "pki/ProposeRevokeX509RootCert", nil)
	cdc.RegisterConcrete(&MsgApproveRevokeX509RootCert{}, "pki/ApproveRevokeX509RootCert", nil)
	cdc.RegisterConcrete(&MsgRevokeX509Cert{}, "pki/RevokeX509Cert", nil)
	cdc.RegisterConcrete(&MsgRejectAddX509RootCert{}, "pki/RejectAddX509RootCert", nil)
	cdc.RegisterConcrete(&MsgAddPkiRevocationDistributionPoint{}, "pki/AddPkiRevocationDistributionPoint", nil)
	cdc.RegisterConcrete(&MsgUpdatePkiRevocationDistributionPoint{}, "pki/UpdatePkiRevocationDistributionPoint", nil)
	cdc.RegisterConcrete(&MsgDeletePkiRevocationDistributionPoint{}, "pki/DeletePkiRevocationDistributionPoint", nil)
	cdc.RegisterConcrete(&MsgAssignVid{}, "pki/AssignVid", nil)
	cdc.RegisterConcrete(&MsgAddNocX509RootCert{}, "pki/AddNocX509RootCert", nil)
	cdc.RegisterConcrete(&MsgRemoveX509Cert{}, "pki/RemoveX509Cert", nil)
	cdc.RegisterConcrete(&MsgAddNocX509IcaCert{}, "pki/AddNocX509Cert", nil)
	cdc.RegisterConcrete(&MsgRevokeNocX509RootCert{}, "pki/RevokeNocRootX509Cert", nil)
	cdc.RegisterConcrete(&MsgRevokeNocX509IcaCert{}, "pki/RevokeNocX509Cert", nil)
	cdc.RegisterConcrete(&MsgRemoveNocX509IcaCert{}, "pki/RemoveNocX509IcaCert", nil)
	cdc.RegisterConcrete(&MsgRemoveNocX509RootCert{}, "pki/RemoveNocX509RootCert", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeAddX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveAddX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddX509Cert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposeRevokeX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveRevokeX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeX509Cert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRejectAddX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddPkiRevocationDistributionPoint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePkiRevocationDistributionPoint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeletePkiRevocationDistributionPoint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAssignVid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddNocX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveX509Cert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddNocX509IcaCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeNocX509RootCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevokeNocX509IcaCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveNocX509IcaCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveNocX509RootCert{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc) //nolint:nosnakecase
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
