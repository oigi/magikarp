// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: publish.proto

package publish

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

// PublishServiceClient is the client API for PublishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishServiceClient interface {
	CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error)
	ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error)
	CountVideo(ctx context.Context, in *CountVideoRequest, opts ...grpc.CallOption) (*CountVideoResponse, error)
}

type publishServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishServiceClient(cc grpc.ClientConnInterface) PublishServiceClient {
	return &publishServiceClient{cc}
}

func (c *publishServiceClient) CreateVideo(ctx context.Context, in *CreateVideoRequest, opts ...grpc.CallOption) (*CreateVideoResponse, error) {
	out := new(CreateVideoResponse)
	err := c.cc.Invoke(ctx, "/publish.PublishService/CreateVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error) {
	out := new(ListVideoResponse)
	err := c.cc.Invoke(ctx, "/publish.PublishService/ListVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishServiceClient) CountVideo(ctx context.Context, in *CountVideoRequest, opts ...grpc.CallOption) (*CountVideoResponse, error) {
	out := new(CountVideoResponse)
	err := c.cc.Invoke(ctx, "/publish.PublishService/CountVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServiceServer is the server API for PublishService service.
// All implementations must embed UnimplementedPublishServiceServer
// for forward compatibility
type PublishServiceServer interface {
	CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error)
	ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error)
	CountVideo(context.Context, *CountVideoRequest) (*CountVideoResponse, error)
	mustEmbedUnimplementedPublishServiceServer()
}

// UnimplementedPublishServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServiceServer struct {
}

func (UnimplementedPublishServiceServer) CreateVideo(context.Context, *CreateVideoRequest) (*CreateVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVideo not implemented")
}
func (UnimplementedPublishServiceServer) ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideo not implemented")
}
func (UnimplementedPublishServiceServer) CountVideo(context.Context, *CountVideoRequest) (*CountVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountVideo not implemented")
}
func (UnimplementedPublishServiceServer) mustEmbedUnimplementedPublishServiceServer() {}

// UnsafePublishServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServiceServer will
// result in compilation errors.
type UnsafePublishServiceServer interface {
	mustEmbedUnimplementedPublishServiceServer()
}

func RegisterPublishServiceServer(s grpc.ServiceRegistrar, srv PublishServiceServer) {
	s.RegisterService(&PublishService_ServiceDesc, srv)
}

func _PublishService_CreateVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).CreateVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.PublishService/CreateVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).CreateVideo(ctx, req.(*CreateVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_ListVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).ListVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.PublishService/ListVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).ListVideo(ctx, req.(*ListVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublishService_CountVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).CountVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publish.PublishService/CountVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).CountVideo(ctx, req.(*CountVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublishService_ServiceDesc is the grpc.ServiceDesc for PublishService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublishService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publish.PublishService",
	HandlerType: (*PublishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVideo",
			Handler:    _PublishService_CreateVideo_Handler,
		},
		{
			MethodName: "ListVideo",
			Handler:    _PublishService_ListVideo_Handler,
		},
		{
			MethodName: "CountVideo",
			Handler:    _PublishService_CountVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publish.proto",
}
