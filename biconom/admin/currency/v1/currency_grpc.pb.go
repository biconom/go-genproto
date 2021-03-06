// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: biconom/admin/currency/v1/currency.proto

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

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyClient interface {
	Get(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error)
	List(ctx context.Context, in *CurrencyListRequest, opts ...grpc.CallOption) (Currency_ListClient, error)
	Create(ctx context.Context, in *CurrencyCreateRequest, opts ...grpc.CallOption) (*currency.Currency, error)
	Publish(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error)
	Unpublish(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) Get(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error) {
	out := new(currency.Currency)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.Currency/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) List(ctx context.Context, in *CurrencyListRequest, opts ...grpc.CallOption) (Currency_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &Currency_ServiceDesc.Streams[0], "/biconom.admin.currency.v1.Currency/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Currency_ListClient interface {
	Recv() (*currency.Currency, error)
	grpc.ClientStream
}

type currencyListClient struct {
	grpc.ClientStream
}

func (x *currencyListClient) Recv() (*currency.Currency, error) {
	m := new(currency.Currency)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *currencyClient) Create(ctx context.Context, in *CurrencyCreateRequest, opts ...grpc.CallOption) (*currency.Currency, error) {
	out := new(currency.Currency)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.Currency/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) Publish(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error) {
	out := new(currency.Currency)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.Currency/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyClient) Unpublish(ctx context.Context, in *currency.Currency_ID, opts ...grpc.CallOption) (*currency.Currency, error) {
	out := new(currency.Currency)
	err := c.cc.Invoke(ctx, "/biconom.admin.currency.v1.Currency/Unpublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
// All implementations must embed UnimplementedCurrencyServer
// for forward compatibility
type CurrencyServer interface {
	Get(context.Context, *currency.Currency_ID) (*currency.Currency, error)
	List(*CurrencyListRequest, Currency_ListServer) error
	Create(context.Context, *CurrencyCreateRequest) (*currency.Currency, error)
	Publish(context.Context, *currency.Currency_ID) (*currency.Currency, error)
	Unpublish(context.Context, *currency.Currency_ID) (*currency.Currency, error)
	mustEmbedUnimplementedCurrencyServer()
}

// UnimplementedCurrencyServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (UnimplementedCurrencyServer) Get(context.Context, *currency.Currency_ID) (*currency.Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCurrencyServer) List(*CurrencyListRequest, Currency_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCurrencyServer) Create(context.Context, *CurrencyCreateRequest) (*currency.Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCurrencyServer) Publish(context.Context, *currency.Currency_ID) (*currency.Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedCurrencyServer) Unpublish(context.Context, *currency.Currency_ID) (*currency.Currency, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpublish not implemented")
}
func (UnimplementedCurrencyServer) mustEmbedUnimplementedCurrencyServer() {}

// UnsafeCurrencyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServer will
// result in compilation errors.
type UnsafeCurrencyServer interface {
	mustEmbedUnimplementedCurrencyServer()
}

func RegisterCurrencyServer(s grpc.ServiceRegistrar, srv CurrencyServer) {
	s.RegisterService(&Currency_ServiceDesc, srv)
}

func _Currency_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(currency.Currency_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.Currency/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).Get(ctx, req.(*currency.Currency_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyServer).List(m, &currencyListServer{stream})
}

type Currency_ListServer interface {
	Send(*currency.Currency) error
	grpc.ServerStream
}

type currencyListServer struct {
	grpc.ServerStream
}

func (x *currencyListServer) Send(m *currency.Currency) error {
	return x.ServerStream.SendMsg(m)
}

func _Currency_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.Currency/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).Create(ctx, req.(*CurrencyCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(currency.Currency_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.Currency/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).Publish(ctx, req.(*currency.Currency_ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Currency_Unpublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(currency.Currency_ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).Unpublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/biconom.admin.currency.v1.Currency/Unpublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).Unpublish(ctx, req.(*currency.Currency_ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Currency_ServiceDesc is the grpc.ServiceDesc for Currency service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Currency_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "biconom.admin.currency.v1.Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Currency_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Currency_Create_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Currency_Publish_Handler,
		},
		{
			MethodName: "Unpublish",
			Handler:    _Currency_Unpublish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Currency_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "biconom/admin/currency/v1/currency.proto",
}
