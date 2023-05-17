// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/v1beta1/client_registration_relationship_registry_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/be-heroes/doxchain/x/did/types"
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

type ClientRegistrationRelationshipRegistryEntry struct {
	OwnerId          types.Did        `protobuf:"bytes,1,opt,name=ownerId,proto3" json:"ownerId"`
	DestinationId    types.Did        `protobuf:"bytes,2,opt,name=destinationId,proto3" json:"destinationId"`
	AccessClientList AccessClientList `protobuf:"bytes,3,opt,name=accessClientList,proto3" json:"accessClientList"`
}

func (m *ClientRegistrationRelationshipRegistryEntry) Reset() {
	*m = ClientRegistrationRelationshipRegistryEntry{}
}
func (m *ClientRegistrationRelationshipRegistryEntry) String() string {
	return proto.CompactTextString(m)
}
func (*ClientRegistrationRelationshipRegistryEntry) ProtoMessage() {}
func (*ClientRegistrationRelationshipRegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f6a09daefd6f447, []int{0}
}
func (m *ClientRegistrationRelationshipRegistryEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientRegistrationRelationshipRegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientRegistrationRelationshipRegistryEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientRegistrationRelationshipRegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRegistrationRelationshipRegistryEntry.Merge(m, src)
}
func (m *ClientRegistrationRelationshipRegistryEntry) XXX_Size() int {
	return m.Size()
}
func (m *ClientRegistrationRelationshipRegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRegistrationRelationshipRegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRegistrationRelationshipRegistryEntry proto.InternalMessageInfo

func (m *ClientRegistrationRelationshipRegistryEntry) GetOwnerId() types.Did {
	if m != nil {
		return m.OwnerId
	}
	return types.Did{}
}

func (m *ClientRegistrationRelationshipRegistryEntry) GetDestinationId() types.Did {
	if m != nil {
		return m.DestinationId
	}
	return types.Did{}
}

func (m *ClientRegistrationRelationshipRegistryEntry) GetAccessClientList() AccessClientList {
	if m != nil {
		return m.AccessClientList
	}
	return AccessClientList{}
}

func init() {
	proto.RegisterType((*ClientRegistrationRelationshipRegistryEntry)(nil), "beheroes.doxchain.idp.v1beta1.ClientRegistrationRelationshipRegistryEntry")
}

func init() {
	proto.RegisterFile("doxchain/idp/v1beta1/client_registration_relationship_registry_entry.proto", fileDescriptor_7f6a09daefd6f447)
}

var fileDescriptor_7f6a09daefd6f447 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0x2a, 0x0a, 0x11, 0x41, 0x82, 0x87, 0x52, 0x70, 0x95, 0x9e, 0x14, 0xe9, 0x2e,
	0xd5, 0x27, 0xb0, 0xd5, 0x43, 0x45, 0x3c, 0xf4, 0xe8, 0x25, 0x24, 0xd9, 0x21, 0x19, 0xa8, 0xbb,
	0x61, 0x77, 0xd5, 0xe6, 0x2d, 0x7c, 0x08, 0x1f, 0xa6, 0xc7, 0x1e, 0x3d, 0x89, 0xb4, 0x2f, 0x22,
	0xdd, 0x6c, 0x42, 0xb5, 0x22, 0x78, 0x1b, 0x66, 0xe6, 0xff, 0xe7, 0x9b, 0x99, 0xe0, 0x96, 0xcb,
	0x69, 0x9a, 0xc7, 0x28, 0x18, 0xf2, 0x82, 0x3d, 0xf7, 0x13, 0x30, 0x71, 0x9f, 0xa5, 0x13, 0x04,
	0x61, 0x22, 0x05, 0x19, 0x6a, 0xa3, 0x62, 0x83, 0x52, 0x44, 0x0a, 0x26, 0x36, 0xd0, 0x39, 0x16,
	0x75, 0xa5, 0x8c, 0x40, 0x18, 0x55, 0xd2, 0x42, 0x49, 0x23, 0xc3, 0xa3, 0x04, 0x72, 0x50, 0x12,
	0x34, 0xad, 0x4d, 0x29, 0xf2, 0x82, 0x3a, 0xd3, 0xce, 0x61, 0x26, 0x33, 0x69, 0x3b, 0xd9, 0x2a,
	0xaa, 0x44, 0x1d, 0xd2, 0x00, 0x70, 0xe4, 0x0d, 0x00, 0x47, 0xee, 0xea, 0xbd, 0x5f, 0x01, 0xe3,
	0x34, 0x05, 0xad, 0x23, 0xc7, 0x39, 0x41, 0x6d, 0xaa, 0xf6, 0xee, 0x5b, 0x2b, 0x38, 0x1f, 0xda,
	0xec, 0x78, 0x0d, 0x7e, 0xbc, 0xc6, 0xee, 0xf2, 0xe5, 0xcd, 0x8a, 0x3c, 0x1c, 0x04, 0xbb, 0xf2,
	0x45, 0x80, 0x1a, 0xf1, 0xb6, 0x7f, 0xe2, 0x9f, 0xee, 0x5d, 0x74, 0xe9, 0xe6, 0x16, 0x2b, 0x1a,
	0x37, 0x99, 0x5e, 0x23, 0x1f, 0x6c, 0xcf, 0x3e, 0x8e, 0xbd, 0x71, 0x2d, 0x0c, 0xef, 0x83, 0x7d,
	0x0e, 0xda, 0xa0, 0xb0, 0x33, 0x46, 0xbc, 0xdd, 0xfa, 0xa7, 0xd3, 0x77, 0x79, 0x18, 0x07, 0x07,
	0xd5, 0x7e, 0xd5, 0x22, 0x77, 0xa8, 0x4d, 0x7b, 0xcb, 0x5a, 0x32, 0xfa, 0xe7, 0x89, 0xe9, 0xd5,
	0x0f, 0x99, 0xf3, 0xdf, 0xb0, 0x1b, 0x0c, 0x67, 0x0b, 0xe2, 0xcf, 0x17, 0xc4, 0xff, 0x5c, 0x10,
	0xff, 0x75, 0x49, 0xbc, 0xf9, 0x92, 0x78, 0xef, 0x4b, 0xe2, 0x3d, 0x9c, 0x65, 0x68, 0xf2, 0xa7,
	0x84, 0xa6, 0xf2, 0x91, 0x25, 0xd0, 0xab, 0xa6, 0xb1, 0xe6, 0x09, 0x53, 0xfb, 0x06, 0x53, 0x16,
	0xa0, 0x93, 0x1d, 0x7b, 0xf2, 0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xd7, 0xed, 0xf6,
	0x44, 0x02, 0x00, 0x00,
}

