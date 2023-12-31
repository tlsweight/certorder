// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/compute/v1/disk_type_service.proto

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

// DiskTypeServiceClient is the client API for DiskTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiskTypeServiceClient interface {
	// Returns the information about specified disk type.
	//
	// To get the list of available disk types, make a [List] request.
	Get(ctx context.Context, in *GetDiskTypeRequest, opts ...grpc.CallOption) (*DiskType, error)
	// Retrieves the list of disk types for the specified folder.
	List(ctx context.Context, in *ListDiskTypesRequest, opts ...grpc.CallOption) (*ListDiskTypesResponse, error)
}

type diskTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiskTypeServiceClient(cc grpc.ClientConnInterface) DiskTypeServiceClient {
	return &diskTypeServiceClient{cc}
}

func (c *diskTypeServiceClient) Get(ctx context.Context, in *GetDiskTypeRequest, opts ...grpc.CallOption) (*DiskType, error) {
	out := new(DiskType)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diskTypeServiceClient) List(ctx context.Context, in *ListDiskTypesRequest, opts ...grpc.CallOption) (*ListDiskTypesResponse, error) {
	out := new(ListDiskTypesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.DiskTypeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiskTypeServiceServer is the server API for DiskTypeService service.
// All implementations should embed UnimplementedDiskTypeServiceServer
// for forward compatibility
type DiskTypeServiceServer interface {
	// Returns the information about specified disk type.
	//
	// To get the list of available disk types, make a [List] request.
	Get(context.Context, *GetDiskTypeRequest) (*DiskType, error)
	// Retrieves the list of disk types for the specified folder.
	List(context.Context, *ListDiskTypesRequest) (*ListDiskTypesResponse, error)
}

// UnimplementedDiskTypeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDiskTypeServiceServer struct {
}

func (UnimplementedDiskTypeServiceServer) Get(context.Context, *GetDiskTypeRequest) (*DiskType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDiskTypeServiceServer) List(context.Context, *ListDiskTypesRequest) (*ListDiskTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeDiskTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiskTypeServiceServer will
// result in compilation errors.
type UnsafeDiskTypeServiceServer interface {
	mustEmbedUnimplementedDiskTypeServiceServer()
}

func RegisterDiskTypeServiceServer(s grpc.ServiceRegistrar, srv DiskTypeServiceServer) {
	s.RegisterService(&DiskTypeService_ServiceDesc, srv)
}

func _DiskTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiskTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskTypeServiceServer).Get(ctx, req.(*GetDiskTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiskTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiskTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiskTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.DiskTypeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiskTypeServiceServer).List(ctx, req.(*ListDiskTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiskTypeService_ServiceDesc is the grpc.ServiceDesc for DiskTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiskTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.DiskTypeService",
	HandlerType: (*DiskTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DiskTypeService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiskTypeService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/disk_type_service.proto",
}
