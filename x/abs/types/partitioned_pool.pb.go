// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/abs/partitioned_pool.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/idp/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
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

type PartitionedPool struct {
	Denom            string                       `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	DesiredState     *ParitionedPoolConfiguration `protobuf:"bytes,2,opt,name=desiredState,proto3" json:"desiredState,omitempty"`
	AccessClientList *types.AccessClientList      `protobuf:"bytes,3,opt,name=accessClientList,proto3" json:"accessClientList,omitempty"`
}

func (m *PartitionedPool) Reset()         { *m = PartitionedPool{} }
func (m *PartitionedPool) String() string { return proto.CompactTextString(m) }
func (*PartitionedPool) ProtoMessage()    {}
func (*PartitionedPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a552114df4ea429, []int{0}
}
func (m *PartitionedPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartitionedPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartitionedPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartitionedPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionedPool.Merge(m, src)
}
func (m *PartitionedPool) XXX_Size() int {
	return m.Size()
}
func (m *PartitionedPool) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionedPool.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionedPool proto.InternalMessageInfo

func (m *PartitionedPool) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *PartitionedPool) GetDesiredState() *ParitionedPoolConfiguration {
	if m != nil {
		return m.DesiredState
	}
	return nil
}

func (m *PartitionedPool) GetAccessClientList() *types.AccessClientList {
	if m != nil {
		return m.AccessClientList
	}
	return nil
}

type ParitionedPoolConfiguration struct {
	Threshold ParitionedPoolThreshold   `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold"`
	Corridors []*ParitionedPoolCorridor `protobuf:"bytes,2,rep,name=Corridors,proto3" json:"Corridors,omitempty"`
}

func (m *ParitionedPoolConfiguration) Reset()         { *m = ParitionedPoolConfiguration{} }
func (m *ParitionedPoolConfiguration) String() string { return proto.CompactTextString(m) }
func (*ParitionedPoolConfiguration) ProtoMessage()    {}
func (*ParitionedPoolConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a552114df4ea429, []int{1}
}
func (m *ParitionedPoolConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParitionedPoolConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParitionedPoolConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParitionedPoolConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParitionedPoolConfiguration.Merge(m, src)
}
func (m *ParitionedPoolConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *ParitionedPoolConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ParitionedPoolConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ParitionedPoolConfiguration proto.InternalMessageInfo

func (m *ParitionedPoolConfiguration) GetThreshold() ParitionedPoolThreshold {
	if m != nil {
		return m.Threshold
	}
	return ParitionedPoolThreshold{}
}

func (m *ParitionedPoolConfiguration) GetCorridors() []*ParitionedPoolCorridor {
	if m != nil {
		return m.Corridors
	}
	return nil
}

type ParitionedPoolCorridor struct {
	Threshold ParitionedPoolThreshold `protobuf:"bytes,1,opt,name=threshold,proto3" json:"threshold"`
	Target    types1.Coin             `protobuf:"bytes,2,opt,name=target,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"target" yaml:"coin"`
}

func (m *ParitionedPoolCorridor) Reset()         { *m = ParitionedPoolCorridor{} }
func (m *ParitionedPoolCorridor) String() string { return proto.CompactTextString(m) }
func (*ParitionedPoolCorridor) ProtoMessage()    {}
func (*ParitionedPoolCorridor) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a552114df4ea429, []int{2}
}
func (m *ParitionedPoolCorridor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParitionedPoolCorridor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParitionedPoolCorridor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParitionedPoolCorridor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParitionedPoolCorridor.Merge(m, src)
}
func (m *ParitionedPoolCorridor) XXX_Size() int {
	return m.Size()
}
func (m *ParitionedPoolCorridor) XXX_DiscardUnknown() {
	xxx_messageInfo_ParitionedPoolCorridor.DiscardUnknown(m)
}

var xxx_messageInfo_ParitionedPoolCorridor proto.InternalMessageInfo

func (m *ParitionedPoolCorridor) GetThreshold() ParitionedPoolThreshold {
	if m != nil {
		return m.Threshold
	}
	return ParitionedPoolThreshold{}
}

func (m *ParitionedPoolCorridor) GetTarget() types1.Coin {
	if m != nil {
		return m.Target
	}
	return types1.Coin{}
}

