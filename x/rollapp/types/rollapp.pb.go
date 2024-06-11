// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/rollapp.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisAccount is a struct for the genesis account for the rollapp
type GenesisAccount struct {
	// amount of coins to be sent to the genesis address
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
	// address is a bech-32 address of the genesis account
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *GenesisAccount) Reset()         { *m = GenesisAccount{} }
func (m *GenesisAccount) String() string { return proto.CompactTextString(m) }
func (*GenesisAccount) ProtoMessage()    {}
func (*GenesisAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c072320fdc0abd9, []int{0}
}
func (m *GenesisAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccount.Merge(m, src)
}
func (m *GenesisAccount) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccount proto.InternalMessageInfo

func (m *GenesisAccount) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *GenesisAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// RollappGenesisState is a partial repr of the state the hub can expect the rollapp to be in upon genesis
type RollappGenesisState struct {
	// genesis_accounts is a list of token allocations
	GenesisAccounts []*GenesisAccount `protobuf:"bytes,1,rep,name=genesis_accounts,json=genesisAccounts,proto3" json:"genesis_accounts,omitempty"`
	// is_genesis_event is a boolean that indicates if the genesis event has occured
	IsGenesisEvent bool `protobuf:"varint,2,opt,name=is_genesis_event,json=isGenesisEvent,proto3" json:"is_genesis_event,omitempty"`
}

