// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/host_type_service.proto

package compute

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

// HostTypeServiceClient is the client API for HostTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostTypeServiceClient interface {
	// Returns information about specified host type.
	Get(ctx context.Context, in *GetHostTypeRequest, opts ...grpc.CallOption) (*HostType, error)
	// List avaliable host types.
	List(ctx context.Context, in *ListHostTypesRequest, opts ...grpc.CallOption) (*ListHostTypesResponse, error)
}

type hostTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostTypeServiceClient(cc grpc.ClientConnInterface) HostTypeServiceClient {
	return &hostTypeServiceClient{cc}
}

func (c *hostTypeServiceClient) Get(ctx context.Context, in *GetHostTypeRequest, opts ...grpc.CallOption) (*HostType, error) {
	out := new(HostType)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.HostTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostTypeServiceClient) List(ctx context.Context, in *ListHostTypesRequest, opts ...grpc.CallOption) (*ListHostTypesResponse, error) {
	out := new(ListHostTypesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.HostTypeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostTypeServiceServer is the server API for HostTypeService service.
// All implementations should embed UnimplementedHostTypeServiceServer
// for forward compatibility
type HostTypeServiceServer interface {
	// Returns information about specified host type.
	Get(context.Context, *GetHostTypeRequest) (*HostType, error)
	// List avaliable host types.
	List(context.Context, *ListHostTypesRequest) (*ListHostTypesResponse, error)
}

// UnimplementedHostTypeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHostTypeServiceServer struct {
}

func (UnimplementedHostTypeServiceServer) Get(context.Context, *GetHostTypeRequest) (*HostType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHostTypeServiceServer) List(context.Context, *ListHostTypesRequest) (*ListHostTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeHostTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostTypeServiceServer will
// result in compilation errors.
type UnsafeHostTypeServiceServer interface {
	mustEmbedUnimplementedHostTypeServiceServer()
}

func RegisterHostTypeServiceServer(s grpc.ServiceRegistrar, srv HostTypeServiceServer) {
	s.RegisterService(&HostTypeService_ServiceDesc, srv)
}

func _HostTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.HostTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostTypeServiceServer).Get(ctx, req.(*GetHostTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHostTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.HostTypeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostTypeServiceServer).List(ctx, req.(*ListHostTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostTypeService_ServiceDesc is the grpc.ServiceDesc for HostTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.HostTypeService",
	HandlerType: (*HostTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _HostTypeService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _HostTypeService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/host_type_service.proto",
}