type ParitionedPoolThreshold struct {
	SoftCap uint64 `protobuf:"varint,1,opt,name=softCap,proto3" json:"softCap,omitempty"`
	HardCap uint64 `protobuf:"varint,2,opt,name=hardCap,proto3" json:"hardCap,omitempty"`
}

func (m *ParitionedPoolThreshold) Reset()         { *m = ParitionedPoolThreshold{} }
func (m *ParitionedPoolThreshold) String() string { return proto.CompactTextString(m) }
func (*ParitionedPoolThreshold) ProtoMessage()    {}
func (*ParitionedPoolThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a552114df4ea429, []int{3}
}
func (m *ParitionedPoolThreshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParitionedPoolThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParitionedPoolThreshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParitionedPoolThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParitionedPoolThreshold.Merge(m, src)
}
func (m *ParitionedPoolThreshold) XXX_Size() int {
	return m.Size()
}
func (m *ParitionedPoolThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_ParitionedPoolThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_ParitionedPoolThreshold proto.InternalMessageInfo

func (m *ParitionedPoolThreshold) GetSoftCap() uint64 {
	if m != nil {
		return m.SoftCap
	}
	return 0
}

func (m *ParitionedPoolThreshold) GetHardCap() uint64 {
	if m != nil {
		return m.HardCap
	}
	return 0
}

func init() {
	proto.RegisterType((*PartitionedPool)(nil), "beheroes.doxchain.abs.PartitionedPool")
	proto.RegisterType((*ParitionedPoolConfiguration)(nil), "beheroes.doxchain.abs.ParitionedPoolConfiguration")
	proto.RegisterType((*ParitionedPoolCorridor)(nil), "beheroes.doxchain.abs.ParitionedPoolCorridor")
	proto.RegisterType((*ParitionedPoolThreshold)(nil), "beheroes.doxchain.abs.ParitionedPoolThreshold")
}

func init() {
	proto.RegisterFile("doxchain/abs/partitioned_pool.proto", fileDescriptor_2a552114df4ea429)
}

var fileDescriptor_2a552114df4ea429 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xbb, 0x31, 0x54, 0x17, 0x09, 0x14, 0x0d, 0x08, 0x43, 0x4a, 0xab, 0x20, 0x44, 0x39,
	0xd4, 0xd6, 0xca, 0x8d, 0x0b, 0xa2, 0xb9, 0x82, 0x54, 0x02, 0x17, 0x10, 0x52, 0xe5, 0xc4, 0x5e,
	0x62, 0x91, 0xe4, 0x45, 0xb6, 0x87, 0xb6, 0x7f, 0xc1, 0xef, 0xe0, 0x77, 0x70, 0xd8, 0x71, 0x47,
	0x4e, 0x63, 0x6a, 0xff, 0x01, 0xbf, 0x00, 0xc5, 0x4e, 0xd7, 0x01, 0x05, 0xed, 0xb2, 0x53, 0xf2,
	0xfc, 0xbe, 0xf7, 0xbd, 0xef, 0xb3, 0xdf, 0xc3, 0x8f, 0x38, 0x1c, 0xa5, 0x39, 0x93, 0x15, 0x65,
	0x89, 0xa6, 0x35, 0x53, 0x46, 0x1a, 0x09, 0x95, 0xe0, 0xf3, 0x1a, 0xa0, 0x20, 0xb5, 0x02, 0x03,
	0xde, 0xdd, 0x44, 0xe4, 0x42, 0x81, 0xd0, 0x64, 0x85, 0x26, 0x2c, 0xd1, 0x7b, 0xbb, 0x19, 0x64,
	0x60, 0x11, 0xb4, 0xf9, 0x73, 0xe0, 0xbd, 0xc7, 0x17, 0x8c, 0x92, 0xd7, 0x94, 0xa5, 0xa9, 0xd0,
	0x7a, 0x9e, 0x16, 0x52, 0x54, 0x66, 0x5e, 0x48, 0x6d, 0x5a, 0x58, 0x90, 0x82, 0x2e, 0x41, 0xd3,
	0x84, 0x69, 0x41, 0x3f, 0xef, 0x27, 0xc2, 0xb0, 0x7d, 0x9a, 0x82, 0xac, 0x5c, 0x3e, 0x3c, 0x47,
	0xf8, 0xf6, 0x6c, 0x2d, 0x67, 0x06, 0x50, 0x78, 0xbb, 0xf8, 0x06, 0x17, 0x15, 0x94, 0x3e, 0x1a,
	0xa2, 0x51, 0x2f, 0x76, 0x81, 0xf7, 0x11, 0xdf, 0xe2, 0x42, 0x4b, 0x25, 0xf8, 0x5b, 0xc3, 0x8c,
	0xf0, 0xbb, 0x43, 0x34, 0xea, 0x4f, 0x26, 0x64, 0xa3, 0x68, 0x32, 0x63, 0xea, 0x12, 0x65, 0x04,
	0xd5, 0x81, 0xcc, 0x0e, 0x15, 0x6b, 0x8e, 0xa6, 0xdb, 0x27, 0x67, 0x03, 0x14, 0xff, 0xc6, 0xe6,
	0xbd, 0xc7, 0x77, 0x9c, 0x87, 0xc8, 0x5a, 0x78, 0x25, 0xb5, 0xf1, 0xb7, 0x6c, 0x87, 0x27, 0x1b,
	0x3a, 0x48, 0x5e, 0x93, 0x97, 0x7f, 0xc0, 0x5b, 0xda, 0xbf, 0x68, 0xc2, 0x6f, 0x08, 0x3f, 0xfc,
	0x8f, 0x1c, 0x2f, 0xc6, 0x3d, 0x93, 0x2b, 0xa1, 0x73, 0x28, 0xb8, 0xb5, 0xdc, 0x9f, 0x90, 0x2b,
	0xb9, 0x7a, 0xb7, 0xaa, 0xb2, 0xad, 0x3b, 0xf1, 0x9a, 0xc6, 0x7b, 0x83, 0x7b, 0x11, 0x28, 0x25,
	0x39, 0x28, 0xed, 0x77, 0x87, 0x5b, 0xa3, 0xfe, 0x64, 0x7c, 0xc5, 0x9b, 0x72, 0x55, 0xad, 0x9b,
	0x35, 0x4b, 0xf3, 0x52, 0xf7, 0x36, 0x63, 0xaf, 0xc5, 0x81, 0xc2, 0x3b, 0x86, 0xa9, 0x4c, 0x98,
	0xf6, 0xa1, 0x1f, 0x10, 0x37, 0x49, 0xa4, 0x99, 0x24, 0xd2, 0x4e, 0x12, 0x89, 0x40, 0x56, 0xd3,
	0x17, 0x4d, 0xed, 0xcf, 0xb3, 0x41, 0xff, 0x98, 0x95, 0xc5, 0xf3, 0xb0, 0x99, 0xae, 0xf0, 0xeb,
	0x8f, 0xc1, 0x28, 0x93, 0x26, 0x3f, 0x4c, 0x48, 0x0a, 0x25, 0x6d, 0xa7, 0xd0, 0x7d, 0xc6, 0x9a,
	0x7f, 0xa2, 0xe6, 0xb8, 0x16, 0xda, 0xd6, 0xeb, 0xb8, 0xed, 0x14, 0xbe, 0xc6, 0xf7, 0xff, 0xa1,
	0xcf, 0xf3, 0xf1, 0x4d, 0x0d, 0x07, 0x26, 0x62, 0xb5, 0x35, 0xb8, 0x1d, 0xaf, 0xc2, 0x26, 0x93,
	0x33, 0xc5, 0x9b, 0x4c, 0xd7, 0x65, 0xda, 0x70, 0x1a, 0x9d, 0x2c, 0x02, 0x74, 0xba, 0x08, 0xd0,
	0xf9, 0x22, 0x40, 0x5f, 0x96, 0x41, 0xe7, 0x74, 0x19, 0x74, 0xbe, 0x2f, 0x83, 0xce, 0x87, 0xa7,
	0x97, 0xa4, 0x25, 0x62, 0xec, 0x2e, 0x8a, 0x5e, 0x6c, 0xd4, 0x91, 0xdd, 0x52, 0xab, 0x30, 0xd9,
	0xb1, 0x7b, 0xf2, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x76, 0x0e, 0xda, 0xc2, 0x03,
	0x00, 0x00,
}

