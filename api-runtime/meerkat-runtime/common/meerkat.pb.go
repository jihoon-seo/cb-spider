// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: meerkat.proto

package common

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message with empty
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// The response message containing a ChildKat status
type Status struct {
	ServerID             string   `protobuf:"bytes,1,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	Reserved             string   `protobuf:"bytes,4,opt,name=Reserved,proto3" json:"Reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32248fb0568de66, []int{1}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

func (m *Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Status) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Status) GetReserved() string {
	if m != nil {
		return m.Reserved
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "meerkatrpc.Empty")
	proto.RegisterType((*Status)(nil), "meerkatrpc.Status")
}

func init() { proto.RegisterFile("meerkat.proto", fileDescriptor_c32248fb0568de66) }

var fileDescriptor_c32248fb0568de66 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x4d, 0x2d,
	0xca, 0x4e, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x72, 0x8b, 0x0a, 0x92,
	0x95, 0xd8, 0xb9, 0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x95, 0x72, 0xb8, 0xd8, 0x82, 0x4b, 0x12,
	0x4b, 0x4a, 0x8b, 0x85, 0xa4, 0xb8, 0x38, 0x82, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x3c, 0x5d, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x31, 0x98, 0x2a, 0x09, 0x26, 0xb0, 0x0c,
	0x4c, 0x8f, 0x10, 0x17, 0x4b, 0x48, 0x66, 0x6e, 0xaa, 0x04, 0x33, 0x58, 0x14, 0xcc, 0x06, 0x99,
	0x13, 0x94, 0x5a, 0x0c, 0xd2, 0x99, 0x22, 0xc1, 0x02, 0x31, 0x07, 0xc6, 0x37, 0xf2, 0xe0, 0xe2,
	0x76, 0xce, 0xc8, 0xcc, 0x49, 0x81, 0x6a, 0xb7, 0xe4, 0xe2, 0x73, 0x4f, 0x2d, 0x41, 0x16, 0x11,
	0xd4, 0x43, 0x38, 0x52, 0x0f, 0xec, 0x42, 0x29, 0x21, 0x64, 0x21, 0x88, 0x32, 0x25, 0x06, 0xa7,
	0xf0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xd7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4,
	0x9c, 0xfc, 0xd2, 0x14, 0xdd, 0xa4, 0xc4, 0xa2, 0xcc, 0xe2, 0x92, 0x44, 0xfd, 0xe4, 0x24, 0xdd,
	0xe2, 0x82, 0xcc, 0x94, 0xd4, 0x22, 0xfd, 0xc4, 0x82, 0x4c, 0xdd, 0xa2, 0xd2, 0xbc, 0x92, 0xcc,
	0xdc, 0x54, 0x7d, 0xa8, 0xa9, 0x70, 0x7e, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x5e, 0x12, 0x1b, 0x38,
	0xb0, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x03, 0xc3, 0xb8, 0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChildStatusClient is the client API for ChildStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChildStatusClient interface {
	// Sends a request of Resource status
	GetChildStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
}

type childStatusClient struct {
	cc *grpc.ClientConn
}

func NewChildStatusClient(cc *grpc.ClientConn) ChildStatusClient {
	return &childStatusClient{cc}
}

func (c *childStatusClient) GetChildStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/meerkatrpc.ChildStatus/GetChildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChildStatusServer is the server API for ChildStatus service.
type ChildStatusServer interface {
	// Sends a request of Resource status
	GetChildStatus(context.Context, *Empty) (*Status, error)
}

// UnimplementedChildStatusServer can be embedded to have forward compatible implementations.
type UnimplementedChildStatusServer struct {
}

func (*UnimplementedChildStatusServer) GetChildStatus(ctx context.Context, req *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChildStatus not implemented")
}

func RegisterChildStatusServer(s *grpc.Server, srv ChildStatusServer) {
	s.RegisterService(&_ChildStatus_serviceDesc, srv)
}

func _ChildStatus_GetChildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChildStatusServer).GetChildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meerkatrpc.ChildStatus/GetChildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChildStatusServer).GetChildStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChildStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "meerkatrpc.ChildStatus",
	HandlerType: (*ChildStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChildStatus",
			Handler:    _ChildStatus_GetChildStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meerkat.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Time) > 0 {
		i -= len(m.Time)
		copy(dAtA[i:], m.Time)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Time)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintMeerkat(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeerkat(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeerkat(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovMeerkat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMeerkat(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeerkat(x uint64) (n int) {
	return sovMeerkat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeerkat
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
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
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
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
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
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
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeerkat
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
				return ErrInvalidLengthMeerkat
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeerkat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMeerkat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeerkat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMeerkat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeerkat
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
					return 0, ErrIntOverflowMeerkat
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
					return 0, ErrIntOverflowMeerkat
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
				return 0, ErrInvalidLengthMeerkat
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeerkat
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeerkat
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeerkat        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeerkat          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeerkat = fmt.Errorf("proto: unexpected end of group")
)