func (m *ClientRegistrationRelationshipRegistryEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRegistrationRelationshipRegistryEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientRegistrationRelationshipRegistryEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AccessClientList.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClientRegistrationRelationshipRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DestinationId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClientRegistrationRelationshipRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.OwnerId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClientRegistrationRelationshipRegistryEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClientRegistrationRelationshipRegistryEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovClientRegistrationRelationshipRegistryEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientRegistrationRelationshipRegistryEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OwnerId.Size()
	n += 1 + l + sovClientRegistrationRelationshipRegistryEntry(uint64(l))
	l = m.DestinationId.Size()
	n += 1 + l + sovClientRegistrationRelationshipRegistryEntry(uint64(l))
	l = m.AccessClientList.Size()
	n += 1 + l + sovClientRegistrationRelationshipRegistryEntry(uint64(l))
	return n
}

func sovClientRegistrationRelationshipRegistryEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClientRegistrationRelationshipRegistryEntry(x uint64) (n int) {
	return sovClientRegistrationRelationshipRegistryEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientRegistrationRelationshipRegistryEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
			return fmt.Errorf("proto: ClientRegistrationRelationshipRegistryEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRegistrationRelationshipRegistryEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DestinationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccessClientList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientRegistrationRelationshipRegistryEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
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
func skipClientRegistrationRelationshipRegistryEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
					return 0, ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
					return 0, ErrIntOverflowClientRegistrationRelationshipRegistryEntry
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
				return 0, ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClientRegistrationRelationshipRegistryEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClientRegistrationRelationshipRegistryEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClientRegistrationRelationshipRegistryEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientRegistrationRelationshipRegistryEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClientRegistrationRelationshipRegistryEntry = fmt.Errorf("proto: unexpected end of group")
)
