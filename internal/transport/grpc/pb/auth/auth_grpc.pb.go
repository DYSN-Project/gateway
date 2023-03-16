// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error)
	ConfirmRegister(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Tokens, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Tokens, error)
	GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error)
	UpdateTokens(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Tokens, error)
	RemovePassword(ctx context.Context, in *RemovePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemovePasswordConfirm(ctx context.Context, in *ConfirmRemovePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.Auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ConfirmRegister(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/auth.Auth/ConfirmRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/auth.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetUserByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateTokens(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Tokens, error) {
	out := new(Tokens)
	err := c.cc.Invoke(ctx, "/auth.Auth/UpdateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RemovePassword(ctx context.Context, in *RemovePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/RemovePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RemovePasswordConfirm(ctx context.Context, in *ConfirmRemovePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/RemovePasswordConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetLanguage(ctx context.Context, in *LanguageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/auth.Auth/SetLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Register(context.Context, *RegisterRequest) (*User, error)
	ConfirmRegister(context.Context, *ConfirmRequest) (*Tokens, error)
	Login(context.Context, *LoginRequest) (*Tokens, error)
	GetUserByToken(context.Context, *Token) (*User, error)
	UpdateTokens(context.Context, *Token) (*Tokens, error)
	RemovePassword(context.Context, *RemovePasswordRequest) (*emptypb.Empty, error)
	RemovePasswordConfirm(context.Context, *ConfirmRemovePasswordRequest) (*emptypb.Empty, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*emptypb.Empty, error)
	SetLanguage(context.Context, *LanguageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) ConfirmRegister(context.Context, *ConfirmRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmRegister not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) GetUserByToken(context.Context, *Token) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByToken not implemented")
}
func (UnimplementedAuthServer) UpdateTokens(context.Context, *Token) (*Tokens, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokens not implemented")
}
func (UnimplementedAuthServer) RemovePassword(context.Context, *RemovePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePassword not implemented")
}
func (UnimplementedAuthServer) RemovePasswordConfirm(context.Context, *ConfirmRemovePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePasswordConfirm not implemented")
}
func (UnimplementedAuthServer) ChangePassword(context.Context, *ChangePasswordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthServer) SetLanguage(context.Context, *LanguageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLanguage not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ConfirmRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ConfirmRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ConfirmRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ConfirmRegister(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetUserByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/UpdateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateTokens(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RemovePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RemovePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RemovePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RemovePassword(ctx, req.(*RemovePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RemovePasswordConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRemovePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RemovePasswordConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/RemovePasswordConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RemovePasswordConfirm(ctx, req.(*ConfirmRemovePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SetLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetLanguage(ctx, req.(*LanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "ConfirmRegister",
			Handler:    _Auth_ConfirmRegister_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "GetUserByToken",
			Handler:    _Auth_GetUserByToken_Handler,
		},
		{
			MethodName: "UpdateTokens",
			Handler:    _Auth_UpdateTokens_Handler,
		},
		{
			MethodName: "RemovePassword",
			Handler:    _Auth_RemovePassword_Handler,
		},
		{
			MethodName: "RemovePasswordConfirm",
			Handler:    _Auth_RemovePasswordConfirm_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Auth_ChangePassword_Handler,
		},
		{
			MethodName: "SetLanguage",
			Handler:    _Auth_SetLanguage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
