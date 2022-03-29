// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: biconom/client/rank/v1/rank_system.proto

package service_client_rank_pb

import (
	context "context"
	rank "github.com/biconom/go-genproto/biconom/type/rank"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RankSystemClient is the client API for RankSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankSystemClient interface {
	Get(ctx context.Context, in *rank.RankSystem_ID, opts ...grpc.CallOption) (*rank.RankSystem, error)
	List(ctx context.Context, in *RankSystemListRequest, opts ...grpc.CallOption) (RankSystem_ListClient, error)
	HeaderGet(ctx context.Context, in *rank.RankSystem_ID, opts ...grpc.CallOption) (*rank.RankSystem_Header, error)
	HeaderList(ctx context.Context, in *RankSystemListRequest, opts ...grpc.CallOption) (RankSystem_HeaderListClient, error)
}

type rankSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewRankSystemClient(cc grpc.ClientConnInterface) RankSystemClient {
	return &rankSystemClient{cc}
}

func (c *rankSystemClient) Get(ctx context.Context, in *rank.RankSystem_ID, opts ...grpc.CallOption) (*rank.RankSystem, error) {
	out := new(rank.RankSystem)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.RankSystem/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankSystemClient) List(ctx context.Context, in *RankSystemListRequest, opts ...grpc.CallOption) (RankSystem_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankSystem_ServiceDesc.Streams[0], "/biconom.client.rank.v1.RankSystem/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankSystemListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RankSystem_ListClient interface {
	Recv() (*rank.RankSystem, error)
	grpc.ClientStream
}

type rankSystemListClient struct {
	grpc.ClientStream
}

func (x *rankSystemListClient) Recv() (*rank.RankSystem, error) {
	m := new(rank.RankSystem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankSystemClient) HeaderGet(ctx context.Context, in *rank.RankSystem_ID, opts ...grpc.CallOption) (*rank.RankSystem_Header, error) {
	out := new(rank.RankSystem_Header)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.RankSystem/HeaderGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankSystemClient) HeaderList(ctx context.Context, in *RankSystemListRequest, opts ...grpc.CallOption) (RankSystem_HeaderListClient, error) {
	stream, err := c.cc.NewStream(ctx, &RankSystem_ServiceDesc.Streams[1], "/biconom.client.rank.v1.RankSystem/HeaderList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankSystemHeaderListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RankSystem_HeaderListClient interface {
	Recv() (*rank.RankSystem_Header, error)
	grpc.ClientStream
}

type rankSystemHeaderListClient struct {
	grpc.ClientStream
}

func (x *rankSystemHeaderListClient) Recv() (*rank.RankSystem_Header, error) {
	m := new(rank.RankSystem_Header)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RankSystemServer is the server API for RankSystem service.
// All implementations must embed UnimplementedRankSystemServer
// for forward compatibility
type RankSystemServer interface {
	Get(context.Context, *rank.RankSystem_ID) (*rank.RankSystem, error)
	List(*RankSystemListRequest, RankSystem_ListServer) error
	HeaderGet(context.Context, *rank.RankSystem_ID) (*rank.RankSystem_Header, error)
	HeaderList(*RankSystemListRequest, RankSystem_HeaderListServer) error
	mustEmbedUnimplementedRankSystemServer()
}

// UnimplementedRankSystemServer must be embedded to have forward compatible implementations.
type UnimplementedRankSystemServer struct {
}

func (UnimplementedRankSystemServer) Get(context.Context, *rank.RankSystem_ID) (*rank.RankSystem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRankSystemServer) List(*RankSystemListRequest, RankSystem_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRankSystemServer) HeaderGet(context.Context, *rank.RankSystem_ID) (*rank.RankSystem_Header, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeaderGet not implemented")
}
func (UnimplementedRankSystemServer) HeaderList(*RankSystemListRequest, RankSystem_HeaderListServer) error {
	return status.Errorf(codes.Unimplemented, "method HeaderList not implemented")
}
func (UnimplementedRankSystemServer) mustEmbedUnimplementedRankSystemServer() {}

// UnsafeRankSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankSystemServer will
// result in compilation errors.
type UnsafeRankSystemServer interface {
	mustEmbedUnimplementedRankSystemServer()
}

func RegisterRankSystemServer(s grpc.ServiceRegistrar, srv RankSystemServer) {
	s.RegisterService(&RankSystem_ServiceDesc, srv)
}

func _RankSystem_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank.RankSystem_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankSystemServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.RankSystem/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankSystemServer).Get(ctx, req.(*rank.RankSystem_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankSystem_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankSystemListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankSystemServer).List(m, &rankSystemListServer{stream})
}

type RankSystem_ListServer interface {
	Send(*rank.RankSystem) error
	grpc.ServerStream
}

type rankSystemListServer struct {
	grpc.ServerStream
}

func (x *rankSystemListServer) Send(m *rank.RankSystem) error {
	return x.ServerStream.SendMsg(m)
}

func _RankSystem_HeaderGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank.RankSystem_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankSystemServer).HeaderGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.RankSystem/HeaderGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankSystemServer).HeaderGet(ctx, req.(*rank.RankSystem_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RankSystem_HeaderList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankSystemListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankSystemServer).HeaderList(m, &rankSystemHeaderListServer{stream})
}

type RankSystem_HeaderListServer interface {
	Send(*rank.RankSystem_Header) error
	grpc.ServerStream
}

type rankSystemHeaderListServer struct {
	grpc.ServerStream
}

func (x *rankSystemHeaderListServer) Send(m *rank.RankSystem_Header) error {
	return x.ServerStream.SendMsg(m)
}

// RankSystem_ServiceDesc is the grpc.ServiceDesc for RankSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RankSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biconom.client.rank.v1.RankSystem",
	HandlerType: (*RankSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RankSystem_Get_Handler,
		},
		{
			MethodName: "HeaderGet",
			Handler:    _RankSystem_HeaderGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _RankSystem_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HeaderList",
			Handler:       _RankSystem_HeaderList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "biconom/client/rank/v1/rank_system.proto",
}