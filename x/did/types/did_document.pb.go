// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/did/did_document.proto

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

type DidDocument struct {
	Context              []string                    `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"`
	Id                   Did                         `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	Controller           *Did                        `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
	AlsoKnownAs          []*Did                      `protobuf:"bytes,4,rep,name=alsoKnownAs,proto3" json:"alsoKnownAs,omitempty"`
	VerificationMethod   []*VerificationMethod       `protobuf:"bytes,5,rep,name=verificationMethod,proto3" json:"verificationMethod,omitempty"`
	Authentication       []*VerificationRelationship `protobuf:"bytes,6,rep,name=authentication,proto3" json:"authentication,omitempty"`
	AssertionMethod      []*VerificationRelationship `protobuf:"bytes,7,rep,name=assertionMethod,proto3" json:"assertionMethod,omitempty"`
	KeyAgreement         []*VerificationRelationship `protobuf:"bytes,8,rep,name=keyAgreement,proto3" json:"keyAgreement,omitempty"`
	CapabilityInvocation []*VerificationRelationship `protobuf:"bytes,9,rep,name=capabilityInvocation,proto3" json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []*VerificationRelationship `protobuf:"bytes,10,rep,name=capabilityDelegation,proto3" json:"capabilityDelegation,omitempty"`
	Service              []*Service                  `protobuf:"bytes,11,rep,name=service,proto3" json:"service,omitempty"`
}

func (m *DidDocument) Reset()         { *m = DidDocument{} }
func (m *DidDocument) String() string { return proto.CompactTextString(m) }
func (*DidDocument) ProtoMessage()    {}
func (*DidDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d66eb31aa545c75, []int{0}
}
func (m *DidDocument) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidDocument.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidDocument.Merge(m, src)
}
func (m *DidDocument) XXX_Size() int {
	return m.Size()
}
func (m *DidDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_DidDocument.DiscardUnknown(m)
}

var xxx_messageInfo_DidDocument proto.InternalMessageInfo

func (m *DidDocument) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *DidDocument) GetId() Did {
	if m != nil {
		return m.Id
	}
	return Did{}
}

func (m *DidDocument) GetController() *Did {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *DidDocument) GetAlsoKnownAs() []*Did {
	if m != nil {
		return m.AlsoKnownAs
	}
	return nil
}

func (m *DidDocument) GetVerificationMethod() []*VerificationMethod {
	if m != nil {
		return m.VerificationMethod
	}
	return nil
}

func (m *DidDocument) GetAuthentication() []*VerificationRelationship {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *DidDocument) GetAssertionMethod() []*VerificationRelationship {
	if m != nil {
		return m.AssertionMethod
	}
	return nil
}

func (m *DidDocument) GetKeyAgreement() []*VerificationRelationship {
	if m != nil {
		return m.KeyAgreement
	}
	return nil
}

func (m *DidDocument) GetCapabilityInvocation() []*VerificationRelationship {
	if m != nil {
		return m.CapabilityInvocation
	}
	return nil
}

func (m *DidDocument) GetCapabilityDelegation() []*VerificationRelationship {
	if m != nil {
		return m.CapabilityDelegation
	}
	return nil
}

func (m *DidDocument) GetService() []*Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*DidDocument)(nil), "beheroes.doxchain.did.DidDocument")
}

func init() { proto.RegisterFile("doxchain/did/did_document.proto", fileDescriptor_7d66eb31aa545c75) }

var fileDescriptor_7d66eb31aa545c75 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x76, 0x6d, 0xdd, 0xa9, 0x28, 0x0c, 0xab, 0x0c, 0x3d, 0x64, 0x8b, 0x07, 0xd9,
	0x3d, 0x98, 0x88, 0xde, 0xc5, 0xad, 0xbd, 0x88, 0x78, 0xa9, 0x20, 0x28, 0x48, 0x49, 0x32, 0xcf,
	0xe4, 0x61, 0x3a, 0x53, 0x66, 0xa6, 0xb5, 0xfd, 0x16, 0x7e, 0xac, 0x3d, 0xee, 0xd1, 0x93, 0x48,
	0xfb, 0x09, 0xfc, 0x06, 0xd2, 0xc9, 0x44, 0x93, 0xda, 0x42, 0xa1, 0x7b, 0x08, 0x49, 0xde, 0xfc,
	0xff, 0xbf, 0xff, 0xe3, 0x31, 0x8f, 0x9c, 0x73, 0xb9, 0x48, 0xb2, 0x08, 0x45, 0xc8, 0x91, 0x6f,
	0x9e, 0x31, 0x97, 0xc9, 0x6c, 0x02, 0xc2, 0x04, 0x53, 0x25, 0x8d, 0xa4, 0x0f, 0x63, 0xc8, 0x40,
	0x49, 0xd0, 0x41, 0xa9, 0x0c, 0x38, 0xf2, 0xde, 0x59, 0x2a, 0x53, 0x69, 0x15, 0xe1, 0xe6, 0xab,
	0x10, 0xf7, 0x1e, 0x6d, 0xd3, 0x5c, 0xfd, 0x49, 0xad, 0x3e, 0x07, 0x85, 0x5f, 0x30, 0x89, 0x0c,
	0x4a, 0x31, 0x9e, 0x80, 0xc9, 0x64, 0xa9, 0xeb, 0xd5, 0x74, 0x1a, 0xd4, 0x1c, 0x13, 0x28, 0xce,
	0x1e, 0xff, 0x6e, 0x93, 0xee, 0x10, 0xf9, 0xd0, 0xb5, 0x47, 0x19, 0xe9, 0x24, 0x52, 0x18, 0x58,
	0x18, 0xe6, 0xf5, 0x5b, 0x17, 0xa7, 0xa3, 0xf2, 0x97, 0x3e, 0x23, 0x4d, 0xe4, 0xac, 0xd9, 0xf7,
	0x2e, 0xba, 0xcf, 0x7b, 0xc1, 0xce, 0xfe, 0x83, 0x21, 0xf2, 0xc1, 0xc9, 0xf5, 0xcf, 0xf3, 0xc6,
	0xa8, 0x89, 0x9c, 0xbe, 0x22, 0x64, 0x63, 0x56, 0x32, 0xcf, 0x41, 0xb1, 0xd6, 0x41, 0x4e, 0x6f,
	0x54, 0xf1, 0xd0, 0x01, 0xe9, 0x46, 0xb9, 0x96, 0x6f, 0x85, 0xfc, 0x26, 0xae, 0x34, 0x3b, 0xe9,
	0xb7, 0x0e, 0x42, 0x54, 0x4d, 0x74, 0x4c, 0x68, 0x75, 0x34, 0xef, 0xec, 0x64, 0xd8, 0x1d, 0x8b,
	0xba, 0xdc, 0x83, 0xfa, 0xf0, 0x9f, 0xc1, 0x91, 0x77, 0xa0, 0xe8, 0x67, 0x72, 0x3f, 0x9a, 0x99,
	0x0c, 0x84, 0x71, 0x75, 0xd6, 0xb6, 0xf0, 0xf0, 0x00, 0xf8, 0x08, 0x72, 0xfb, 0xd6, 0x19, 0x4e,
	0x5d, 0xc4, 0x16, 0x8c, 0x8e, 0xc9, 0x83, 0x48, 0x6b, 0x50, 0x95, 0xe6, 0x3b, 0xc7, 0xf0, 0xb7,
	0x69, 0xf4, 0x23, 0xb9, 0xf7, 0x15, 0x96, 0x57, 0xa9, 0x02, 0xd8, 0x5c, 0x01, 0x76, 0xf7, 0x18,
	0x7a, 0x0d, 0x45, 0x91, 0x9c, 0x25, 0xd1, 0x34, 0x8a, 0x31, 0x47, 0xb3, 0x7c, 0x23, 0xe6, 0xd2,
	0x0d, 0xe8, 0xf4, 0x98, 0x88, 0x9d, 0xc8, 0x7a, 0xd4, 0x10, 0x72, 0x48, 0x8b, 0x28, 0x72, 0x4b,
	0x51, 0xff, 0x90, 0xf4, 0x25, 0xe9, 0xb8, 0x25, 0x62, 0x5d, 0x4b, 0xf7, 0xf7, 0xd0, 0xdf, 0x17,
	0x2a, 0x07, 0x2b, 0x4d, 0x83, 0xd7, 0xd7, 0x2b, 0xdf, 0xbb, 0x59, 0xf9, 0xde, 0xaf, 0x95, 0xef,
	0x7d, 0x5f, 0xfb, 0x8d, 0x9b, 0xb5, 0xdf, 0xf8, 0xb1, 0xf6, 0x1b, 0x9f, 0x2e, 0x53, 0x34, 0xd9,
	0x2c, 0x0e, 0x12, 0x39, 0x09, 0x63, 0x78, 0x5a, 0x30, 0xc3, 0xbf, 0xeb, 0xbb, 0xb0, 0x0b, 0x6c,
	0x96, 0x53, 0xd0, 0x71, 0xdb, 0xee, 0xef, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x01,
	0xbe, 0x8a, 0x6b, 0x04, 0x00, 0x00,
}

func (m *DidDocument) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidDocument) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidDocument) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Service) > 0 {
		for iNdEx := len(m.Service) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Service[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for iNdEx := len(m.CapabilityDelegation) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CapabilityDelegation[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for iNdEx := len(m.CapabilityInvocation) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CapabilityInvocation[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.KeyAgreement) > 0 {
		for iNdEx := len(m.KeyAgreement) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyAgreement[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.AssertionMethod) > 0 {
		for iNdEx := len(m.AssertionMethod) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssertionMethod[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authentication[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.VerificationMethod) > 0 {
		for iNdEx := len(m.VerificationMethod) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerificationMethod[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.AlsoKnownAs) > 0 {
		for iNdEx := len(m.AlsoKnownAs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AlsoKnownAs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDidDocument(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Controller != nil {
		{
			size, err := m.Controller.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDidDocument(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDidDocument(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintDidDocument(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDidDocument(dAtA []byte, offset int, v uint64) int {
	offset -= sovDidDocument(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidDocument) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	l = m.Id.Size()
	n += 1 + l + sovDidDocument(uint64(l))
	if m.Controller != nil {
		l = m.Controller.Size()
		n += 1 + l + sovDidDocument(uint64(l))
	}
	if len(m.AlsoKnownAs) > 0 {
		for _, e := range m.AlsoKnownAs {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.VerificationMethod) > 0 {
		for _, e := range m.VerificationMethod {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.Authentication) > 0 {
		for _, e := range m.Authentication {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.AssertionMethod) > 0 {
		for _, e := range m.AssertionMethod {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.KeyAgreement) > 0 {
		for _, e := range m.KeyAgreement {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.CapabilityInvocation) > 0 {
		for _, e := range m.CapabilityInvocation {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.CapabilityDelegation) > 0 {
		for _, e := range m.CapabilityDelegation {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	if len(m.Service) > 0 {
		for _, e := range m.Service {
			l = e.Size()
			n += 1 + l + sovDidDocument(uint64(l))
		}
	}
	return n
}

func sovDidDocument(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDidDocument(x uint64) (n int) {
	return sovDidDocument(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidDocument) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDidDocument
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
			return fmt.Errorf("proto: DidDocument: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidDocument: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Controller == nil {
				m.Controller = &Did{}
			}
			if err := m.Controller.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlsoKnownAs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AlsoKnownAs = append(m.AlsoKnownAs, &Did{})
			if err := m.AlsoKnownAs[len(m.AlsoKnownAs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethod = append(m.VerificationMethod, &VerificationMethod{})
			if err := m.VerificationMethod[len(m.VerificationMethod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authentication = append(m.Authentication, &VerificationRelationship{})
			if err := m.Authentication[len(m.Authentication)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionMethod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionMethod = append(m.AssertionMethod, &VerificationRelationship{})
			if err := m.AssertionMethod[len(m.AssertionMethod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyAgreement", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyAgreement = append(m.KeyAgreement, &VerificationRelationship{})
			if err := m.KeyAgreement[len(m.KeyAgreement)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityInvocation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityInvocation = append(m.CapabilityInvocation, &VerificationRelationship{})
			if err := m.CapabilityInvocation[len(m.CapabilityInvocation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapabilityDelegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CapabilityDelegation = append(m.CapabilityDelegation, &VerificationRelationship{})
			if err := m.CapabilityDelegation[len(m.CapabilityDelegation)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDidDocument
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
				return ErrInvalidLengthDidDocument
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDidDocument
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Service = append(m.Service, &Service{})
			if err := m.Service[len(m.Service)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDidDocument(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDidDocument
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
func skipDidDocument(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDidDocument
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
					return 0, ErrIntOverflowDidDocument
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
					return 0, ErrIntOverflowDidDocument
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
				return 0, ErrInvalidLengthDidDocument
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDidDocument
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDidDocument
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDidDocument        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDidDocument          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDidDocument = fmt.Errorf("proto: unexpected end of group")
)
