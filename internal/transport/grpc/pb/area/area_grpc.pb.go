// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: area.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AreaClient is the client API for Area service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AreaClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*AreaList, error)
	CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*AreaFull, error)
	UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteArea(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*empty.Empty, error)
	Show(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*AreaFull, error)
}

type areaClient struct {
	cc grpc.ClientConnInterface
}

func NewAreaClient(cc grpc.ClientConnInterface) AreaClient {
	return &areaClient{cc}
}

func (c *areaClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/area.Area/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*AreaList, error) {
	out := new(AreaList)
	err := c.cc.Invoke(ctx, "/area.Area/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) CreateArea(ctx context.Context, in *CreateAreaRequest, opts ...grpc.CallOption) (*AreaFull, error) {
	out := new(AreaFull)
	err := c.cc.Invoke(ctx, "/area.Area/CreateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) UpdateArea(ctx context.Context, in *UpdateAreaRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/area.Area/UpdateArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) DeleteArea(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/area.Area/DeleteArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaClient) Show(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*AreaFull, error) {
	out := new(AreaFull)
	err := c.cc.Invoke(ctx, "/area.Area/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AreaServer is the server API for Area service.
// All implementations must embed UnimplementedAreaServer
// for forward compatibility
type AreaServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	List(context.Context, *ListRequest) (*AreaList, error)
	CreateArea(context.Context, *CreateAreaRequest) (*AreaFull, error)
	UpdateArea(context.Context, *UpdateAreaRequest) (*empty.Empty, error)
	DeleteArea(context.Context, *Identity) (*empty.Empty, error)
	Show(context.Context, *Identity) (*AreaFull, error)
	mustEmbedUnimplementedAreaServer()
}

// UnimplementedAreaServer must be embedded to have forward compatible implementations.
type UnimplementedAreaServer struct {
}

func (UnimplementedAreaServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAreaServer) List(context.Context, *ListRequest) (*AreaList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAreaServer) CreateArea(context.Context, *CreateAreaRequest) (*AreaFull, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArea not implemented")
}
func (UnimplementedAreaServer) UpdateArea(context.Context, *UpdateAreaRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArea not implemented")
}
func (UnimplementedAreaServer) DeleteArea(context.Context, *Identity) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArea not implemented")
}
func (UnimplementedAreaServer) Show(context.Context, *Identity) (*AreaFull, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Show not implemented")
}
func (UnimplementedAreaServer) mustEmbedUnimplementedAreaServer() {}

// UnsafeAreaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AreaServer will
// result in compilation errors.
type UnsafeAreaServer interface {
	mustEmbedUnimplementedAreaServer()
}

func RegisterAreaServer(s grpc.ServiceRegistrar, srv AreaServer) {
	s.RegisterService(&Area_ServiceDesc, srv)
}

func _Area_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_CreateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).CreateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/CreateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).CreateArea(ctx, req.(*CreateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_UpdateArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAreaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).UpdateArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/UpdateArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).UpdateArea(ctx, req.(*UpdateAreaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_DeleteArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).DeleteArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/DeleteArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).DeleteArea(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Area_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AreaServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/area.Area/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AreaServer).Show(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

// Area_ServiceDesc is the grpc.ServiceDesc for Area service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Area_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "area.Area",
	HandlerType: (*AreaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Area_Ping_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Area_List_Handler,
		},
		{
			MethodName: "CreateArea",
			Handler:    _Area_CreateArea_Handler,
		},
		{
			MethodName: "UpdateArea",
			Handler:    _Area_UpdateArea_Handler,
		},
		{
			MethodName: "DeleteArea",
			Handler:    _Area_DeleteArea_Handler,
		},
		{
			MethodName: "Show",
			Handler:    _Area_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "area.proto",
}