func (m *PartitionedPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartitionedPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartitionedPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccessClientList != nil {
		{
			size, err := m.AccessClientList.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.DesiredState != nil {
		{
			size, err := m.DesiredState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPartitionedPool(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ParitionedPoolConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParitionedPoolConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParitionedPoolConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Corridors) > 0 {
		for iNdEx := len(m.Corridors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Corridors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Threshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ParitionedPoolCorridor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParitionedPoolCorridor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParitionedPoolCorridor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Target.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Threshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPartitionedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ParitionedPoolThreshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParitionedPoolThreshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParitionedPoolThreshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HardCap != 0 {
		i = encodeVarintPartitionedPool(dAtA, i, uint64(m.HardCap))
		i--
		dAtA[i] = 0x10
	}
	if m.SoftCap != 0 {
		i = encodeVarintPartitionedPool(dAtA, i, uint64(m.SoftCap))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPartitionedPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPartitionedPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PartitionedPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPartitionedPool(uint64(l))
	}
	if m.DesiredState != nil {
		l = m.DesiredState.Size()
		n += 1 + l + sovPartitionedPool(uint64(l))
	}
	if m.AccessClientList != nil {
		l = m.AccessClientList.Size()
		n += 1 + l + sovPartitionedPool(uint64(l))
	}
	return n
}

func (m *ParitionedPoolConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Threshold.Size()
	n += 1 + l + sovPartitionedPool(uint64(l))
	if len(m.Corridors) > 0 {
		for _, e := range m.Corridors {
			l = e.Size()
			n += 1 + l + sovPartitionedPool(uint64(l))
		}
	}
	return n
}

func (m *ParitionedPoolCorridor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Threshold.Size()
	n += 1 + l + sovPartitionedPool(uint64(l))
	l = m.Target.Size()
	n += 1 + l + sovPartitionedPool(uint64(l))
	return n
}

func (m *ParitionedPoolThreshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SoftCap != 0 {
		n += 1 + sovPartitionedPool(uint64(m.SoftCap))
	}
	if m.HardCap != 0 {
		n += 1 + sovPartitionedPool(uint64(m.HardCap))
	}
	return n
}

func sovPartitionedPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPartitionedPool(x uint64) (n int) {
	return sovPartitionedPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PartitionedPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPartitionedPool
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
			return fmt.Errorf("proto: PartitionedPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartitionedPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DesiredState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DesiredState == nil {
				m.DesiredState = &ParitionedPoolConfiguration{}
			}
			if err := m.DesiredState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessClientList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessClientList == nil {
				m.AccessClientList = &types.AccessClientList{}
			}
			if err := m.AccessClientList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPartitionedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPartitionedPool
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
func (m *ParitionedPoolConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPartitionedPool
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
			return fmt.Errorf("proto: ParitionedPoolConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParitionedPoolConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Threshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Corridors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Corridors = append(m.Corridors, &ParitionedPoolCorridor{})
			if err := m.Corridors[len(m.Corridors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPartitionedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPartitionedPool
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
func (m *ParitionedPoolCorridor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPartitionedPool
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
			return fmt.Errorf("proto: ParitionedPoolCorridor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParitionedPoolCorridor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Threshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
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
				return ErrInvalidLengthPartitionedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPartitionedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Target.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPartitionedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPartitionedPool
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
func (m *ParitionedPoolThreshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPartitionedPool
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
			return fmt.Errorf("proto: ParitionedPoolThreshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParitionedPoolThreshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftCap", wireType)
			}
			m.SoftCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftCap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardCap", wireType)
			}
			m.HardCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPartitionedPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HardCap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPartitionedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPartitionedPool
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
func skipPartitionedPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPartitionedPool
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
					return 0, ErrIntOverflowPartitionedPool
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
					return 0, ErrIntOverflowPartitionedPool
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
				return 0, ErrInvalidLengthPartitionedPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPartitionedPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPartitionedPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPartitionedPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPartitionedPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPartitionedPool = fmt.Errorf("proto: unexpected end of group")
)