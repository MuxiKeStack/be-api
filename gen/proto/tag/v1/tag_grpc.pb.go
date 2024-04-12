// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/tag/v1/tag.proto

package tagv1

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
	TagService_AttachAssessmentTags_FullMethodName = "/tag.v1.TagService/AttachAssessmentTags"
	TagService_AttachFeaturesTags_FullMethodName   = "/tag.v1.TagService/AttachFeaturesTags"
)

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagServiceClient interface {
	AttachAssessmentTags(ctx context.Context, in *AttachAssessmentTagsRequest, opts ...grpc.CallOption) (*AttachAssessmentTagsResponse, error)
	AttachFeaturesTags(ctx context.Context, in *AttachFeaturesTagsRequest, opts ...grpc.CallOption) (*AttachFeaturesTagsResponse, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) AttachAssessmentTags(ctx context.Context, in *AttachAssessmentTagsRequest, opts ...grpc.CallOption) (*AttachAssessmentTagsResponse, error) {
	out := new(AttachAssessmentTagsResponse)
	err := c.cc.Invoke(ctx, TagService_AttachAssessmentTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) AttachFeaturesTags(ctx context.Context, in *AttachFeaturesTagsRequest, opts ...grpc.CallOption) (*AttachFeaturesTagsResponse, error) {
	out := new(AttachFeaturesTagsResponse)
	err := c.cc.Invoke(ctx, TagService_AttachFeaturesTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
// All implementations must embed UnimplementedTagServiceServer
// for forward compatibility
type TagServiceServer interface {
	AttachAssessmentTags(context.Context, *AttachAssessmentTagsRequest) (*AttachAssessmentTagsResponse, error)
	AttachFeaturesTags(context.Context, *AttachFeaturesTagsRequest) (*AttachFeaturesTagsResponse, error)
	mustEmbedUnimplementedTagServiceServer()
}

// UnimplementedTagServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (UnimplementedTagServiceServer) AttachAssessmentTags(context.Context, *AttachAssessmentTagsRequest) (*AttachAssessmentTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachAssessmentTags not implemented")
}
func (UnimplementedTagServiceServer) AttachFeaturesTags(context.Context, *AttachFeaturesTagsRequest) (*AttachFeaturesTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachFeaturesTags not implemented")
}
func (UnimplementedTagServiceServer) mustEmbedUnimplementedTagServiceServer() {}

// UnsafeTagServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagServiceServer will
// result in compilation errors.
type UnsafeTagServiceServer interface {
	mustEmbedUnimplementedTagServiceServer()
}

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	s.RegisterService(&TagService_ServiceDesc, srv)
}

func _TagService_AttachAssessmentTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachAssessmentTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AttachAssessmentTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_AttachAssessmentTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AttachAssessmentTags(ctx, req.(*AttachAssessmentTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_AttachFeaturesTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachFeaturesTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AttachFeaturesTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TagService_AttachFeaturesTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AttachFeaturesTags(ctx, req.(*AttachFeaturesTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TagService_ServiceDesc is the grpc.ServiceDesc for TagService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TagService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tag.v1.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttachAssessmentTags",
			Handler:    _TagService_AttachAssessmentTags_Handler,
		},
		{
			MethodName: "AttachFeaturesTags",
			Handler:    _TagService_AttachFeaturesTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tag/v1/tag.proto",
}