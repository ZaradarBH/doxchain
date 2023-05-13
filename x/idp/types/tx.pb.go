// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: doxchain/idp/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgAuthenticationRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Tenant  string `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (m *MsgAuthenticationRequest) Reset()         { *m = MsgAuthenticationRequest{} }
func (m *MsgAuthenticationRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticationRequest) ProtoMessage()    {}
func (*MsgAuthenticationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d303193384f0350d, []int{0}
}
func (m *MsgAuthenticationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticationRequest.Merge(m, src)
}
func (m *MsgAuthenticationRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticationRequest proto.InternalMessageInfo

func (m *MsgAuthenticationRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAuthenticationRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

type MsgAuthenticationResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *MsgAuthenticationResponse) Reset()         { *m = MsgAuthenticationResponse{} }
func (m *MsgAuthenticationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAuthenticationResponse) ProtoMessage()    {}
func (*MsgAuthenticationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d303193384f0350d, []int{1}
}
func (m *MsgAuthenticationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAuthenticationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAuthenticationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAuthenticationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAuthenticationResponse.Merge(m, src)
}
func (m *MsgAuthenticationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAuthenticationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAuthenticationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAuthenticationResponse proto.InternalMessageInfo

func (m *MsgAuthenticationResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgAuthenticationRequest)(nil), "beheroes.doxchain.idp.MsgAuthenticationRequest")
	proto.RegisterType((*MsgAuthenticationResponse)(nil), "beheroes.doxchain.idp.MsgAuthenticationResponse")
}

func init() { proto.RegisterFile("doxchain/idp/tx.proto", fileDescriptor_d303193384f0350d) }

var fileDescriptor_d303193384f0350d = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x63, 0x50, 0x8b, 0xf0, 0x68, 0x51, 0x14, 0x18, 0x2c, 0xd4, 0x09, 0x06, 0x6c, 0xfe,
	0x3c, 0x01, 0xb0, 0xb6, 0x4b, 0x47, 0xb6, 0xfc, 0x39, 0x25, 0x06, 0x71, 0xe7, 0xc6, 0x17, 0x29,
	0xbc, 0x05, 0x8f, 0xc5, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0x54, 0xb7, 0x4c, 0x41, 0x62, 0xfc,
	0xe9, 0x74, 0x9f, 0x3e, 0x7d, 0x72, 0x56, 0x52, 0x57, 0xd4, 0x99, 0x43, 0xeb, 0x4a, 0x6f, 0xb9,
	0x33, 0xbe, 0x21, 0x26, 0x35, 0xcb, 0xa1, 0x86, 0x86, 0x20, 0x98, 0xfd, 0xdd, 0xb8, 0xd2, 0xcf,
	0x17, 0x32, 0x5d, 0x86, 0xea, 0xa1, 0xe5, 0x1a, 0x90, 0x5d, 0x91, 0xb1, 0x23, 0x5c, 0xc1, 0xba,
	0x85, 0xc0, 0x2a, 0x95, 0x47, 0x45, 0x03, 0x19, 0x53, 0x93, 0x8a, 0x0b, 0x71, 0x79, 0xbc, 0xda,
	0x4f, 0x75, 0x2a, 0xa7, 0x0c, 0x98, 0x21, 0xa7, 0x07, 0xdb, 0xc3, 0x6e, 0xcd, 0x6f, 0xe5, 0xd9,
	0x08, 0x2d, 0x78, 0xc2, 0x00, 0xea, 0x44, 0x4e, 0x98, 0x5e, 0x01, 0x77, 0xb0, 0x38, 0xee, 0xd6,
	0xf2, 0x70, 0x19, 0x2a, 0xf5, 0x22, 0x27, 0x0b, 0xaa, 0x1c, 0x2a, 0x6b, 0x46, 0x45, 0xcd, 0x5f,
	0x96, 0xe7, 0x37, 0xff, 0x7f, 0x88, 0x22, 0x8f, 0x4f, 0x9f, 0xbd, 0x16, 0x9b, 0x5e, 0x8b, 0xef,
	0x5e, 0x8b, 0x8f, 0x41, 0x27, 0x9b, 0x41, 0x27, 0x5f, 0x83, 0x4e, 0x9e, 0xaf, 0x2a, 0xc7, 0x75,
	0x9b, 0x9b, 0x82, 0xde, 0x6c, 0x0e, 0xd7, 0x11, 0x6b, 0x7f, 0x83, 0x76, 0x31, 0xe9, 0xbb, 0x87,
	0x90, 0x4f, 0xb7, 0x59, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xc7, 0x29, 0xb4, 0x6f,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Login(ctx context.Context, in *MsgAuthenticationRequest, opts ...grpc.CallOption) (*MsgAuthenticationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Login(ctx context.Context, in *MsgAuthenticationRequest, opts ...grpc.CallOption) (*MsgAuthenticationResponse, error) {
	out := new(MsgAuthenticationResponse)
	err := c.cc.Invoke(ctx, "/beheroes.doxchain.idp.Msg/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Login(context.Context, *MsgAuthenticationRequest) (*MsgAuthenticationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Login(ctx context.Context, req *MsgAuthenticationRequest) (*MsgAuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beheroes.doxchain.idp.Msg/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Login(ctx, req.(*MsgAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beheroes.doxchain.idp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Msg_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doxchain/idp/tx.proto",
}

func (m *MsgAuthenticationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tenant) > 0 {
		i -= len(m.Tenant)
		copy(dAtA[i:], m.Tenant)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Tenant)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAuthenticationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAuthenticationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAuthenticationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAuthenticationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Tenant)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAuthenticationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAuthenticationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAuthenticationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAuthenticationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAuthenticationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAuthenticationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
