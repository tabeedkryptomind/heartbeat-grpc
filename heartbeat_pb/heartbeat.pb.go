// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heartbeat_pb/heartbeat.proto

package heartbeat_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type HeartBeat struct {
	Bpm                  int32    `protobuf:"varint,1,opt,name=bpm,proto3" json:"bpm,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd80506f8386f83, []int{0}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func (m *HeartBeat) GetBpm() int32 {
	if m != nil {
		return m.Bpm
	}
	return 0
}

func (m *HeartBeat) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type HeartBeatRequest struct {
	Heartbeat            *HeartBeat `protobuf:"bytes,1,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HeartBeatRequest) Reset()         { *m = HeartBeatRequest{} }
func (m *HeartBeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRequest) ProtoMessage()    {}
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd80506f8386f83, []int{1}
}

func (m *HeartBeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRequest.Unmarshal(m, b)
}
func (m *HeartBeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRequest.Marshal(b, m, deterministic)
}
func (m *HeartBeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRequest.Merge(m, src)
}
func (m *HeartBeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRequest.Size(m)
}
func (m *HeartBeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRequest proto.InternalMessageInfo

func (m *HeartBeatRequest) GetHeartbeat() *HeartBeat {
	if m != nil {
		return m.Heartbeat
	}
	return nil
}

type HeartBeatResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatResponse) Reset()         { *m = HeartBeatResponse{} }
func (m *HeartBeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartBeatResponse) ProtoMessage()    {}
func (*HeartBeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fd80506f8386f83, []int{2}
}

func (m *HeartBeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatResponse.Unmarshal(m, b)
}
func (m *HeartBeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatResponse.Marshal(b, m, deterministic)
}
func (m *HeartBeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatResponse.Merge(m, src)
}
func (m *HeartBeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartBeatResponse.Size(m)
}
func (m *HeartBeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatResponse proto.InternalMessageInfo

func (m *HeartBeatResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*HeartBeat)(nil), "heartbeat.HeartBeat")
	proto.RegisterType((*HeartBeatRequest)(nil), "heartbeat.HeartBeatRequest")
	proto.RegisterType((*HeartBeatResponse)(nil), "heartbeat.HeartBeatResponse")
}

func init() { proto.RegisterFile("heartbeat_pb/heartbeat.proto", fileDescriptor_7fd80506f8386f83) }

var fileDescriptor_7fd80506f8386f83 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0x48, 0x4d, 0x2c,
	0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0x89, 0x2f, 0x48, 0xd2, 0x87, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x38, 0xe1, 0x02, 0x4a, 0x96, 0x5c, 0x9c, 0x1e, 0x20, 0x8e, 0x53, 0x6a, 0x62, 0x89,
	0x90, 0x00, 0x17, 0x73, 0x52, 0x41, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0x29,
	0x24, 0xc5, 0xc5, 0x51, 0x5a, 0x9c, 0x5a, 0x94, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x19, 0x04, 0xe7, 0x2b, 0xb9, 0x71, 0x09, 0xc0, 0xb5, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x19, 0x71, 0x21, 0xcc, 0x06, 0x9b, 0xc3, 0x6d, 0x24, 0xa2, 0x87, 0xb0, 0x1e, 0xa1,
	0x1e, 0xc9, 0x09, 0xda, 0x5c, 0x82, 0x48, 0xe6, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89,
	0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x40, 0x4c, 0xe1, 0x0c, 0x82, 0xf2, 0x8c, 0x12, 0x90,
	0x2c, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xf2, 0xe1, 0xe2, 0x0d, 0x2d, 0x4e, 0x2d,
	0x42, 0xf8, 0x43, 0x1a, 0xab, 0x95, 0x10, 0x27, 0x4a, 0xc9, 0x60, 0x97, 0x84, 0xd8, 0xab, 0xc4,
	0xe0, 0xc4, 0x1f, 0xc5, 0xab, 0x8f, 0x1c, 0x7a, 0x49, 0x6c, 0xe0, 0x40, 0x33, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xa1, 0x40, 0x58, 0x9d, 0x54, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartBeatServiceClient is the client API for HeartBeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartBeatServiceClient interface {
	UserHeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
}

type heartBeatServiceClient struct {
	cc *grpc.ClientConn
}

func NewHeartBeatServiceClient(cc *grpc.ClientConn) HeartBeatServiceClient {
	return &heartBeatServiceClient{cc}
}

func (c *heartBeatServiceClient) UserHeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/heartbeat.HeartBeatService/UserHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServiceServer is the server API for HeartBeatService service.
type HeartBeatServiceServer interface {
	UserHeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
}

// UnimplementedHeartBeatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServiceServer struct {
}

func (*UnimplementedHeartBeatServiceServer) UserHeartBeat(ctx context.Context, req *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHeartBeat not implemented")
}

func RegisterHeartBeatServiceServer(s *grpc.Server, srv HeartBeatServiceServer) {
	s.RegisterService(&_HeartBeatService_serviceDesc, srv)
}

func _HeartBeatService_UserHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServiceServer).UserHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat.HeartBeatService/UserHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServiceServer).UserHeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeartBeatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heartbeat.HeartBeatService",
	HandlerType: (*HeartBeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserHeartBeat",
			Handler:    _HeartBeatService_UserHeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat_pb/heartbeat.proto",
}
