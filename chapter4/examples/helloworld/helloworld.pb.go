// Code generated by protoc-gen-go.
// source: helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld.proto

It has these top-level messages:
	HelloRequest
	HelloReply
	TokenRequest
	TokenReply
*/
package helloworld

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TokenRequest struct {
	GrantType string `protobuf:"bytes,1,opt,name=grantType" json:"grantType,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Scope     string `protobuf:"bytes,4,opt,name=scope" json:"scope,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenRequest) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *TokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *TokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *TokenRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

type TokenReply struct {
	AccessToken      string `protobuf:"bytes,1,opt,name=accessToken" json:"accessToken,omitempty"`
	TokenType        string `protobuf:"bytes,2,opt,name=tokenType" json:"tokenType,omitempty"`
	ExpiresIn        int32  `protobuf:"varint,3,opt,name=expiresIn" json:"expiresIn,omitempty"`
	RefreshToken     string `protobuf:"bytes,4,opt,name=refreshToken" json:"refreshToken,omitempty"`
	ExampleParameter string `protobuf:"bytes,5,opt,name=exampleParameter" json:"exampleParameter,omitempty"`
}

func (m *TokenReply) Reset()                    { *m = TokenReply{} }
func (m *TokenReply) String() string            { return proto.CompactTextString(m) }
func (*TokenReply) ProtoMessage()               {}
func (*TokenReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *TokenReply) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *TokenReply) GetExpiresIn() int32 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *TokenReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *TokenReply) GetExampleParameter() string {
	if m != nil {
		return m.ExampleParameter
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*TokenRequest)(nil), "helloworld.TokenRequest")
	proto.RegisterType((*TokenReply)(nil), "helloworld.TokenReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Token(context.Context, *TokenRequest) (*TokenReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Greeter_Token_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x4a, 0xfb, 0x40,
	0x10, 0x6e, 0x7e, 0xbf, 0xc6, 0xb6, 0x63, 0xc1, 0xb2, 0x88, 0x84, 0xea, 0xa1, 0xec, 0x41, 0xc4,
	0x43, 0x10, 0x3d, 0x0a, 0x1e, 0x7a, 0x51, 0x6f, 0xa5, 0x16, 0x3c, 0xaf, 0xe9, 0xd8, 0x16, 0x93,
	0xec, 0x3a, 0x93, 0xd2, 0xc6, 0x07, 0xf0, 0xa1, 0x7c, 0x3a, 0xd9, 0x4d, 0x62, 0x12, 0xf4, 0x36,
	0xdf, 0x9f, 0xdd, 0xef, 0x63, 0x67, 0x61, 0xb4, 0xc6, 0x38, 0xd6, 0x3b, 0x4d, 0xf1, 0x32, 0x34,
	0xa4, 0x33, 0x2d, 0xa0, 0x66, 0xa4, 0x84, 0xe1, 0x83, 0x45, 0x73, 0x7c, 0xdf, 0x22, 0x67, 0x42,
	0x40, 0x37, 0x55, 0x09, 0x06, 0xde, 0xc4, 0xbb, 0x18, 0xcc, 0xdd, 0x2c, 0xcf, 0x01, 0x4a, 0x8f,
	0x89, 0x73, 0x11, 0x40, 0x2f, 0x41, 0x66, 0xb5, 0xaa, 0x4c, 0x15, 0x94, 0x1f, 0x30, 0x5c, 0xe8,
	0x37, 0x4c, 0xab, 0xbb, 0xce, 0x60, 0xb0, 0x22, 0x95, 0x66, 0x8b, 0xdc, 0x54, 0xde, 0x9a, 0x10,
	0x63, 0xe8, 0x6f, 0x19, 0xc9, 0xa5, 0xfd, 0x73, 0xe2, 0x0f, 0xb6, 0x9a, 0x51, 0xcc, 0x3b, 0x4d,
	0xcb, 0xe0, 0x7f, 0xa1, 0x55, 0x58, 0x1c, 0x83, 0xcf, 0x91, 0x36, 0x18, 0x74, 0x9d, 0x50, 0x00,
	0xf9, 0xe5, 0x01, 0x94, 0xe1, 0xb6, 0xe4, 0x04, 0x0e, 0x55, 0x14, 0x21, 0xb3, 0xe3, 0xca, 0xf0,
	0x26, 0x65, 0xcb, 0x65, 0x76, 0x70, 0xe5, 0x8a, 0xfc, 0x9a, 0xb0, 0x2a, 0xee, 0xcd, 0x86, 0x90,
	0x1f, 0x53, 0xd7, 0xc0, 0x9f, 0xd7, 0x84, 0x90, 0x30, 0x24, 0x7c, 0x25, 0xe4, 0x75, 0x71, 0x7d,
	0xd1, 0xa4, 0xc5, 0x89, 0x4b, 0x18, 0xe1, 0x5e, 0x25, 0x26, 0xc6, 0x99, 0x22, 0x95, 0x60, 0x86,
	0x14, 0xf8, 0xce, 0xf7, 0x8b, 0xbf, 0xfe, 0xf4, 0xa0, 0x77, 0x4f, 0x68, 0x67, 0x71, 0x07, 0xfd,
	0x27, 0x95, 0xbb, 0xf7, 0x16, 0x41, 0xd8, 0xd8, 0x5d, 0x73, 0x4d, 0xe3, 0x93, 0x3f, 0x14, 0x13,
	0xe7, 0xb2, 0x23, 0x6e, 0xc1, 0x2f, 0x0a, 0xb4, 0x0e, 0x37, 0xf7, 0xd2, 0x3e, 0x5c, 0x3f, 0x9a,
	0xec, 0x4c, 0xaf, 0xe0, 0x74, 0xa3, 0xc3, 0x15, 0x99, 0x28, 0x2c, 0x4b, 0x72, 0xc3, 0x3b, 0x3d,
	0x72, 0x49, 0xcf, 0x76, 0x9e, 0xd9, 0x9f, 0x34, 0xf3, 0x5e, 0x0e, 0xdc, 0x97, 0xba, 0xf9, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0xaa, 0xa1, 0x0a, 0x66, 0x02, 0x00, 0x00,
}