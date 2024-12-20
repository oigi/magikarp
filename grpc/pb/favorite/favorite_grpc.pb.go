// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc3
// source: favorite.proto

package favorite

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

// FavoriteServiceClient is the client API for FavoriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteServiceClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error)
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
	FavoriteCount(ctx context.Context, in *FavoriteCountReq, opts ...grpc.CallOption) (*FavoriteCountResp, error)
	IsFavorite(ctx context.Context, in *IsFavoriteReq, opts ...grpc.CallOption) (*IsFavoriteResp, error)
}

type favoriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteServiceClient(cc grpc.ClientConnInterface) FavoriteServiceClient {
	return &favoriteServiceClient{cc}
}

func (c *favoriteServiceClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	out := new(FavoriteActionResp)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteService/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteService/FavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) FavoriteCount(ctx context.Context, in *FavoriteCountReq, opts ...grpc.CallOption) (*FavoriteCountResp, error) {
	out := new(FavoriteCountResp)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteService/FavoriteCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteServiceClient) IsFavorite(ctx context.Context, in *IsFavoriteReq, opts ...grpc.CallOption) (*IsFavoriteResp, error) {
	out := new(IsFavoriteResp)
	err := c.cc.Invoke(ctx, "/favorite.FavoriteService/IsFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteServiceServer is the server API for FavoriteService service.
// All implementations must embed UnimplementedFavoriteServiceServer
// for forward compatibility
type FavoriteServiceServer interface {
	FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error)
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	FavoriteCount(context.Context, *FavoriteCountReq) (*FavoriteCountResp, error)
	IsFavorite(context.Context, *IsFavoriteReq) (*IsFavoriteResp, error)
	mustEmbedUnimplementedFavoriteServiceServer()
}

// UnimplementedFavoriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteServiceServer struct {
}

func (UnimplementedFavoriteServiceServer) FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteServiceServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteServiceServer) FavoriteCount(context.Context, *FavoriteCountReq) (*FavoriteCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteCount not implemented")
}
func (UnimplementedFavoriteServiceServer) IsFavorite(context.Context, *IsFavoriteReq) (*IsFavoriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFavorite not implemented")
}
func (UnimplementedFavoriteServiceServer) mustEmbedUnimplementedFavoriteServiceServer() {}

// UnsafeFavoriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteServiceServer will
// result in compilation errors.
type UnsafeFavoriteServiceServer interface {
	mustEmbedUnimplementedFavoriteServiceServer()
}

func RegisterFavoriteServiceServer(s grpc.ServiceRegistrar, srv FavoriteServiceServer) {
	s.RegisterService(&FavoriteService_ServiceDesc, srv)
}

func _FavoriteService_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteService/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteService/FavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_FavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).FavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteService/FavoriteCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).FavoriteCount(ctx, req.(*FavoriteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteService_IsFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFavoriteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.FavoriteService/IsFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteServiceServer).IsFavorite(ctx, req.(*IsFavoriteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteService_ServiceDesc is the grpc.ServiceDesc for FavoriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.FavoriteService",
	HandlerType: (*FavoriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _FavoriteService_FavoriteAction_Handler,
		},
		{
			MethodName: "FavoriteList",
			Handler:    _FavoriteService_FavoriteList_Handler,
		},
		{
			MethodName: "FavoriteCount",
			Handler:    _FavoriteService_FavoriteCount_Handler,
		},
		{
			MethodName: "IsFavorite",
			Handler:    _FavoriteService_IsFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
