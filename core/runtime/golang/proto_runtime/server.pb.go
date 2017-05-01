// Code generated by protoc-gen-gogo.
// source: server.proto
// DO NOT EDIT!

/*
Package proto_runtime is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	Ok
	InvokeParam
	InvokeResponse
*/
package proto_runtime

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Ok struct {
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{0} }

func (m *Ok) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type InvokeParam struct {
	ID       string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Function string            `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	Params   []string          `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
	Header   map[string]string `protobuf:"bytes,4,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InvokeParam) Reset()                    { *m = InvokeParam{} }
func (m *InvokeParam) String() string            { return proto.CompactTextString(m) }
func (*InvokeParam) ProtoMessage()               {}
func (*InvokeParam) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{1} }

func (m *InvokeParam) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *InvokeParam) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *InvokeParam) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *InvokeParam) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

type InvokeResponse struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *InvokeResponse) Reset()                    { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string            { return proto.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()               {}
func (*InvokeResponse) Descriptor() ([]byte, []int) { return fileDescriptorServer, []int{2} }

func (m *InvokeResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *InvokeResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *InvokeResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Ok)(nil), "proto_runtime.Ok")
	proto.RegisterType((*InvokeParam)(nil), "proto_runtime.InvokeParam")
	proto.RegisterType((*InvokeResponse)(nil), "proto_runtime.InvokeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Stub service

type StubClient interface {
	HealthCheck(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Ok, error)
	Invoke(ctx context.Context, in *InvokeParam, opts ...grpc.CallOption) (*InvokeResponse, error)
}

type stubClient struct {
	cc *grpc.ClientConn
}

func NewStubClient(cc *grpc.ClientConn) StubClient {
	return &stubClient{cc}
}

func (c *stubClient) HealthCheck(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := grpc.Invoke(ctx, "/proto_runtime.Stub/HealthCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stubClient) Invoke(ctx context.Context, in *InvokeParam, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/proto_runtime.Stub/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stub service

type StubServer interface {
	HealthCheck(context.Context, *Ok) (*Ok, error)
	Invoke(context.Context, *InvokeParam) (*InvokeResponse, error)
}

func RegisterStubServer(s *grpc.Server, srv StubServer) {
	s.RegisterService(&_Stub_serviceDesc, srv)
}

func _Stub_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ok)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StubServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_runtime.Stub/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StubServer).HealthCheck(ctx, req.(*Ok))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stub_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StubServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_runtime.Stub/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StubServer).Invoke(ctx, req.(*InvokeParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_runtime.Stub",
	HandlerType: (*StubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Stub_HealthCheck_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Stub_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptorServer) }

var fileDescriptorServer = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0x4d, 0x1b, 0xec, 0xb4, 0x16, 0x1d, 0x44, 0x42, 0x50, 0x08, 0x3d, 0x48, 0x4e,
	0x39, 0xb4, 0x17, 0xf5, 0xe0, 0xa5, 0x15, 0x2c, 0x08, 0x95, 0xf5, 0x01, 0x24, 0x69, 0x57, 0x52,
	0xd2, 0xee, 0x86, 0xdd, 0x4d, 0x20, 0x37, 0x1f, 0xd0, 0x87, 0x92, 0x6c, 0xa2, 0x46, 0xab, 0xa7,
	0x9d, 0x6f, 0xfe, 0xed, 0x6f, 0x3e, 0x18, 0x69, 0xae, 0x4a, 0xae, 0xa2, 0x5c, 0x49, 0x23, 0xf1,
	0xd8, 0x3e, 0x2f, 0xaa, 0x10, 0x66, 0xbb, 0xe7, 0x93, 0x0b, 0xa0, 0xab, 0x0c, 0xcf, 0xc1, 0xd5,
	0x26, 0x36, 0x85, 0xf6, 0x48, 0x40, 0xc2, 0x3e, 0x6b, 0xd5, 0xe4, 0x9d, 0xc0, 0x70, 0x29, 0x4a,
	0x99, 0xf1, 0xa7, 0x58, 0xc5, 0x7b, 0x1c, 0x03, 0x5d, 0x2e, 0x6c, 0xcf, 0x80, 0xd1, 0xe5, 0x02,
	0x7d, 0x38, 0x7a, 0x2d, 0xc4, 0xda, 0x6c, 0xa5, 0xf0, 0xa8, 0xcd, 0x7e, 0xe9, 0x7a, 0x67, 0x5e,
	0x0f, 0x69, 0xcf, 0x09, 0x9c, 0x70, 0xc0, 0x5a, 0x85, 0x77, 0xe0, 0xa6, 0x3c, 0xde, 0x70, 0xe5,
	0xf5, 0x02, 0x27, 0x1c, 0x4e, 0xaf, 0xa2, 0x1f, 0x44, 0x51, 0xe7, 0xbf, 0xe8, 0xc1, 0x36, 0xde,
	0x0b, 0xa3, 0x2a, 0xd6, 0x4e, 0xf9, 0x37, 0x30, 0xec, 0xa4, 0xf1, 0x04, 0x9c, 0x8c, 0x57, 0x2d,
	0x53, 0x1d, 0xe2, 0x19, 0xf4, 0xcb, 0x78, 0x57, 0xf0, 0x96, 0xa8, 0x11, 0xb7, 0xf4, 0x9a, 0x4c,
	0x1e, 0x61, 0xdc, 0x6c, 0x67, 0x5c, 0xe7, 0x52, 0x68, 0x7e, 0x70, 0xd0, 0xb7, 0x11, 0xb4, 0x6b,
	0x04, 0x22, 0xf4, 0x12, 0xb9, 0xa9, 0x3c, 0x27, 0x20, 0xe1, 0x88, 0xd9, 0x78, 0xfa, 0x46, 0xa0,
	0xf7, 0x6c, 0x8a, 0x04, 0x67, 0x96, 0x68, 0x67, 0xd2, 0x79, 0xca, 0xd7, 0x19, 0x9e, 0xfe, 0x3a,
	0x68, 0x95, 0xf9, 0x87, 0x29, 0x9c, 0x83, 0xdb, 0xb0, 0xa0, 0xff, 0xbf, 0x01, 0xfe, 0xe5, 0x9f,
	0xb5, 0x4f, 0xfc, 0xc4, 0xb5, 0xd5, 0xd9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xa4, 0xaa,
	0x82, 0xe3, 0x01, 0x00, 0x00,
}