func (m *RollappGenesisState) Reset()         { *m = RollappGenesisState{} }
func (m *RollappGenesisState) String() string { return proto.CompactTextString(m) }
func (*RollappGenesisState) ProtoMessage()    {}
func (*RollappGenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c072320fdc0abd9, []int{1}
}
func (m *RollappGenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappGenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappGenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappGenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappGenesisState.Merge(m, src)
}
func (m *RollappGenesisState) XXX_Size() int {
	return m.Size()
}
func (m *RollappGenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappGenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_RollappGenesisState proto.InternalMessageInfo

func (m *RollappGenesisState) GetGenesisAccounts() []*GenesisAccount {
	if m != nil {
		return m.GenesisAccounts
	}
	return nil
}

func (m *RollappGenesisState) GetIsGenesisEvent() bool {
	if m != nil {
		return m.IsGenesisEvent
	}
	return false
}

// Rollapp defines a rollapp object. First the RollApp is created and then
// sequencers can be created and attached. The RollApp is identified by rollappId
type Rollapp struct {
	// The unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// creator is the bech32-encoded address of the rollapp creator.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// version is the software and configuration version.
	// starts from 1 and increases by one on every MsgUpdateState
	Version uint64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// maxSequencers is the maximum number of sequencers.
	MaxSequencers uint64 `protobuf:"varint,4,opt,name=maxSequencers,proto3" json:"maxSequencers,omitempty"`
	// permissionedAddresses is a bech32-encoded address list of the sequencers that are allowed to serve this rollappId.
	// In the case of an empty list, the rollapp is considered permissionless.
	PermissionedAddresses []string `protobuf:"bytes,5,rep,name=permissionedAddresses,proto3" json:"permissionedAddresses,omitempty"`
	// tokenMetadata is a list of TokenMetadata that are registered on this rollapp
	TokenMetadata []*TokenMetadata `protobuf:"bytes,6,rep,name=tokenMetadata,proto3" json:"tokenMetadata,omitempty"`
	// genesis_state is a partial repr of the state the hub can expect the rollapp to be in upon genesis
	GenesisState RollappGenesisState `protobuf:"bytes,7,opt,name=genesis_state,json=genesisState,proto3" json:"genesis_state"`
	// channel_id will be set to the canonical IBC channel of the rollapp.
	ChannelId string `protobuf:"bytes,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// frozen is a boolean that indicates if the rollapp is frozen.
	Frozen bool `protobuf:"varint,9,opt,name=frozen,proto3" json:"frozen,omitempty"`
	// registeredDenoms is a list of registered denom bases on this rollapp
	RegisteredDenoms []string `protobuf:"bytes,10,rep,name=registeredDenoms,proto3" json:"registeredDenoms,omitempty"`
}

func (m *Rollapp) Reset()         { *m = Rollapp{} }
func (m *Rollapp) String() string { return proto.CompactTextString(m) }
func (*Rollapp) ProtoMessage()    {}
func (*Rollapp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c072320fdc0abd9, []int{2}
}
func (m *Rollapp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rollapp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rollapp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rollapp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rollapp.Merge(m, src)
}
func (m *Rollapp) XXX_Size() int {
	return m.Size()
}
func (m *Rollapp) XXX_DiscardUnknown() {
	xxx_messageInfo_Rollapp.DiscardUnknown(m)
}

var xxx_messageInfo_Rollapp proto.InternalMessageInfo

func (m *Rollapp) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *Rollapp) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Rollapp) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Rollapp) GetMaxSequencers() uint64 {
	if m != nil {
		return m.MaxSequencers
	}
	return 0
}

func (m *Rollapp) GetPermissionedAddresses() []string {
	if m != nil {
		return m.PermissionedAddresses
	}
	return nil
}

func (m *Rollapp) GetTokenMetadata() []*TokenMetadata {
	if m != nil {
		return m.TokenMetadata
	}
	return nil
}

func (m *Rollapp) GetGenesisState() RollappGenesisState {
	if m != nil {
		return m.GenesisState
	}
	return RollappGenesisState{}
}

func (m *Rollapp) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *Rollapp) GetFrozen() bool {
	if m != nil {
		return m.Frozen
	}
	return false
}

func (m *Rollapp) GetRegisteredDenoms() []string {
	if m != nil {
		return m.RegisteredDenoms
	}
	return nil
}

// Rollapp summary is a compact representation of Rollapp
type RollappSummary struct {
	// The unique identifier of the rollapp chain.
	// The rollappId follows the same standard as cosmos chain_id.
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
	// Defines the index of the last rollapp UpdateState.
	LatestStateIndex *StateInfoIndex `protobuf:"bytes,2,opt,name=latestStateIndex,proto3" json:"latestStateIndex,omitempty"`
	// Defines the index of the last rollapp UpdateState that was finalized.
	LatestFinalizedStateIndex *StateInfoIndex `protobuf:"bytes,3,opt,name=latestFinalizedStateIndex,proto3" json:"latestFinalizedStateIndex,omitempty"`
}

func (m *RollappSummary) Reset()         { *m = RollappSummary{} }
func (m *RollappSummary) String() string { return proto.CompactTextString(m) }
func (*RollappSummary) ProtoMessage()    {}
func (*RollappSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c072320fdc0abd9, []int{3}
}
func (m *RollappSummary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappSummary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappSummary.Merge(m, src)
}
func (m *RollappSummary) XXX_Size() int {
	return m.Size()
}
func (m *RollappSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RollappSummary proto.InternalMessageInfo

func (m *RollappSummary) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *RollappSummary) GetLatestStateIndex() *StateInfoIndex {
	if m != nil {
		return m.LatestStateIndex
	}
	return nil
}

func (m *RollappSummary) GetLatestFinalizedStateIndex() *StateInfoIndex {
	if m != nil {
		return m.LatestFinalizedStateIndex
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisAccount)(nil), "dymensionxyz.dymension.rollapp.GenesisAccount")
	proto.RegisterType((*RollappGenesisState)(nil), "dymensionxyz.dymension.rollapp.RollappGenesisState")
	proto.RegisterType((*Rollapp)(nil), "dymensionxyz.dymension.rollapp.Rollapp")
	proto.RegisterType((*RollappSummary)(nil), "dymensionxyz.dymension.rollapp.RollappSummary")
}

func init() { proto.RegisterFile("dymension/rollapp/rollapp.proto", fileDescriptor_2c072320fdc0abd9) }

var fileDescriptor_2c072320fdc0abd9 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x8e, 0x49, 0x49, 0x9b, 0x2d, 0x2d, 0xd1, 0xf2, 0xd0, 0xb6, 0x2a, 0x6e, 0x14, 0x71, 0xb0,
	0x90, 0xb0, 0xd5, 0x87, 0xc4, 0xb9, 0xe5, 0xa5, 0x1e, 0xe0, 0xe0, 0x70, 0xa1, 0x07, 0xa2, 0x8d,
	0x3d, 0x75, 0x57, 0xb5, 0x77, 0xc3, 0xee, 0x26, 0x4a, 0xfa, 0x2b, 0x38, 0xf3, 0x8b, 0x7a, 0xa3,
	0x47, 0x4e, 0x08, 0x25, 0x7f, 0x82, 0x23, 0xf2, 0x7a, 0x9d, 0x26, 0x4a, 0x21, 0x88, 0x93, 0xfd,
	0xcd, 0xe3, 0x9b, 0x99, 0x6f, 0x46, 0x8b, 0x76, 0xe3, 0x51, 0x06, 0x5c, 0x31, 0xc1, 0x03, 0x29,
	0xd2, 0x94, 0xf6, 0x7a, 0xe5, 0xd7, 0xef, 0x49, 0xa1, 0x05, 0x76, 0xa7, 0x01, 0xc3, 0xd1, 0xa5,
	0x3f, 0x05, 0xbe, 0x8d, 0xda, 0x7e, 0x98, 0x88, 0x44, 0x98, 0xd0, 0x20, 0xff, 0x2b, 0xb2, 0xb6,
	0x5b, 0x8b, 0xb4, 0x4a, 0x53, 0x0d, 0x1d, 0xc6, 0xcf, 0xca, 0x98, 0x9d, 0xc5, 0x98, 0x2e, 0xe5,
	0x17, 0xd6, 0xeb, 0x46, 0x42, 0x65, 0x42, 0x05, 0x5d, 0xaa, 0x20, 0x18, 0xec, 0x75, 0x41, 0xd3,
	0xbd, 0x20, 0x12, 0x8c, 0x17, 0xfe, 0x56, 0x84, 0x36, 0xdf, 0x02, 0x07, 0xc5, 0xd4, 0x51, 0x14,
	0x89, 0x3e, 0xd7, 0xf8, 0x05, 0xaa, 0xd1, 0x2c, 0xff, 0x23, 0x4e, 0xd3, 0xf1, 0xd6, 0xf7, 0xb7,
	0xfc, 0x82, 0xc2, 0xcf, 0x29, 0x7c, 0x4b, 0xe1, 0xbf, 0x14, 0x8c, 0x1f, 0xaf, 0x5c, 0xfd, 0xd8,
	0xad, 0x84, 0x36, 0x1c, 0x13, 0xb4, 0x4a, 0xe3, 0x58, 0x82, 0x52, 0xe4, 0x4e, 0xd3, 0xf1, 0xea,
	0x61, 0x09, 0x5b, 0x5f, 0x1d, 0xf4, 0x20, 0x2c, 0x7a, 0xb3, 0xc5, 0xda, 0xf9, 0x14, 0xf8, 0x23,
	0x6a, 0x24, 0x05, 0xee, 0xd0, 0xa2, 0xba, 0x22, 0x4e, 0xb3, 0xea, 0xad, 0xef, 0xfb, 0xfe, 0xdf,
	0xf5, 0xf2, 0xe7, 0x9b, 0x0e, 0xef, 0x27, 0x73, 0x58, 0x61, 0x0f, 0x35, 0x98, 0xea, 0x94, 0xec,
	0x30, 0x00, 0xae, 0x4d, 0x57, 0x6b, 0xe1, 0x26, 0x53, 0x36, 0xf9, 0x75, 0x6e, 0x6d, 0x7d, 0xab,
	0xa2, 0x55, 0xdb, 0x1c, 0xde, 0x41, 0x75, 0x5b, 0xe0, 0x24, 0x36, 0xe3, 0xd7, 0xc3, 0x1b, 0x43,
	0x3e, 0x60, 0x24, 0x81, 0x6a, 0x21, 0xcb, 0x01, 0x2d, 0xcc, 0x3d, 0x03, 0x90, 0x79, 0x83, 0xa4,
	0xda, 0x74, 0xbc, 0x95, 0xb0, 0x84, 0xf8, 0x29, 0xda, 0xc8, 0xe8, 0xb0, 0x0d, 0x9f, 0xfb, 0xc0,
	0x23, 0x90, 0x8a, 0xac, 0x18, 0xff, 0xbc, 0x11, 0x1f, 0xa2, 0x47, 0x3d, 0x90, 0x19, 0x53, 0x79,
	0x0e, 0xc4, 0x47, 0x85, 0x6e, 0xa0, 0xc8, 0xdd, 0x66, 0xd5, 0xab, 0x87, 0xb7, 0x3b, 0x71, 0x1b,
	0x6d, 0x68, 0x71, 0x01, 0xfc, 0x1d, 0x68, 0x1a, 0x53, 0x4d, 0x49, 0xcd, 0x68, 0xf7, 0x7c, 0x99,
	0x76, 0x1f, 0x66, 0x93, 0xc2, 0x79, 0x0e, 0xfc, 0x09, 0x6d, 0x94, 0xaa, 0x99, 0x53, 0x23, 0xab,
	0xe6, 0x0a, 0x0e, 0x96, 0x91, 0xde, 0xb2, 0x5f, 0x7b, 0x1f, 0xf7, 0x92, 0xd9, 0x9d, 0x3f, 0x41,
	0x28, 0x3a, 0xa7, 0x9c, 0x43, 0xda, 0x61, 0x31, 0x59, 0x2b, 0x34, 0xb6, 0x96, 0x93, 0x18, 0x3f,
	0x46, 0xb5, 0x33, 0x29, 0x2e, 0x81, 0x93, 0xba, 0xd9, 0x96, 0x45, 0xf8, 0x19, 0x6a, 0x48, 0x48,
	0x98, 0xd2, 0x20, 0x21, 0x7e, 0x05, 0x5c, 0x64, 0x8a, 0x20, 0x23, 0xce, 0x82, 0xbd, 0xf5, 0xcb,
	0x41, 0x9b, 0xb6, 0x9d, 0x76, 0x3f, 0xcb, 0xa8, 0x1c, 0x2d, 0x59, 0xec, 0x29, 0x6a, 0xa4, 0x54,
	0x83, 0xd2, 0xa6, 0xc5, 0x13, 0x1e, 0xc3, 0xd0, 0x6c, 0xf8, 0x1f, 0xee, 0xd0, 0x66, 0x9c, 0x09,
	0x93, 0x15, 0x2e, 0xf0, 0xe0, 0x14, 0x6d, 0x15, 0xb6, 0x37, 0x8c, 0xd3, 0x94, 0x5d, 0x42, 0x3c,
	0x53, 0xa4, 0xfa, 0x5f, 0x45, 0xfe, 0x4c, 0x78, 0xfc, 0xfe, 0x6a, 0xec, 0x3a, 0xd7, 0x63, 0xd7,
	0xf9, 0x39, 0x76, 0x9d, 0x2f, 0x13, 0xb7, 0x72, 0x3d, 0x71, 0x2b, 0xdf, 0x27, 0x6e, 0xe5, 0xf4,
	0x30, 0x61, 0xfa, 0xbc, 0xdf, 0xf5, 0x23, 0x91, 0x05, 0xb3, 0xe5, 0x6e, 0x40, 0x30, 0x38, 0x08,
	0x86, 0xd3, 0x37, 0x44, 0x8f, 0x7a, 0xa0, 0xba, 0x35, 0xf3, 0x4a, 0x1c, 0xfc, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x75, 0x02, 0x7c, 0xc2, 0xe0, 0x04, 0x00, 0x00,
}

func (m *GenesisAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRollapp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RollappGenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappGenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappGenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsGenesisEvent {
		i--
		if m.IsGenesisEvent {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.GenesisAccounts) > 0 {
		for iNdEx := len(m.GenesisAccounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisAccounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRollapp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Rollapp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rollapp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rollapp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RegisteredDenoms) > 0 {
		for iNdEx := len(m.RegisteredDenoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RegisteredDenoms[iNdEx])
			copy(dAtA[i:], m.RegisteredDenoms[iNdEx])
			i = encodeVarintRollapp(dAtA, i, uint64(len(m.RegisteredDenoms[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if m.Frozen {
		i--
		if m.Frozen {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x42
	}
	{
		size, err := m.GenesisState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRollapp(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.TokenMetadata) > 0 {
		for iNdEx := len(m.TokenMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRollapp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PermissionedAddresses) > 0 {
		for iNdEx := len(m.PermissionedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PermissionedAddresses[iNdEx])
			copy(dAtA[i:], m.PermissionedAddresses[iNdEx])
			i = encodeVarintRollapp(dAtA, i, uint64(len(m.PermissionedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.MaxSequencers != 0 {
		i = encodeVarintRollapp(dAtA, i, uint64(m.MaxSequencers))
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintRollapp(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RollappSummary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappSummary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappSummary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestFinalizedStateIndex != nil {
		{
			size, err := m.LatestFinalizedStateIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRollapp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LatestStateIndex != nil {
		{
			size, err := m.LatestStateIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRollapp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintRollapp(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRollapp(dAtA []byte, offset int, v uint64) int {
	offset -= sovRollapp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovRollapp(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	return n
}

func (m *RollappGenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GenesisAccounts) > 0 {
		for _, e := range m.GenesisAccounts {
			l = e.Size()
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	if m.IsGenesisEvent {
		n += 2
	}
	return n
}

func (m *Rollapp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovRollapp(uint64(m.Version))
	}
	if m.MaxSequencers != 0 {
		n += 1 + sovRollapp(uint64(m.MaxSequencers))
	}
	if len(m.PermissionedAddresses) > 0 {
		for _, s := range m.PermissionedAddresses {
			l = len(s)
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	if len(m.TokenMetadata) > 0 {
		for _, e := range m.TokenMetadata {
			l = e.Size()
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	l = m.GenesisState.Size()
	n += 1 + l + sovRollapp(uint64(l))
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.Frozen {
		n += 2
	}
	if len(m.RegisteredDenoms) > 0 {
		for _, s := range m.RegisteredDenoms {
			l = len(s)
			n += 1 + l + sovRollapp(uint64(l))
		}
	}
	return n
}

func (m *RollappSummary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.LatestStateIndex != nil {
		l = m.LatestStateIndex.Size()
		n += 1 + l + sovRollapp(uint64(l))
	}
	if m.LatestFinalizedStateIndex != nil {
		l = m.LatestFinalizedStateIndex.Size()
		n += 1 + l + sovRollapp(uint64(l))
	}
	return n
}

func sovRollapp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRollapp(x uint64) (n int) {
	return sovRollapp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RollappGenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RollappGenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappGenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisAccounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisAccounts = append(m.GenesisAccounts, &GenesisAccount{})
			if err := m.GenesisAccounts[len(m.GenesisAccounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsGenesisEvent", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsGenesisEvent = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Rollapp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rollapp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rollapp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSequencers", wireType)
			}
			m.MaxSequencers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSequencers |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PermissionedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PermissionedAddresses = append(m.PermissionedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenMetadata = append(m.TokenMetadata, &TokenMetadata{})
			if err := m.TokenMetadata[len(m.TokenMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Frozen = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredDenoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisteredDenoms = append(m.RegisteredDenoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RollappSummary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRollapp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RollappSummary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappSummary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestStateIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestStateIndex == nil {
				m.LatestStateIndex = &StateInfoIndex{}
			}
			if err := m.LatestStateIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestFinalizedStateIndex", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRollapp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRollapp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestFinalizedStateIndex == nil {
				m.LatestFinalizedStateIndex = &StateInfoIndex{}
			}
			if err := m.LatestFinalizedStateIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRollapp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRollapp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRollapp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRollapp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRollapp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRollapp
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRollapp
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRollapp
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRollapp        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRollapp          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRollapp = fmt.Errorf("proto: unexpected end of group")
)
