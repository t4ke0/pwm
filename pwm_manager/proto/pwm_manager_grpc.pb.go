// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	StorePassword(ctx context.Context, in *ManagerRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdatePassword(ctx context.Context, in *ManagerUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	GetPasswords(ctx context.Context, in *GetPasswordsRequest, opts ...grpc.CallOption) (*UserPasswords, error)
	GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratedPassword, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) StorePassword(ctx context.Context, in *ManagerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Manager/StorePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdatePassword(ctx context.Context, in *ManagerUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Manager/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetPasswords(ctx context.Context, in *GetPasswordsRequest, opts ...grpc.CallOption) (*UserPasswords, error) {
	out := new(UserPasswords)
	err := c.cc.Invoke(ctx, "/proto.Manager/GetPasswords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GeneratePassword(ctx context.Context, in *GeneratePasswordRequest, opts ...grpc.CallOption) (*GeneratedPassword, error) {
	out := new(GeneratedPassword)
	err := c.cc.Invoke(ctx, "/proto.Manager/GeneratePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	StorePassword(context.Context, *ManagerRequest) (*Empty, error)
	UpdatePassword(context.Context, *ManagerUpdateRequest) (*Empty, error)
	GetPasswords(context.Context, *GetPasswordsRequest) (*UserPasswords, error)
	GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratedPassword, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) StorePassword(context.Context, *ManagerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePassword not implemented")
}
func (UnimplementedManagerServer) UpdatePassword(context.Context, *ManagerUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedManagerServer) GetPasswords(context.Context, *GetPasswordsRequest) (*UserPasswords, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPasswords not implemented")
}
func (UnimplementedManagerServer) GeneratePassword(context.Context, *GeneratePasswordRequest) (*GeneratedPassword, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePassword not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_StorePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).StorePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/StorePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).StorePassword(ctx, req.(*ManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdatePassword(ctx, req.(*ManagerUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPasswordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/GetPasswords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetPasswords(ctx, req.(*GetPasswordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GeneratePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneratePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GeneratePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Manager/GeneratePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GeneratePassword(ctx, req.(*GeneratePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StorePassword",
			Handler:    _Manager_StorePassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _Manager_UpdatePassword_Handler,
		},
		{
			MethodName: "GetPasswords",
			Handler:    _Manager_GetPasswords_Handler,
		},
		{
			MethodName: "GeneratePassword",
			Handler:    _Manager_GeneratePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pwm_manager.proto",
}
