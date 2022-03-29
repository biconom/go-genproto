// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: biconom/client/rank/v1/rank.proto

package service_client_rank_pb

import (
	context "context"
	condition "github.com/biconom/go-genproto/biconom/type/condition"
	rank_system "github.com/biconom/go-genproto/biconom/type/rank_system"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RankClient is the client API for Rank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RankClient interface {
	Get(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Option, error)
	List(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ListClient, error)
	ListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ListByRankSystemClient, error)
	ValueGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Rank, error)
	ValueList(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ValueListClient, error)
	ValueListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ValueListByRankSystemClient, error)
	ValueHeaderGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Rank_Header, error)
	ValueHeaderList(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ValueHeaderListClient, error)
	ValueHeaderListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ValueHeaderListByRankSystemClient, error)
	ConditionGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*condition.Condition, error)
}

type rankClient struct {
	cc grpc.ClientConnInterface
}

func NewRankClient(cc grpc.ClientConnInterface) RankClient {
	return &rankClient{cc}
}

func (c *rankClient) Get(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Option, error) {
	out := new(rank_system.RankSystem_Option)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.Rank/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankClient) List(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[0], "/biconom.client.rank.v1.Rank/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ListClient interface {
	Recv() (*rank_system.RankSystem_Option, error)
	grpc.ClientStream
}

type rankListClient struct {
	grpc.ClientStream
}

