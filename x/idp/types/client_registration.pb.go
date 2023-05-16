// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/client_registration.proto

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

type ClientRegistration struct {
	Id                         types.Did         `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	AppId                      types.Did         `protobuf:"bytes,2,opt,name=appId,proto3" json:"appId"`
	Name                       string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	AppRoles                   []string          `protobuf:"bytes,4,rep,name=appRoles,proto3" json:"appRoles,omitempty"`
	AccessTokenAcceptedVersion uint32            `protobuf:"varint,5,opt,name=accessTokenAcceptedVersion,proto3" json:"accessTokenAcceptedVersion,omitempty"`
	AllowPublicClient          uint32            `protobuf:"varint,6,opt,name=allowPublicClient,proto3" json:"allowPublicClient,omitempty"`
	ReplyUrls                  []string          `protobuf:"bytes,7,rep,name=replyUrls,proto3" json:"replyUrls,omitempty"`
	Tags                       []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	AccessClientList           *AccessClientList `protobuf:"bytes,9,opt,name=accessClientList,proto3" json:"accessClientList,omitempty"`
}

func (m *ClientRegistration) Reset()         { *m = ClientRegistration{} }
func (m *ClientRegistration) String() string { return proto.CompactTextString(m) }
func (*ClientRegistration) ProtoMessage()    {}
func (*ClientRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b44e633ad8377f2d, []int{0}
}
func (m *ClientRegistration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientRegistration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientRegistration.Merge(m, src)
}
func (m *ClientRegistration) XXX_Size() int {
	return m.Size()
}
func (m *ClientRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_ClientRegistration proto.InternalMessageInfo

func (m *ClientRegistration) GetId() types.Did {
	if m != nil {
		return m.Id
	}
	return types.Did{}
}

func (m *ClientRegistration) GetAppId() types.Did {
	if m != nil {
		return m.AppId
	}
	return types.Did{}
}

func (m *ClientRegistration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClientRegistration) GetAppRoles() []string {
	if m != nil {
		return m.AppRoles
	}
	return nil
}

func (m *ClientRegistration) GetAccessTokenAcceptedVersion() uint32 {
	if m != nil {
		return m.AccessTokenAcceptedVersion
	}
	return 0
}

func (m *ClientRegistration) GetAllowPublicClient() uint32 {
	if m != nil {
		return m.AllowPublicClient
	}
	return 0
}

func (m *ClientRegistration) GetReplyUrls() []string {
	if m != nil {
		return m.ReplyUrls
	}
	return nil
}

func (m *ClientRegistration) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ClientRegistration) GetAccessClientList() *AccessClientList {
	if m != nil {
		return m.AccessClientList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientRegistration)(nil), "beheroes.doxchain.idp.ClientRegistration")
}

func init() {
	proto.RegisterFile("doxchain/idp/client_registration.proto", fileDescriptor_b44e633ad8377f2d)
}

var fileDescriptor_b44e633ad8377f2d = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0xaa, 0x13, 0x31,
	0x14, 0x9e, 0xf4, 0xcf, 0x4e, 0x44, 0xd0, 0xa0, 0x12, 0x06, 0x19, 0x07, 0x41, 0x1d, 0x41, 0x67,
	0x44, 0xc1, 0xa5, 0xd0, 0xd6, 0x8d, 0xe0, 0x42, 0x06, 0x15, 0x74, 0x53, 0x32, 0x49, 0x98, 0x06,
	0xd3, 0x49, 0x98, 0xa4, 0xd8, 0xbe, 0x85, 0xaf, 0xe1, 0x9b, 0x74, 0xd9, 0xa5, 0x2b, 0x91, 0xf6,
	0x45, 0x2e, 0x93, 0xf4, 0xf6, 0xfe, 0xf4, 0x72, 0xb9, 0x8b, 0xc0, 0xc9, 0x39, 0xdf, 0x39, 0xdf,
	0x77, 0x7e, 0xe0, 0x33, 0xa6, 0x96, 0x74, 0x46, 0x44, 0x9d, 0x0b, 0xa6, 0x73, 0x2a, 0x05, 0xaf,
	0xed, 0xb4, 0xe1, 0x95, 0x30, 0xb6, 0x21, 0x56, 0xa8, 0x3a, 0xd3, 0x8d, 0xb2, 0x0a, 0x3d, 0x28,
	0xf9, 0x8c, 0x37, 0x8a, 0x9b, 0xec, 0x34, 0x21, 0x13, 0x4c, 0x47, 0xf7, 0x2b, 0x55, 0x29, 0x87,
	0xc8, 0x5b, 0xcb, 0x83, 0xa3, 0x87, 0x87, 0xa2, 0x4c, 0xb0, 0xf6, 0xed, 0xfd, 0x4f, 0x2f, 0x90,
	0x11, 0x4a, 0xb9, 0x31, 0xd3, 0x3d, 0xa7, 0x14, 0xc6, 0x7a, 0xd8, 0x93, 0x3f, 0x5d, 0x88, 0x26,
	0xce, 0x5b, 0x9c, 0x13, 0x82, 0x5e, 0xc3, 0x8e, 0x60, 0x18, 0x24, 0x20, 0xbd, 0xfd, 0x26, 0xca,
	0x8e, 0xf5, 0xb4, 0x3c, 0x1f, 0x04, 0x1b, 0xf7, 0xd6, 0xff, 0x1e, 0x07, 0x45, 0x47, 0x30, 0xf4,
	0x0e, 0xf6, 0x89, 0xd6, 0x1f, 0x19, 0xee, 0xdc, 0x30, 0xc9, 0xc3, 0x11, 0x82, 0xbd, 0x9a, 0xcc,
	0x39, 0xee, 0x26, 0x20, 0x0d, 0x0b, 0x67, 0xa3, 0x08, 0x0e, 0x89, 0xd6, 0x85, 0x92, 0xdc, 0xe0,
	0x5e, 0xd2, 0x4d, 0xc3, 0xe2, 0xf0, 0x47, 0xef, 0x61, 0xe4, 0x9b, 0xf9, 0xa2, 0x7e, 0xf2, 0x7a,
	0x44, 0x29, 0xd7, 0x96, 0xb3, 0x6f, 0xbc, 0x31, 0x42, 0xd5, 0xb8, 0x9f, 0x80, 0xf4, 0x4e, 0x71,
	0x0d, 0x02, 0xbd, 0x84, 0xf7, 0x88, 0x94, 0xea, 0xd7, 0xe7, 0x45, 0x29, 0x05, 0xf5, 0xad, 0xe3,
	0x81, 0x4b, 0x3b, 0x0e, 0xa0, 0x47, 0x30, 0x6c, 0xb8, 0x96, 0xab, 0xaf, 0x8d, 0x34, 0xf8, 0x96,
	0x93, 0x72, 0xe6, 0x68, 0xb5, 0x5b, 0x52, 0x19, 0x3c, 0x74, 0x01, 0x67, 0xa3, 0xef, 0xf0, 0xae,
	0x67, 0xf7, 0x15, 0x3e, 0x09, 0x63, 0x71, 0xe8, 0x46, 0xf2, 0x3c, 0xbb, 0x72, 0xaf, 0xd9, 0xe8,
	0x12, 0xdc, 0xcd, 0x07, 0x14, 0x47, 0x65, 0xc6, 0x93, 0xf5, 0x36, 0x06, 0x9b, 0x6d, 0x0c, 0xfe,
	0x6f, 0x63, 0xf0, 0x7b, 0x17, 0x07, 0x9b, 0x5d, 0x1c, 0xfc, 0xdd, 0xc5, 0xc1, 0x8f, 0x17, 0x95,
	0xb0, 0xb3, 0x45, 0x99, 0x51, 0x35, 0xcf, 0x4b, 0xfe, 0xca, 0xb3, 0xe4, 0x87, 0x0b, 0x58, 0xba,
	0x1b, 0xb0, 0x2b, 0xcd, 0x4d, 0x39, 0x70, 0x7b, 0x7f, 0x7b, 0x12, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x45, 0x27, 0x48, 0x8d, 0x02, 0x00, 0x00,
}

func (m *ClientRegistration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientRegistration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientRegistration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintClientRegistration(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintClientRegistration(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ReplyUrls) > 0 {
		for iNdEx := len(m.ReplyUrls) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReplyUrls[iNdEx])
			copy(dAtA[i:], m.ReplyUrls[iNdEx])
			i = encodeVarintClientRegistration(dAtA, i, uint64(len(m.ReplyUrls[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.AllowPublicClient != 0 {
		i = encodeVarintClientRegistration(dAtA, i, uint64(m.AllowPublicClient))
		i--
		dAtA[i] = 0x30
	}
	if m.AccessTokenAcceptedVersion != 0 {
		i = encodeVarintClientRegistration(dAtA, i, uint64(m.AccessTokenAcceptedVersion))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AppRoles) > 0 {
		for iNdEx := len(m.AppRoles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AppRoles[iNdEx])
			copy(dAtA[i:], m.AppRoles[iNdEx])
			i = encodeVarintClientRegistration(dAtA, i, uint64(len(m.AppRoles[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintClientRegistration(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.AppId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClientRegistration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClientRegistration(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClientRegistration(dAtA []byte, offset int, v uint64) int {
	offset -= sovClientRegistration(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientRegistration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Id.Size()
	n += 1 + l + sovClientRegistration(uint64(l))
	l = m.AppId.Size()
	n += 1 + l + sovClientRegistration(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClientRegistration(uint64(l))
	}
	if len(m.AppRoles) > 0 {
		for _, s := range m.AppRoles {
			l = len(s)
			n += 1 + l + sovClientRegistration(uint64(l))
		}
	}
	if m.AccessTokenAcceptedVersion != 0 {
		n += 1 + sovClientRegistration(uint64(m.AccessTokenAcceptedVersion))
	}
	if m.AllowPublicClient != 0 {
		n += 1 + sovClientRegistration(uint64(m.AllowPublicClient))
	}
	if len(m.ReplyUrls) > 0 {
		for _, s := range m.ReplyUrls {
			l = len(s)
			n += 1 + l + sovClientRegistration(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovClientRegistration(uint64(l))
		}
	}
	if m.AccessClientList != nil {
		l = m.AccessClientList.Size()
		n += 1 + l + sovClientRegistration(uint64(l))
	}
	return n
}

func sovClientRegistration(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClientRegistration(x uint64) (n int) {
	return sovClientRegistration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientRegistration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientRegistration
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
			return fmt.Errorf("proto: ClientRegistration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientRegistration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AppId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppRoles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppRoles = append(m.AppRoles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessTokenAcceptedVersion", wireType)
			}
			m.AccessTokenAcceptedVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessTokenAcceptedVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowPublicClient", wireType)
			}
			m.AllowPublicClient = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AllowPublicClient |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyUrls", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplyUrls = append(m.ReplyUrls, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessClientList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientRegistration
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
				return ErrInvalidLengthClientRegistration
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClientRegistration
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessClientList == nil {
				m.AccessClientList = &AccessClientList{}
			}
			if err := m.AccessClientList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientRegistration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClientRegistration
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
func skipClientRegistration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientRegistration
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
					return 0, ErrIntOverflowClientRegistration
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
					return 0, ErrIntOverflowClientRegistration
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
				return 0, ErrInvalidLengthClientRegistration
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClientRegistration
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClientRegistration
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClientRegistration        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientRegistration          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClientRegistration = fmt.Errorf("proto: unexpected end of group")
)
