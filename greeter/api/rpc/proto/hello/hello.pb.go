// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/examples/greeter/api/rpc/proto/hello/hello.proto

/*
Package go_micro_api_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/greeter/api/rpc/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_api_greeter

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.api.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.api.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.api.greeter.Greeter/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.api.greeter.Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.api.greeter.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/examples/greeter/api/rpc/proto/hello/hello.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/examples/greeter/api/rpc/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xc1, 0xaa, 0xc2, 0x30,
	0x10, 0x45, 0x5f, 0x79, 0x6a, 0x35, 0x2b, 0x09, 0x2e, 0x44, 0xac, 0x48, 0x57, 0xae, 0x66, 0x40,
	0xbf, 0xc0, 0x95, 0x5d, 0xd7, 0x2f, 0x48, 0xcb, 0x90, 0x06, 0x9a, 0x26, 0x26, 0x29, 0xf8, 0xf9,
	0xd2, 0x34, 0x4b, 0xdd, 0x0c, 0x97, 0xb9, 0x9c, 0x39, 0xc3, 0xee, 0x52, 0x85, 0x6e, 0x6c, 0xa0,
	0x35, 0x1a, 0xb5, 0x6a, 0x9d, 0x41, 0x7a, 0x0b, 0x6d, 0x7b, 0xf2, 0x28, 0x1d, 0x51, 0x20, 0x87,
	0xc2, 0x2a, 0x74, 0xb6, 0x45, 0xeb, 0x4c, 0x30, 0xd8, 0x51, 0xdf, 0xa7, 0x09, 0x71, 0xc3, 0x77,
	0xd2, 0x40, 0x44, 0x41, 0x58, 0x05, 0x89, 0x2a, 0x0b, 0x96, 0xd7, 0xf4, 0x1a, 0xc9, 0x07, 0xce,
	0xd9, 0x62, 0x10, 0x9a, 0xf6, 0xd9, 0x39, 0xbb, 0x6c, 0xea, 0x98, 0xcb, 0x23, 0x5b, 0xd7, 0xe4,
	0xad, 0x19, 0x3c, 0xf1, 0x2d, 0xfb, 0xd7, 0x5e, 0xa6, 0x7a, 0x8a, 0xd7, 0x27, 0xcb, 0x1f, 0xf3,
	0x1d, 0x5e, 0xb1, 0x65, 0x35, 0xc9, 0x78, 0x01, 0xdf, 0x3c, 0x90, 0x24, 0x87, 0xd3, 0xaf, 0x7a,
	0x96, 0x94, 0x7f, 0xcd, 0x2a, 0xbe, 0x7b, 0xfb, 0x04, 0x00, 0x00, 0xff, 0xff, 0x04, 0xed, 0xc3,
	0x67, 0xf3, 0x00, 0x00, 0x00,
}
