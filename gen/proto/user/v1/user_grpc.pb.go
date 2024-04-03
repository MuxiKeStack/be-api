// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/user/v1/user.proto

package userv1

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

const (
	UserService_LoginByCCNU_FullMethodName            = "/user.v1.UserService/LoginByCCNU"
	UserService_UpdateNonSensitiveInfo_FullMethodName = "/user.v1.UserService/UpdateNonSensitiveInfo"
	UserService_Profile_FullMethodName                = "/user.v1.UserService/Profile"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	LoginByCCNU(ctx context.Context, in *LoginByCCNURequest, opts ...grpc.CallOption) (*LoginByCCNUResponse, error)
	UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) LoginByCCNU(ctx context.Context, in *LoginByCCNURequest, opts ...grpc.CallOption) (*LoginByCCNUResponse, error) {
	out := new(LoginByCCNUResponse)
	err := c.cc.Invoke(ctx, UserService_LoginByCCNU_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateNonSensitiveInfo(ctx context.Context, in *UpdateNonSensitiveInfoRequest, opts ...grpc.CallOption) (*UpdateNonSensitiveInfoResponse, error) {
	out := new(UpdateNonSensitiveInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateNonSensitiveInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, UserService_Profile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	LoginByCCNU(context.Context, *LoginByCCNURequest) (*LoginByCCNUResponse, error)
	UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error)
	Profile(context.Context, *ProfileRequest) (*ProfileResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) LoginByCCNU(context.Context, *LoginByCCNURequest) (*LoginByCCNUResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByCCNU not implemented")
}
func (UnimplementedUserServiceServer) UpdateNonSensitiveInfo(context.Context, *UpdateNonSensitiveInfoRequest) (*UpdateNonSensitiveInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNonSensitiveInfo not implemented")
}
func (UnimplementedUserServiceServer) Profile(context.Context, *ProfileRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_LoginByCCNU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByCCNURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginByCCNU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginByCCNU_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginByCCNU(ctx, req.(*LoginByCCNURequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateNonSensitiveInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNonSensitiveInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateNonSensitiveInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateNonSensitiveInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateNonSensitiveInfo(ctx, req.(*UpdateNonSensitiveInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Profile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*ProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByCCNU",
			Handler:    _UserService_LoginByCCNU_Handler,
		},
		{
			MethodName: "UpdateNonSensitiveInfo",
			Handler:    _UserService_UpdateNonSensitiveInfo_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/v1/user.proto",
}