func (x *rankListClient) Recv() (*rank_system.RankSystem_Option, error) {
	m := new(rank_system.RankSystem_Option)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ListByRankSystemClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[1], "/biconom.client.rank.v1.Rank/ListByRankSystem", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankListByRankSystemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ListByRankSystemClient interface {
	Recv() (*rank_system.RankSystem_Option, error)
	grpc.ClientStream
}

type rankListByRankSystemClient struct {
	grpc.ClientStream
}

func (x *rankListByRankSystemClient) Recv() (*rank_system.RankSystem_Option, error) {
	m := new(rank_system.RankSystem_Option)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ValueGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Rank, error) {
	out := new(rank_system.RankSystem_Rank)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.Rank/ValueGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankClient) ValueList(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ValueListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[2], "/biconom.client.rank.v1.Rank/ValueList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankValueListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ValueListClient interface {
	Recv() (*rank_system.RankSystem_Rank, error)
	grpc.ClientStream
}

type rankValueListClient struct {
	grpc.ClientStream
}

func (x *rankValueListClient) Recv() (*rank_system.RankSystem_Rank, error) {
	m := new(rank_system.RankSystem_Rank)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ValueListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ValueListByRankSystemClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[3], "/biconom.client.rank.v1.Rank/ValueListByRankSystem", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankValueListByRankSystemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ValueListByRankSystemClient interface {
	Recv() (*rank_system.RankSystem_Rank, error)
	grpc.ClientStream
}

type rankValueListByRankSystemClient struct {
	grpc.ClientStream
}

func (x *rankValueListByRankSystemClient) Recv() (*rank_system.RankSystem_Rank, error) {
	m := new(rank_system.RankSystem_Rank)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ValueHeaderGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*rank_system.RankSystem_Rank_Header, error) {
	out := new(rank_system.RankSystem_Rank_Header)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.Rank/ValueHeaderGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rankClient) ValueHeaderList(ctx context.Context, in *RankListRequest, opts ...grpc.CallOption) (Rank_ValueHeaderListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[4], "/biconom.client.rank.v1.Rank/ValueHeaderList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankValueHeaderListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ValueHeaderListClient interface {
	Recv() (*rank_system.RankSystem_Rank_Header, error)
	grpc.ClientStream
}

type rankValueHeaderListClient struct {
	grpc.ClientStream
}

func (x *rankValueHeaderListClient) Recv() (*rank_system.RankSystem_Rank_Header, error) {
	m := new(rank_system.RankSystem_Rank_Header)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ValueHeaderListByRankSystem(ctx context.Context, in *RankListByRankSystemRequest, opts ...grpc.CallOption) (Rank_ValueHeaderListByRankSystemClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rank_ServiceDesc.Streams[5], "/biconom.client.rank.v1.Rank/ValueHeaderListByRankSystem", opts...)
	if err != nil {
		return nil, err
	}
	x := &rankValueHeaderListByRankSystemClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rank_ValueHeaderListByRankSystemClient interface {
	Recv() (*rank_system.RankSystem_Rank_Header, error)
	grpc.ClientStream
}

type rankValueHeaderListByRankSystemClient struct {
	grpc.ClientStream
}

func (x *rankValueHeaderListByRankSystemClient) Recv() (*rank_system.RankSystem_Rank_Header, error) {
	m := new(rank_system.RankSystem_Rank_Header)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rankClient) ConditionGet(ctx context.Context, in *rank_system.RankSystem_Rank_ID, opts ...grpc.CallOption) (*condition.Condition, error) {
	out := new(condition.Condition)
	err := c.cc.Invoke(ctx, "/biconom.client.rank.v1.Rank/ConditionGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServer is the server API for Rank service.
// All implementations must embed UnimplementedRankServer
// for forward compatibility
type RankServer interface {
	Get(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Option, error)
	List(*RankListRequest, Rank_ListServer) error
	ListByRankSystem(*RankListByRankSystemRequest, Rank_ListByRankSystemServer) error
	ValueGet(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Rank, error)
	ValueList(*RankListRequest, Rank_ValueListServer) error
	ValueListByRankSystem(*RankListByRankSystemRequest, Rank_ValueListByRankSystemServer) error
	ValueHeaderGet(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Rank_Header, error)
	ValueHeaderList(*RankListRequest, Rank_ValueHeaderListServer) error
	ValueHeaderListByRankSystem(*RankListByRankSystemRequest, Rank_ValueHeaderListByRankSystemServer) error
	ConditionGet(context.Context, *rank_system.RankSystem_Rank_ID) (*condition.Condition, error)
	mustEmbedUnimplementedRankServer()
}

// UnimplementedRankServer must be embedded to have forward compatible implementations.
type UnimplementedRankServer struct {
}

func (UnimplementedRankServer) Get(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Option, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRankServer) List(*RankListRequest, Rank_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRankServer) ListByRankSystem(*RankListByRankSystemRequest, Rank_ListByRankSystemServer) error {
	return status.Errorf(codes.Unimplemented, "method ListByRankSystem not implemented")
}
func (UnimplementedRankServer) ValueGet(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Rank, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValueGet not implemented")
}
func (UnimplementedRankServer) ValueList(*RankListRequest, Rank_ValueListServer) error {
	return status.Errorf(codes.Unimplemented, "method ValueList not implemented")
}
func (UnimplementedRankServer) ValueListByRankSystem(*RankListByRankSystemRequest, Rank_ValueListByRankSystemServer) error {
	return status.Errorf(codes.Unimplemented, "method ValueListByRankSystem not implemented")
}
func (UnimplementedRankServer) ValueHeaderGet(context.Context, *rank_system.RankSystem_Rank_ID) (*rank_system.RankSystem_Rank_Header, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValueHeaderGet not implemented")
}
func (UnimplementedRankServer) ValueHeaderList(*RankListRequest, Rank_ValueHeaderListServer) error {
	return status.Errorf(codes.Unimplemented, "method ValueHeaderList not implemented")
}
func (UnimplementedRankServer) ValueHeaderListByRankSystem(*RankListByRankSystemRequest, Rank_ValueHeaderListByRankSystemServer) error {
	return status.Errorf(codes.Unimplemented, "method ValueHeaderListByRankSystem not implemented")
}
func (UnimplementedRankServer) ConditionGet(context.Context, *rank_system.RankSystem_Rank_ID) (*condition.Condition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConditionGet not implemented")
}
func (UnimplementedRankServer) mustEmbedUnimplementedRankServer() {}

// UnsafeRankServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RankServer will
// result in compilation errors.
type UnsafeRankServer interface {
	mustEmbedUnimplementedRankServer()
}

func RegisterRankServer(s grpc.ServiceRegistrar, srv RankServer) {
	s.RegisterService(&Rank_ServiceDesc, srv)
}

func _Rank_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank_system.RankSystem_Rank_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.Rank/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).Get(ctx, req.(*rank_system.RankSystem_Rank_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rank_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).List(m, &rankListServer{stream})
}

type Rank_ListServer interface {
	Send(*rank_system.RankSystem_Option) error
	grpc.ServerStream
}

type rankListServer struct {
	grpc.ServerStream
}

func (x *rankListServer) Send(m *rank_system.RankSystem_Option) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ListByRankSystem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListByRankSystemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).ListByRankSystem(m, &rankListByRankSystemServer{stream})
}

type Rank_ListByRankSystemServer interface {
	Send(*rank_system.RankSystem_Option) error
	grpc.ServerStream
}

type rankListByRankSystemServer struct {
	grpc.ServerStream
}

func (x *rankListByRankSystemServer) Send(m *rank_system.RankSystem_Option) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ValueGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank_system.RankSystem_Rank_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).ValueGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.Rank/ValueGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).ValueGet(ctx, req.(*rank_system.RankSystem_Rank_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rank_ValueList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).ValueList(m, &rankValueListServer{stream})
}

