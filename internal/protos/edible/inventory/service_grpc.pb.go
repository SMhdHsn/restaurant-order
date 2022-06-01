// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/edible/inventory/service.proto

package eipb

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

// EdibleInventoryServiceClient is the client API for EdibleInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdibleInventoryServiceClient interface {
	// Unary
	Use(ctx context.Context, in *InventoryUseRequest, opts ...grpc.CallOption) (*InventoryUseResponse, error)
}

type edibleInventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEdibleInventoryServiceClient(cc grpc.ClientConnInterface) EdibleInventoryServiceClient {
	return &edibleInventoryServiceClient{cc}
}

func (c *edibleInventoryServiceClient) Use(ctx context.Context, in *InventoryUseRequest, opts ...grpc.CallOption) (*InventoryUseResponse, error) {
	out := new(InventoryUseResponse)
	err := c.cc.Invoke(ctx, "/edible.inventory.service.EdibleInventoryService/Use", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdibleInventoryServiceServer is the server API for EdibleInventoryService service.
// All implementations should embed UnimplementedEdibleInventoryServiceServer
// for forward compatibility
type EdibleInventoryServiceServer interface {
	// Unary
	Use(context.Context, *InventoryUseRequest) (*InventoryUseResponse, error)
}

// UnimplementedEdibleInventoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEdibleInventoryServiceServer struct {
}

func (UnimplementedEdibleInventoryServiceServer) Use(context.Context, *InventoryUseRequest) (*InventoryUseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Use not implemented")
}

// UnsafeEdibleInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdibleInventoryServiceServer will
// result in compilation errors.
type UnsafeEdibleInventoryServiceServer interface {
	mustEmbedUnimplementedEdibleInventoryServiceServer()
}

func RegisterEdibleInventoryServiceServer(s grpc.ServiceRegistrar, srv EdibleInventoryServiceServer) {
	s.RegisterService(&EdibleInventoryService_ServiceDesc, srv)
}

func _EdibleInventoryService_Use_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InventoryUseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdibleInventoryServiceServer).Use(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edible.inventory.service.EdibleInventoryService/Use",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdibleInventoryServiceServer).Use(ctx, req.(*InventoryUseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EdibleInventoryService_ServiceDesc is the grpc.ServiceDesc for EdibleInventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EdibleInventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edible.inventory.service.EdibleInventoryService",
	HandlerType: (*EdibleInventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Use",
			Handler:    _EdibleInventoryService_Use_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/edible/inventory/service.proto",
}
