// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package shoeslxdmulti

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShoesLXDMultiClient is the client API for ShoesLXDMulti service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShoesLXDMultiClient interface {
	AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*AddInstanceResponse, error)
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error)
}

type shoesLXDMultiClient struct {
	cc grpc.ClientConnInterface
}

func NewShoesLXDMultiClient(cc grpc.ClientConnInterface) ShoesLXDMultiClient {
	return &shoesLXDMultiClient{cc}
}

func (c *shoesLXDMultiClient) AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*AddInstanceResponse, error) {
	out := new(AddInstanceResponse)
	err := c.cc.Invoke(ctx, "/shoeslxdmulti.ShoesLXDMulti/AddInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoesLXDMultiClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceResponse, error) {
	out := new(DeleteInstanceResponse)
	err := c.cc.Invoke(ctx, "/shoeslxdmulti.ShoesLXDMulti/DeleteInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoesLXDMultiServer is the server API for ShoesLXDMulti service.
// All implementations must embed UnimplementedShoesLXDMultiServer
// for forward compatibility
type ShoesLXDMultiServer interface {
	AddInstance(context.Context, *AddInstanceRequest) (*AddInstanceResponse, error)
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error)
	mustEmbedUnimplementedShoesLXDMultiServer()
}

// UnimplementedShoesLXDMultiServer must be embedded to have forward compatible implementations.
type UnimplementedShoesLXDMultiServer struct {
}

func (UnimplementedShoesLXDMultiServer) AddInstance(context.Context, *AddInstanceRequest) (*AddInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInstance not implemented")
}
func (UnimplementedShoesLXDMultiServer) DeleteInstance(context.Context, *DeleteInstanceRequest) (*DeleteInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstance not implemented")
}
func (UnimplementedShoesLXDMultiServer) mustEmbedUnimplementedShoesLXDMultiServer() {}

// UnsafeShoesLXDMultiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShoesLXDMultiServer will
// result in compilation errors.
type UnsafeShoesLXDMultiServer interface {
	mustEmbedUnimplementedShoesLXDMultiServer()
}

func RegisterShoesLXDMultiServer(s grpc.ServiceRegistrar, srv ShoesLXDMultiServer) {
	s.RegisterService(&ShoesLXDMulti_ServiceDesc, srv)
}

func _ShoesLXDMulti_AddInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoesLXDMultiServer).AddInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoeslxdmulti.ShoesLXDMulti/AddInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoesLXDMultiServer).AddInstance(ctx, req.(*AddInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoesLXDMulti_DeleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoesLXDMultiServer).DeleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoeslxdmulti.ShoesLXDMulti/DeleteInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoesLXDMultiServer).DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShoesLXDMulti_ServiceDesc is the grpc.ServiceDesc for ShoesLXDMulti service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShoesLXDMulti_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shoeslxdmulti.ShoesLXDMulti",
	HandlerType: (*ShoesLXDMultiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInstance",
			Handler:    _ShoesLXDMulti_AddInstance_Handler,
		},
		{
			MethodName: "DeleteInstance",
			Handler:    _ShoesLXDMulti_DeleteInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shoes-lxd-multi.proto",
}
