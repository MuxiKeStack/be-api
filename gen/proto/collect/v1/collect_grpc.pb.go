// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/collect/v1/collect.proto

package collectv1

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
	CollectService_AddCollection_FullMethodName    = "/collect.v1.CollectService/AddCollection"
	CollectService_RemoveCollection_FullMethodName = "/collect.v1.CollectService/RemoveCollection"
	CollectService_ListCollections_FullMethodName  = "/collect.v1.CollectService/ListCollections"
	CollectService_CountCollections_FullMethodName = "/collect.v1.CollectService/CountCollections"
	CollectService_CheckCollection_FullMethodName  = "/collect.v1.CollectService/CheckCollection"
)

// CollectServiceClient is the client API for CollectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectServiceClient interface {
	AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error)
	RemoveCollection(ctx context.Context, in *RemoveCollectionRequest, opts ...grpc.CallOption) (*RemoveCollectionResponse, error)
	ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error)
	CountCollections(ctx context.Context, in *CountCollectionsRequest, opts ...grpc.CallOption) (*CountCollectionsResponse, error)
	CheckCollection(ctx context.Context, in *CheckCollectionRequest, opts ...grpc.CallOption) (*CheckCollectionResponse, error)
}

type collectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectServiceClient(cc grpc.ClientConnInterface) CollectServiceClient {
	return &collectServiceClient{cc}
}

func (c *collectServiceClient) AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*AddCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCollectionResponse)
	err := c.cc.Invoke(ctx, CollectService_AddCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) RemoveCollection(ctx context.Context, in *RemoveCollectionRequest, opts ...grpc.CallOption) (*RemoveCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveCollectionResponse)
	err := c.cc.Invoke(ctx, CollectService_RemoveCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCollectionsResponse)
	err := c.cc.Invoke(ctx, CollectService_ListCollections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CountCollections(ctx context.Context, in *CountCollectionsRequest, opts ...grpc.CallOption) (*CountCollectionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountCollectionsResponse)
	err := c.cc.Invoke(ctx, CollectService_CountCollections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectServiceClient) CheckCollection(ctx context.Context, in *CheckCollectionRequest, opts ...grpc.CallOption) (*CheckCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckCollectionResponse)
	err := c.cc.Invoke(ctx, CollectService_CheckCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectServiceServer is the server API for CollectService service.
// All implementations must embed UnimplementedCollectServiceServer
// for forward compatibility.
type CollectServiceServer interface {
	AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error)
	RemoveCollection(context.Context, *RemoveCollectionRequest) (*RemoveCollectionResponse, error)
	ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error)
	CountCollections(context.Context, *CountCollectionsRequest) (*CountCollectionsResponse, error)
	CheckCollection(context.Context, *CheckCollectionRequest) (*CheckCollectionResponse, error)
	mustEmbedUnimplementedCollectServiceServer()
}

// UnimplementedCollectServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCollectServiceServer struct{}

func (UnimplementedCollectServiceServer) AddCollection(context.Context, *AddCollectionRequest) (*AddCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollection not implemented")
}
func (UnimplementedCollectServiceServer) RemoveCollection(context.Context, *RemoveCollectionRequest) (*RemoveCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCollection not implemented")
}
func (UnimplementedCollectServiceServer) ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollections not implemented")
}
func (UnimplementedCollectServiceServer) CountCollections(context.Context, *CountCollectionsRequest) (*CountCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountCollections not implemented")
}
func (UnimplementedCollectServiceServer) CheckCollection(context.Context, *CheckCollectionRequest) (*CheckCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCollection not implemented")
}
func (UnimplementedCollectServiceServer) mustEmbedUnimplementedCollectServiceServer() {}
func (UnimplementedCollectServiceServer) testEmbeddedByValue()                        {}

// UnsafeCollectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectServiceServer will
// result in compilation errors.
type UnsafeCollectServiceServer interface {
	mustEmbedUnimplementedCollectServiceServer()
}

func RegisterCollectServiceServer(s grpc.ServiceRegistrar, srv CollectServiceServer) {
	// If the following call pancis, it indicates UnimplementedCollectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CollectService_ServiceDesc, srv)
}

func _CollectService_AddCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).AddCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_AddCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).AddCollection(ctx, req.(*AddCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_RemoveCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).RemoveCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_RemoveCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).RemoveCollection(ctx, req.(*RemoveCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_ListCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).ListCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_ListCollections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).ListCollections(ctx, req.(*ListCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CountCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CountCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_CountCollections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CountCollections(ctx, req.(*CountCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectService_CheckCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectServiceServer).CheckCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectService_CheckCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectServiceServer).CheckCollection(ctx, req.(*CheckCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectService_ServiceDesc is the grpc.ServiceDesc for CollectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "collect.v1.CollectService",
	HandlerType: (*CollectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCollection",
			Handler:    _CollectService_AddCollection_Handler,
		},
		{
			MethodName: "RemoveCollection",
			Handler:    _CollectService_RemoveCollection_Handler,
		},
		{
			MethodName: "ListCollections",
			Handler:    _CollectService_ListCollections_Handler,
		},
		{
			MethodName: "CountCollections",
			Handler:    _CollectService_CountCollections_Handler,
		},
		{
			MethodName: "CheckCollection",
			Handler:    _CollectService_CheckCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/collect/v1/collect.proto",
}