type Rank_ValueListServer interface {
	Send(*rank_system.RankSystem_Rank) error
	grpc.ServerStream
}

type rankValueListServer struct {
	grpc.ServerStream
}

func (x *rankValueListServer) Send(m *rank_system.RankSystem_Rank) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ValueListByRankSystem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListByRankSystemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).ValueListByRankSystem(m, &rankValueListByRankSystemServer{stream})
}

type Rank_ValueListByRankSystemServer interface {
	Send(*rank_system.RankSystem_Rank) error
	grpc.ServerStream
}

type rankValueListByRankSystemServer struct {
	grpc.ServerStream
}

func (x *rankValueListByRankSystemServer) Send(m *rank_system.RankSystem_Rank) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ValueHeaderGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank_system.RankSystem_Rank_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).ValueHeaderGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.Rank/ValueHeaderGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).ValueHeaderGet(ctx, req.(*rank_system.RankSystem_Rank_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rank_ValueHeaderList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).ValueHeaderList(m, &rankValueHeaderListServer{stream})
}

type Rank_ValueHeaderListServer interface {
	Send(*rank_system.RankSystem_Rank_Header) error
	grpc.ServerStream
}

type rankValueHeaderListServer struct {
	grpc.ServerStream
}

func (x *rankValueHeaderListServer) Send(m *rank_system.RankSystem_Rank_Header) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ValueHeaderListByRankSystem_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RankListByRankSystemRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RankServer).ValueHeaderListByRankSystem(m, &rankValueHeaderListByRankSystemServer{stream})
}

type Rank_ValueHeaderListByRankSystemServer interface {
	Send(*rank_system.RankSystem_Rank_Header) error
	grpc.ServerStream
}

type rankValueHeaderListByRankSystemServer struct {
	grpc.ServerStream
}

func (x *rankValueHeaderListByRankSystemServer) Send(m *rank_system.RankSystem_Rank_Header) error {
	return x.ServerStream.SendMsg(m)
}

func _Rank_ConditionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(rank_system.RankSystem_Rank_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServer).ConditionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.client.rank.v1.Rank/ConditionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServer).ConditionGet(ctx, req.(*rank_system.RankSystem_Rank_ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Rank_ServiceDesc is the grpc.ServiceDesc for Rank service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rank_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biconom.client.rank.v1.Rank",
	HandlerType: (*RankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Rank_Get_Handler,
		},
		{
			MethodName: "ValueGet",
			Handler:    _Rank_ValueGet_Handler,
		},
		{
			MethodName: "ValueHeaderGet",
			Handler:    _Rank_ValueHeaderGet_Handler,
		},
		{
			MethodName: "ConditionGet",
			Handler:    _Rank_ConditionGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Rank_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListByRankSystem",
			Handler:       _Rank_ListByRankSystem_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ValueList",
			Handler:       _Rank_ValueList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ValueListByRankSystem",
			Handler:       _Rank_ValueListByRankSystem_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ValueHeaderList",
			Handler:       _Rank_ValueHeaderList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ValueHeaderListByRankSystem",
			Handler:       _Rank_ValueHeaderListByRankSystem_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "biconom/client/rank/v1/rank.proto",
}
