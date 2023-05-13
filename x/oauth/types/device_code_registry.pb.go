// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/oauth/device_code_registry.proto

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

type DeviceCodeRegistry struct {
	Tenant string           `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Codes  []DeviceCodeInfo `protobuf:"bytes,2,rep,name=codes,proto3" json:"codes"`
}

func (m *DeviceCodeRegistry) Reset()         { *m = DeviceCodeRegistry{} }
func (m *DeviceCodeRegistry) String() string { return proto.CompactTextString(m) }
func (*DeviceCodeRegistry) ProtoMessage()    {}
func (*DeviceCodeRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_817a404171c0e891, []int{0}
}
func (m *DeviceCodeRegistry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCodeRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCodeRegistry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCodeRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCodeRegistry.Merge(m, src)
}
func (m *DeviceCodeRegistry) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCodeRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCodeRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCodeRegistry proto.InternalMessageInfo

func (m *DeviceCodeRegistry) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *DeviceCodeRegistry) GetCodes() []DeviceCodeInfo {
	if m != nil {
		return m.Codes
	}
	return nil
}

type DeviceCodeInfo struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	DeviceCode string `protobuf:"bytes,2,opt,name=deviceCode,proto3" json:"deviceCode,omitempty"`
	ExpiresAt  int64  `protobuf:"varint,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
}

func (m *DeviceCodeInfo) Reset()         { *m = DeviceCodeInfo{} }
func (m *DeviceCodeInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceCodeInfo) ProtoMessage()    {}
func (*DeviceCodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_817a404171c0e891, []int{1}
}
func (m *DeviceCodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceCodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceCodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceCodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCodeInfo.Merge(m, src)
}
func (m *DeviceCodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *DeviceCodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCodeInfo proto.InternalMessageInfo

func (m *DeviceCodeInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DeviceCodeInfo) GetDeviceCode() string {
	if m != nil {
		return m.DeviceCode
	}
	return ""
}

func (m *DeviceCodeInfo) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceCodeRegistry)(nil), "beheroes.doxchain.oauth.DeviceCodeRegistry")
	proto.RegisterType((*DeviceCodeInfo)(nil), "beheroes.doxchain.oauth.DeviceCodeInfo")
}

func init() {
	proto.RegisterFile("doxchain/oauth/device_code_registry.proto", fileDescriptor_817a404171c0e891)
}

var fileDescriptor_817a404171c0e891 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0x3d, 0x4e, 0xc3, 0x30,
	0x14, 0x8e, 0x1b, 0x28, 0xaa, 0x91, 0x18, 0x2c, 0x04, 0x11, 0x42, 0x26, 0xea, 0x42, 0x10, 0xc2,
	0x91, 0xe0, 0x04, 0xb4, 0x30, 0xb0, 0x66, 0x64, 0xa9, 0xf2, 0xf3, 0x48, 0x3c, 0x90, 0x17, 0x6c,
	0x17, 0xa5, 0xb7, 0xe0, 0x58, 0x1d, 0x3b, 0x32, 0x21, 0x94, 0x5c, 0x04, 0xd5, 0x49, 0x5b, 0x18,
	0xd8, 0xfc, 0xbe, 0x5f, 0xf9, 0xa3, 0x57, 0x19, 0xd6, 0x69, 0x11, 0xcb, 0x32, 0xc4, 0x78, 0x6e,
	0x8a, 0x30, 0x83, 0x77, 0x99, 0xc2, 0x2c, 0xc5, 0x0c, 0x66, 0x0a, 0x72, 0xa9, 0x8d, 0x5a, 0x88,
	0x4a, 0xa1, 0x41, 0x76, 0x9a, 0x40, 0x01, 0x0a, 0x41, 0x8b, 0x8d, 0x47, 0x58, 0xcf, 0xd9, 0x71,
	0x8e, 0x39, 0x5a, 0x4d, 0xb8, 0x7e, 0x75, 0xf2, 0xf1, 0x1b, 0x65, 0x0f, 0x36, 0x6c, 0x8a, 0x19,
	0x44, 0x7d, 0x14, 0x3b, 0xa1, 0x43, 0x03, 0x65, 0x5c, 0x1a, 0x8f, 0xf8, 0x24, 0x18, 0x45, 0xfd,
	0xc5, 0xa6, 0x74, 0x7f, 0xdd, 0xa9, 0xbd, 0x81, 0xef, 0x06, 0x87, 0xb7, 0x97, 0xe2, 0x9f, 0x32,
	0xb1, 0xcb, 0x7c, 0x2a, 0x5f, 0x70, 0xb2, 0xb7, 0xfc, 0xba, 0x70, 0xa2, 0xce, 0x3b, 0x2e, 0xe8,
	0xd1, 0x5f, 0x9a, 0x79, 0xf4, 0x20, 0x55, 0x10, 0x1b, 0x54, 0x7d, 0xdf, 0xe6, 0x64, 0x9c, 0xd2,
	0x6c, 0xab, 0xf5, 0x06, 0x96, 0xfc, 0x85, 0xb0, 0x73, 0x3a, 0x82, 0xba, 0x92, 0x0a, 0xf4, 0xbd,
	0xf1, 0x5c, 0x9f, 0x04, 0x6e, 0xb4, 0x03, 0x26, 0x8f, 0xcb, 0x86, 0x93, 0x55, 0xc3, 0xc9, 0x77,
	0xc3, 0xc9, 0x47, 0xcb, 0x9d, 0x55, 0xcb, 0x9d, 0xcf, 0x96, 0x3b, 0xcf, 0xd7, 0xb9, 0x34, 0xc5,
	0x3c, 0x11, 0x29, 0xbe, 0x86, 0x09, 0xdc, 0x74, 0x9f, 0x08, 0xb7, 0x2b, 0xd7, 0xfd, 0xce, 0x66,
	0x51, 0x81, 0x4e, 0x86, 0x76, 0xaa, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0xa1, 0x90,
	0xc8, 0x86, 0x01, 0x00, 0x00,
}

