// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/vpc/v1/route_table_service.proto

package vpc

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RouteTableServiceClient is the client API for RouteTableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteTableServiceClient interface {
	// Returns the specified RouteTable resource.
	//
	// To get the list of available RouteTable resources, make a [List] request.
	Get(ctx context.Context, in *GetRouteTableRequest, opts ...grpc.CallOption) (*RouteTable, error)
	// Retrieves the list of RouteTable resources in the specified folder.
	List(ctx context.Context, in *ListRouteTablesRequest, opts ...grpc.CallOption) (*ListRouteTablesResponse, error)
	// Creates a route table in the specified folder and network.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(ctx context.Context, in *CreateRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified route table.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(ctx context.Context, in *UpdateRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified route table.
	Delete(ctx context.Context, in *DeleteRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// List operations for the specified route table.
	ListOperations(ctx context.Context, in *ListRouteTableOperationsRequest, opts ...grpc.CallOption) (*ListRouteTableOperationsResponse, error)
	// Move route table to another folder.
	Move(ctx context.Context, in *MoveRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type routeTableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteTableServiceClient(cc grpc.ClientConnInterface) RouteTableServiceClient {
	return &routeTableServiceClient{cc}
}

func (c *routeTableServiceClient) Get(ctx context.Context, in *GetRouteTableRequest, opts ...grpc.CallOption) (*RouteTable, error) {
	out := new(RouteTable)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) List(ctx context.Context, in *ListRouteTablesRequest, opts ...grpc.CallOption) (*ListRouteTablesResponse, error) {
	out := new(ListRouteTablesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) Create(ctx context.Context, in *CreateRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) Update(ctx context.Context, in *UpdateRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) Delete(ctx context.Context, in *DeleteRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) ListOperations(ctx context.Context, in *ListRouteTableOperationsRequest, opts ...grpc.CallOption) (*ListRouteTableOperationsResponse, error) {
	out := new(ListRouteTableOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTableServiceClient) Move(ctx context.Context, in *MoveRouteTableRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.vpc.v1.RouteTableService/Move", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteTableServiceServer is the server API for RouteTableService service.
// All implementations should embed UnimplementedRouteTableServiceServer
// for forward compatibility
type RouteTableServiceServer interface {
	// Returns the specified RouteTable resource.
	//
	// To get the list of available RouteTable resources, make a [List] request.
	Get(context.Context, *GetRouteTableRequest) (*RouteTable, error)
	// Retrieves the list of RouteTable resources in the specified folder.
	List(context.Context, *ListRouteTablesRequest) (*ListRouteTablesResponse, error)
	// Creates a route table in the specified folder and network.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(context.Context, *CreateRouteTableRequest) (*operation.Operation, error)
	// Updates the specified route table.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(context.Context, *UpdateRouteTableRequest) (*operation.Operation, error)
	// Deletes the specified route table.
	Delete(context.Context, *DeleteRouteTableRequest) (*operation.Operation, error)
	// List operations for the specified route table.
	ListOperations(context.Context, *ListRouteTableOperationsRequest) (*ListRouteTableOperationsResponse, error)
	// Move route table to another folder.
	Move(context.Context, *MoveRouteTableRequest) (*operation.Operation, error)
}

// UnimplementedRouteTableServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRouteTableServiceServer struct {
}

func (UnimplementedRouteTableServiceServer) Get(context.Context, *GetRouteTableRequest) (*RouteTable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRouteTableServiceServer) List(context.Context, *ListRouteTablesRequest) (*ListRouteTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRouteTableServiceServer) Create(context.Context, *CreateRouteTableRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRouteTableServiceServer) Update(context.Context, *UpdateRouteTableRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRouteTableServiceServer) Delete(context.Context, *DeleteRouteTableRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRouteTableServiceServer) ListOperations(context.Context, *ListRouteTableOperationsRequest) (*ListRouteTableOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedRouteTableServiceServer) Move(context.Context, *MoveRouteTableRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}

// UnsafeRouteTableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteTableServiceServer will
// result in compilation errors.
type UnsafeRouteTableServiceServer interface {
	mustEmbedUnimplementedRouteTableServiceServer()
}

func RegisterRouteTableServiceServer(s grpc.ServiceRegistrar, srv RouteTableServiceServer) {
	s.RegisterService(&RouteTableService_ServiceDesc, srv)
}

func _RouteTableService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).Get(ctx, req.(*GetRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRouteTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).List(ctx, req.(*ListRouteTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).Create(ctx, req.(*CreateRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).Update(ctx, req.(*UpdateRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).Delete(ctx, req.(*DeleteRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRouteTableOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).ListOperations(ctx, req.(*ListRouteTableOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTableService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRouteTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTableServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.vpc.v1.RouteTableService/Move",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTableServiceServer).Move(ctx, req.(*MoveRouteTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RouteTableService_ServiceDesc is the grpc.ServiceDesc for RouteTableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RouteTableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.vpc.v1.RouteTableService",
	HandlerType: (*RouteTableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RouteTableService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RouteTableService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RouteTableService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RouteTableService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RouteTableService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _RouteTableService_ListOperations_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _RouteTableService_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/vpc/v1/route_table_service.proto",
}