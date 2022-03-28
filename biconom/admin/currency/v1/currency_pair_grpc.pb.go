// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: biconom/admin/currency/v1/currency_pair.proto

package service_admin_currency_pb

import (
	context "context"
	currency "github.com/biconom/go-genproto/biconom/type/currency"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CurrencyPairClient is the client API for CurrencyPair service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyPairClient interface {
	Get(ctx context.Context, in *currency.Currency_Pair_ID, opts ...grpc.CallOption) (*currency.Currency_Pair, error)
	List(ctx context.Context, in *CurrencyPairListRequest, opts ...grpc.CallOption) (CurrencyPair_ListClient, error)
	ListBySource(ctx context.Context, in *CurrencyPairListBySourceRequest, opts ...grpc.CallOption) (CurrencyPair_ListBySourceClient, error)
	ListByTarget(ctx context.Context, in *CurrencyPairListByTargetRequest, opts ...grpc.CallOption) (CurrencyPair_ListByTargetClient, error)
	ListByCurrency(ctx context.Context, in *CurrencyPairListByCurrencyRequest, opts ...grpc.CallOption) (CurrencyPair_ListByCurrencyClient, error)
	Create(ctx context.Context, in *CurrencyPairCreateRequest, opts ...grpc.CallOption) (*currency.Currency_Pair, error)
}

type currencyPairClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyPairClient(cc grpc.ClientConnInterface) CurrencyPairClient {
	return &currencyPairClient{cc}
}

func (c *currencyPairClient) Get(ctx context.Context, in *currency.Currency_Pair_ID, opts ...grpc.CallOption) (*currency.Currency_Pair, error) {
	out := new(currency.Currency_Pair)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.CurrencyPair/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyPairClient) List(ctx context.Context, in *CurrencyPairListRequest, opts ...grpc.CallOption) (CurrencyPair_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyPair_ServiceDesc.Streams[0], "/biconom.admin.currency.v1.CurrencyPair/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyPairListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyPair_ListClient interface {
	Recv() (*currency.Currency_Pair, error)
	grpc.ClientStream
}

type currencyPairListClient struct {
	grpc.ClientStream
}

func (x *currencyPairListClient) Recv() (*currency.Currency_Pair, error) {
	m := new(currency.Currency_Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyPairClient) ListBySource(ctx context.Context, in *CurrencyPairListBySourceRequest, opts ...grpc.CallOption) (CurrencyPair_ListBySourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyPair_ServiceDesc.Streams[1], "/biconom.admin.currency.v1.CurrencyPair/ListBySource", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyPairListBySourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyPair_ListBySourceClient interface {
	Recv() (*currency.Currency_Pair, error)
	grpc.ClientStream
}

type currencyPairListBySourceClient struct {
	grpc.ClientStream
}

func (x *currencyPairListBySourceClient) Recv() (*currency.Currency_Pair, error) {
	m := new(currency.Currency_Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyPairClient) ListByTarget(ctx context.Context, in *CurrencyPairListByTargetRequest, opts ...grpc.CallOption) (CurrencyPair_ListByTargetClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyPair_ServiceDesc.Streams[2], "/biconom.admin.currency.v1.CurrencyPair/ListByTarget", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyPairListByTargetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyPair_ListByTargetClient interface {
	Recv() (*currency.Currency_Pair, error)
	grpc.ClientStream
}

type currencyPairListByTargetClient struct {
	grpc.ClientStream
}

func (x *currencyPairListByTargetClient) Recv() (*currency.Currency_Pair, error) {
	m := new(currency.Currency_Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyPairClient) ListByCurrency(ctx context.Context, in *CurrencyPairListByCurrencyRequest, opts ...grpc.CallOption) (CurrencyPair_ListByCurrencyClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyPair_ServiceDesc.Streams[3], "/biconom.admin.currency.v1.CurrencyPair/ListByCurrency", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyPairListByCurrencyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyPair_ListByCurrencyClient interface {
	Recv() (*currency.Currency_Pair, error)
	grpc.ClientStream
}

type currencyPairListByCurrencyClient struct {
	grpc.ClientStream
}

func (x *currencyPairListByCurrencyClient) Recv() (*currency.Currency_Pair, error) {
	m := new(currency.Currency_Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyPairClient) Create(ctx context.Context, in *CurrencyPairCreateRequest, opts ...grpc.CallOption) (*currency.Currency_Pair, error) {
	out := new(currency.Currency_Pair)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.CurrencyPair/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyPairServer is the server API for CurrencyPair service.
// All implementations must embed UnimplementedCurrencyPairServer
// for forward compatibility
type CurrencyPairServer interface {
	Get(context.Context, *currency.Currency_Pair_ID) (*currency.Currency_Pair, error)
	List(*CurrencyPairListRequest, CurrencyPair_ListServer) error
	ListBySource(*CurrencyPairListBySourceRequest, CurrencyPair_ListBySourceServer) error
	ListByTarget(*CurrencyPairListByTargetRequest, CurrencyPair_ListByTargetServer) error
	ListByCurrency(*CurrencyPairListByCurrencyRequest, CurrencyPair_ListByCurrencyServer) error
	Create(context.Context, *CurrencyPairCreateRequest) (*currency.Currency_Pair, error)
	mustEmbedUnimplementedCurrencyPairServer()
}

// UnimplementedCurrencyPairServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyPairServer struct {
}

func (UnimplementedCurrencyPairServer) Get(context.Context, *currency.Currency_Pair_ID) (*currency.Currency_Pair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCurrencyPairServer) List(*CurrencyPairListRequest, CurrencyPair_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCurrencyPairServer) ListBySource(*CurrencyPairListBySourceRequest, CurrencyPair_ListBySourceServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBySource not implemented")
}
func (UnimplementedCurrencyPairServer) ListByTarget(*CurrencyPairListByTargetRequest, CurrencyPair_ListByTargetServer) error {
	return status.Errorf(codes.Unimplemented, "method ListByTarget not implemented")
}
func (UnimplementedCurrencyPairServer) ListByCurrency(*CurrencyPairListByCurrencyRequest, CurrencyPair_ListByCurrencyServer) error {
	return status.Errorf(codes.Unimplemented, "method ListByCurrency not implemented")
}
func (UnimplementedCurrencyPairServer) Create(context.Context, *CurrencyPairCreateRequest) (*currency.Currency_Pair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCurrencyPairServer) mustEmbedUnimplementedCurrencyPairServer() {}

// UnsafeCurrencyPairServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyPairServer will
// result in compilation errors.
type UnsafeCurrencyPairServer interface {
	mustEmbedUnimplementedCurrencyPairServer()
}

func RegisterCurrencyPairServer(s grpc.ServiceRegistrar, srv CurrencyPairServer) {
	s.RegisterService(&CurrencyPair_ServiceDesc, srv)
}

func _CurrencyPair_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(currency.Currency_Pair_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyPairServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.CurrencyPair/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyPairServer).Get(ctx, req.(*currency.Currency_Pair_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyPair_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyPairListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyPairServer).List(m, &currencyPairListServer{stream})
}

type CurrencyPair_ListServer interface {
	Send(*currency.Currency_Pair) error
	grpc.ServerStream
}

type currencyPairListServer struct {
	grpc.ServerStream
}

func (x *currencyPairListServer) Send(m *currency.Currency_Pair) error {
	return x.ServerStream.SendMsg(m)
}

func _CurrencyPair_ListBySource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyPairListBySourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyPairServer).ListBySource(m, &currencyPairListBySourceServer{stream})
}

type CurrencyPair_ListBySourceServer interface {
	Send(*currency.Currency_Pair) error
	grpc.ServerStream
}

type currencyPairListBySourceServer struct {
	grpc.ServerStream
}

func (x *currencyPairListBySourceServer) Send(m *currency.Currency_Pair) error {
	return x.ServerStream.SendMsg(m)
}

func _CurrencyPair_ListByTarget_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyPairListByTargetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyPairServer).ListByTarget(m, &currencyPairListByTargetServer{stream})
}

type CurrencyPair_ListByTargetServer interface {
	Send(*currency.Currency_Pair) error
	grpc.ServerStream
}

type currencyPairListByTargetServer struct {
	grpc.ServerStream
}

func (x *currencyPairListByTargetServer) Send(m *currency.Currency_Pair) error {
	return x.ServerStream.SendMsg(m)
}

func _CurrencyPair_ListByCurrency_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyPairListByCurrencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyPairServer).ListByCurrency(m, &currencyPairListByCurrencyServer{stream})
}

type CurrencyPair_ListByCurrencyServer interface {
	Send(*currency.Currency_Pair) error
	grpc.ServerStream
}

type currencyPairListByCurrencyServer struct {
	grpc.ServerStream
}

func (x *currencyPairListByCurrencyServer) Send(m *currency.Currency_Pair) error {
	return x.ServerStream.SendMsg(m)
}

func _CurrencyPair_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyPairCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyPairServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.CurrencyPair/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyPairServer).Create(ctx, req.(*CurrencyPairCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyPair_ServiceDesc is the grpc.ServiceDesc for CurrencyPair service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyPair_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biconom.admin.currency.v1.CurrencyPair",
	HandlerType: (*CurrencyPairServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CurrencyPair_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CurrencyPair_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _CurrencyPair_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListBySource",
			Handler:       _CurrencyPair_ListBySource_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListByTarget",
			Handler:       _CurrencyPair_ListByTarget_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListByCurrency",
			Handler:       _CurrencyPair_ListByCurrency_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "biconom/admin/currency/v1/currency_pair.proto",
}