func (m *DeviceCodeRegistry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCodeRegistry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCodeRegistry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Codes) > 0 {
		for iNdEx := len(m.Codes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Codes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeviceCodeRegistry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintDeviceCodeRegistry(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeviceCodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceCodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceCodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiresAt != 0 {
		i = encodeVarintDeviceCodeRegistry(dAtA, i, uint64(m.ExpiresAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DeviceCode) > 0 {
		i -= len(m.DeviceCode)
		copy(dAtA[i:], m.DeviceCode)
		i = encodeVarintDeviceCodeRegistry(dAtA, i, uint64(len(m.DeviceCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintDeviceCodeRegistry(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeviceCodeRegistry(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeviceCodeRegistry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceCodeRegistry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovDeviceCodeRegistry(uint64(l))
	}
	if len(m.Codes) > 0 {
		for _, e := range m.Codes {
			l = e.Size()
			n += 1 + l + sovDeviceCodeRegistry(uint64(l))
		}
	}
	return n
}

func (m *DeviceCodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovDeviceCodeRegistry(uint64(l))
	}
	l = len(m.DeviceCode)
	if l > 0 {
		n += 1 + l + sovDeviceCodeRegistry(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovDeviceCodeRegistry(uint64(m.ExpiresAt))
	}
	return n
}

func sovDeviceCodeRegistry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeviceCodeRegistry(x uint64) (n int) {
	return sovDeviceCodeRegistry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceCodeRegistry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceCodeRegistry
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
			return fmt.Errorf("proto: DeviceCodeRegistry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCodeRegistry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodeRegistry
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
				return ErrInvalidLengthDeviceCodeRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Codes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodeRegistry
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
				return ErrInvalidLengthDeviceCodeRegistry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Codes = append(m.Codes, DeviceCodeInfo{})
			if err := m.Codes[len(m.Codes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceCodeRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
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
func (m *DeviceCodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeviceCodeRegistry
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
			return fmt.Errorf("proto: DeviceCodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceCodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodeRegistry
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
				return ErrInvalidLengthDeviceCodeRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodeRegistry
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
				return ErrInvalidLengthDeviceCodeRegistry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeviceCodeRegistry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeviceCodeRegistry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeviceCodeRegistry
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
func skipDeviceCodeRegistry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeviceCodeRegistry
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
					return 0, ErrIntOverflowDeviceCodeRegistry
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
					return 0, ErrIntOverflowDeviceCodeRegistry
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
				return 0, ErrInvalidLengthDeviceCodeRegistry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeviceCodeRegistry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeviceCodeRegistry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeviceCodeRegistry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeviceCodeRegistry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeviceCodeRegistry = fmt.Errorf("proto: unexpected end of group")
)
