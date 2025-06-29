// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/gogurt_vendor.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GogurtVendor_GetGogurt_FullMethodName = "/proto.GogurtVendor/GetGogurt"
)

// GogurtVendorClient is the client API for GogurtVendor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GogurtVendorClient interface {
	GetGogurt(ctx context.Context, in *GetGogurtRequest, opts ...grpc.CallOption) (*Gogurt, error)
}

type gogurtVendorClient struct {
	cc grpc.ClientConnInterface
}

func NewGogurtVendorClient(cc grpc.ClientConnInterface) GogurtVendorClient {
	return &gogurtVendorClient{cc}
}

func (c *gogurtVendorClient) GetGogurt(ctx context.Context, in *GetGogurtRequest, opts ...grpc.CallOption) (*Gogurt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Gogurt)
	err := c.cc.Invoke(ctx, GogurtVendor_GetGogurt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GogurtVendorServer is the server API for GogurtVendor service.
// All implementations must embed UnimplementedGogurtVendorServer
// for forward compatibility.
type GogurtVendorServer interface {
	GetGogurt(context.Context, *GetGogurtRequest) (*Gogurt, error)
	mustEmbedUnimplementedGogurtVendorServer()
}

// UnimplementedGogurtVendorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGogurtVendorServer struct{}

func (UnimplementedGogurtVendorServer) GetGogurt(context.Context, *GetGogurtRequest) (*Gogurt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGogurt not implemented")
}
func (UnimplementedGogurtVendorServer) mustEmbedUnimplementedGogurtVendorServer() {}
func (UnimplementedGogurtVendorServer) testEmbeddedByValue()                      {}

// UnsafeGogurtVendorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GogurtVendorServer will
// result in compilation errors.
type UnsafeGogurtVendorServer interface {
	mustEmbedUnimplementedGogurtVendorServer()
}

func RegisterGogurtVendorServer(s grpc.ServiceRegistrar, srv GogurtVendorServer) {
	// If the following call pancis, it indicates UnimplementedGogurtVendorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GogurtVendor_ServiceDesc, srv)
}

func _GogurtVendor_GetGogurt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGogurtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GogurtVendorServer).GetGogurt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GogurtVendor_GetGogurt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GogurtVendorServer).GetGogurt(ctx, req.(*GetGogurtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GogurtVendor_ServiceDesc is the grpc.ServiceDesc for GogurtVendor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GogurtVendor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GogurtVendor",
	HandlerType: (*GogurtVendorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGogurt",
			Handler:    _GogurtVendor_GetGogurt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gogurt_vendor.proto",
}
