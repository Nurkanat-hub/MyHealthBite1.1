// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/menu.proto

package menu

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
	MenuService_GetAllDishes_FullMethodName = "/menu.MenuService/GetAllDishes"
	MenuService_CreateDish_FullMethodName   = "/menu.MenuService/CreateDish"
	MenuService_GetDishById_FullMethodName  = "/menu.MenuService/GetDishById"
	MenuService_UpdateDish_FullMethodName   = "/menu.MenuService/UpdateDish"
	MenuService_DeleteDish_FullMethodName   = "/menu.MenuService/DeleteDish"
)

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	GetAllDishes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DishList, error)
	CreateDish(ctx context.Context, in *CreateDishRequest, opts ...grpc.CallOption) (*Dish, error)
	GetDishById(ctx context.Context, in *DishIdRequest, opts ...grpc.CallOption) (*Dish, error)
	UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*Dish, error)
	DeleteDish(ctx context.Context, in *DishIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) GetAllDishes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DishList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DishList)
	err := c.cc.Invoke(ctx, MenuService_GetAllDishes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateDish(ctx context.Context, in *CreateDishRequest, opts ...grpc.CallOption) (*Dish, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dish)
	err := c.cc.Invoke(ctx, MenuService_CreateDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetDishById(ctx context.Context, in *DishIdRequest, opts ...grpc.CallOption) (*Dish, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dish)
	err := c.cc.Invoke(ctx, MenuService_GetDishById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*Dish, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Dish)
	err := c.cc.Invoke(ctx, MenuService_UpdateDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteDish(ctx context.Context, in *DishIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, MenuService_DeleteDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility.
type MenuServiceServer interface {
	GetAllDishes(context.Context, *Empty) (*DishList, error)
	CreateDish(context.Context, *CreateDishRequest) (*Dish, error)
	GetDishById(context.Context, *DishIdRequest) (*Dish, error)
	UpdateDish(context.Context, *UpdateDishRequest) (*Dish, error)
	DeleteDish(context.Context, *DishIdRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMenuServiceServer struct{}

func (UnimplementedMenuServiceServer) GetAllDishes(context.Context, *Empty) (*DishList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDishes not implemented")
}
func (UnimplementedMenuServiceServer) CreateDish(context.Context, *CreateDishRequest) (*Dish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDish not implemented")
}
func (UnimplementedMenuServiceServer) GetDishById(context.Context, *DishIdRequest) (*Dish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDishById not implemented")
}
func (UnimplementedMenuServiceServer) UpdateDish(context.Context, *UpdateDishRequest) (*Dish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDish not implemented")
}
func (UnimplementedMenuServiceServer) DeleteDish(context.Context, *DishIdRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDish not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}
func (UnimplementedMenuServiceServer) testEmbeddedByValue()                     {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	// If the following call pancis, it indicates UnimplementedMenuServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_GetAllDishes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetAllDishes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetAllDishes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetAllDishes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateDish(ctx, req.(*CreateDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetDishById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DishIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetDishById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetDishById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetDishById(ctx, req.(*DishIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateDish(ctx, req.(*UpdateDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DishIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_DeleteDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteDish(ctx, req.(*DishIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllDishes",
			Handler:    _MenuService_GetAllDishes_Handler,
		},
		{
			MethodName: "CreateDish",
			Handler:    _MenuService_CreateDish_Handler,
		},
		{
			MethodName: "GetDishById",
			Handler:    _MenuService_GetDishById_Handler,
		},
		{
			MethodName: "UpdateDish",
			Handler:    _MenuService_UpdateDish_Handler,
		},
		{
			MethodName: "DeleteDish",
			Handler:    _MenuService_DeleteDish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/menu.proto",
}
