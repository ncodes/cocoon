// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	InvokeParam
	InvokeResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type InvokeParam struct {
	ID       string   `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Function string   `protobuf:"bytes,2,opt,name=function" json:"function,omitempty"`
	Params   []string `protobuf:"bytes,3,rep,name=params" json:"params,omitempty"`
}

func (m *InvokeParam) Reset()                    { *m = InvokeParam{} }
func (m *InvokeParam) String() string            { return proto1.CompactTextString(m) }
func (*InvokeParam) ProtoMessage()               {}
func (*InvokeParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

type InvokeResponse struct {
	ID     string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *InvokeResponse) Reset()                    { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string            { return proto1.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()               {}
func (*InvokeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	proto1.RegisterType((*InvokeParam)(nil), "proto.InvokeParam")
	proto1.RegisterType((*InvokeResponse)(nil), "proto.InvokeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Stub service

type StubClient interface {
	Invoke(ctx context.Context, in *InvokeParam, opts ...grpc.CallOption) (*InvokeResponse, error)
}

type stubClient struct {
	cc *grpc.ClientConn
}

func NewStubClient(cc *grpc.ClientConn) StubClient {
	return &stubClient{cc}
}

func (c *stubClient) Invoke(ctx context.Context, in *InvokeParam, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/proto.Stub/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stub service

type StubServer interface {
	Invoke(context.Context, *InvokeParam) (*InvokeResponse, error)
}

func RegisterStubServer(s *grpc.Server, srv StubServer) {
	s.RegisterService(&_Stub_serviceDesc, srv)
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
		FullMethod: "/proto.Stub/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StubServer).Invoke(ctx, req.(*InvokeParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Stub",
	HandlerType: (*StubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _Stub_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x81, 0x5c, 0xdc,
	0x9e, 0x79, 0x65, 0xf9, 0xd9, 0xa9, 0x01, 0x89, 0x45, 0x89, 0xb9, 0x42, 0x7c, 0x5c, 0x4c, 0x9e,
	0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x52, 0x5c, 0x1c, 0x69,
	0xa5, 0x79, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x51, 0x38, 0x5f, 0x48, 0x8c, 0x8b,
	0xad, 0x00, 0xa4, 0xa9, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0x33, 0x08, 0xca, 0x53, 0xf2, 0xe1,
	0xe2, 0x83, 0x18, 0x19, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x61, 0xaa, 0x18, 0x17,
	0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0xd8, 0x4c, 0xd6, 0x20, 0x28, 0x4f, 0x48, 0x88, 0x8b,
	0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x36, 0xb2, 0xe6,
	0x62, 0x09, 0x2e, 0x29, 0x4d, 0x12, 0x32, 0xe6, 0x62, 0x83, 0x98, 0x2a, 0x24, 0x04, 0xf1, 0x81,
	0x1e, 0x92, 0xbb, 0xa5, 0x44, 0x51, 0xc4, 0x60, 0x16, 0x27, 0xb1, 0x81, 0x45, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0xa2, 0x63, 0xa3, 0xfb, 0x00, 0x00, 0x00,
}
