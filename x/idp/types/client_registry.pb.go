// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/client_registry.proto

package types

import (
	fmt "fmt"
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

type ClientRegistry struct {
	Creator       string                `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Registrations []*ClientRegistration `protobuf:"bytes,2,rep,name=registrations,proto3" json:"registrations,omitempty"`
}

func (m *ClientRegistry) Reset()         { *m = ClientRegistry{} }
func (m *ClientRegistry) String() string { return proto.CompactTextString(m) }
func (*ClientRegistry) ProtoMessage()    {}
func (*ClientRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_596120f11a70689b, []int{0}
}
func (m *ClientRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRegistry.Merge(m, src)
}
func (m *ClientRegistry) XXX_Size() int {
	return m.Size()
}
func (m *ClientRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRegistry proto.InternalMessageInfo

func (m *ClientRegistry) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ClientRegistry) GetRegistrations() []*ClientRegistration {
	if m != nil {
		return m.Registrations
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientRegistry)(nil), "beheroes.doxchain.idp.ClientRegistry")
}

func init() {
	proto.RegisterFile("doxchain/idp/client_registry.proto", fileDescriptor_596120f11a70689b)
}

var fileDescriptor_596120f11a70689b = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xc9, 0xaf, 0x48,
	0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0x4c, 0x29, 0xd0, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89,
	0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x4d, 0x4a, 0xcd, 0x48, 0x2d, 0xca, 0x4f, 0x2d, 0xd6, 0x83, 0x29, 0xd6, 0xcb, 0x4c, 0x29, 0x90,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0xa5, 0xd4, 0xf0, 0x18,
	0x98, 0x58, 0x92, 0x99, 0x9f, 0x07, 0x51, 0xa7, 0xd4, 0xc8, 0xc8, 0xc5, 0xe7, 0x0c, 0x96, 0x0d,
	0x82, 0xda, 0x26, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x85, 0x72, 0xf1, 0x22, 0x1b, 0x51, 0x2c, 0xc1, 0xa4,
	0xc0, 0xac, 0xc1, 0x6d, 0xa4, 0xa9, 0x87, 0xd5, 0x65, 0x7a, 0x28, 0xe6, 0x82, 0x75, 0x38, 0xb1,
	0x9c, 0xb8, 0x27, 0xcf, 0x18, 0x84, 0x6a, 0x8a, 0x93, 0xf3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37,
	0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x27, 0xa5, 0xea, 0x42, 0x2c, 0xd1, 0x87, 0x7b, 0xad, 0x02, 0xec, 0xb9, 0x92, 0xca, 0x82, 0xd4,
	0xe2, 0x24, 0x36, 0xb0, 0x7f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0xcc, 0x54, 0x28,
	0x4a, 0x01, 0x00, 0x00,
}

func (m *ClientRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Registrations) > 0 {
		for iNdEx := len(m.Registrations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Registrations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClientRegistry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintClientRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClientRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovClientRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovClientRegistry(uint64(l))
	}
	if len(m.Registrations) > 0 {
		for _, e := range m.Registrations {
			l = e.Size()
			n += 1 + l + sovClientRegistry(uint64(l))
		}
	}
	return n
}

func sovClientRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClientRegistry(x uint64) (n int) {
	return sovClientRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientRegistry
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
			return fmt.Errorf("proto: ClientRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistry
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
				return ErrInvalidLengthClientRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Registrations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistry
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
				return ErrInvalidLengthClientRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Registrations = append(m.Registrations, &ClientRegistration{})
			if err := m.Registrations[len(m.Registrations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClientRegistry
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
func skipClientRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientRegistry
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
					return 0, ErrIntOverflowClientRegistry
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
					return 0, ErrIntOverflowClientRegistry
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
				return 0, ErrInvalidLengthClientRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClientRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClientRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClientRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClientRegistry = fmt.Errorf("proto: unexpected end of group")
)