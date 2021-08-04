// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pay_business

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

// PayBusinessServiceClient is the client API for PayBusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayBusinessServiceClient interface {
	// 统一收单支付
	TradePay(ctx context.Context, in *TradePayRequest, opts ...grpc.CallOption) (*TradePayResponse, error)
	// 创建账户
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	// 获取账户
	FindAccount(ctx context.Context, in *FindAccountRequest, opts ...grpc.CallOption) (*FindAccountResponse, error)
	// 账户充值
	AccountCharge(ctx context.Context, in *AccountChargeRequest, opts ...grpc.CallOption) (*AccountChargeResponse, error)
	// 获取交易唯一ID
	GetTradeUUID(ctx context.Context, in *GetTradeUUIDRequest, opts ...grpc.CallOption) (*GetTradeUUIDResponse, error)
}

type payBusinessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayBusinessServiceClient(cc grpc.ClientConnInterface) PayBusinessServiceClient {
	return &payBusinessServiceClient{cc}
}

func (c *payBusinessServiceClient) TradePay(ctx context.Context, in *TradePayRequest, opts ...grpc.CallOption) (*TradePayResponse, error) {
	out := new(TradePayResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/TradePay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payBusinessServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payBusinessServiceClient) FindAccount(ctx context.Context, in *FindAccountRequest, opts ...grpc.CallOption) (*FindAccountResponse, error) {
	out := new(FindAccountResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/FindAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payBusinessServiceClient) AccountCharge(ctx context.Context, in *AccountChargeRequest, opts ...grpc.CallOption) (*AccountChargeResponse, error) {
	out := new(AccountChargeResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/AccountCharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payBusinessServiceClient) GetTradeUUID(ctx context.Context, in *GetTradeUUIDRequest, opts ...grpc.CallOption) (*GetTradeUUIDResponse, error) {
	out := new(GetTradeUUIDResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/GetTradeUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayBusinessServiceServer is the server API for PayBusinessService service.
// All implementations must embed UnimplementedPayBusinessServiceServer
// for forward compatibility
type PayBusinessServiceServer interface {
	// 统一收单支付
	TradePay(context.Context, *TradePayRequest) (*TradePayResponse, error)
	// 创建账户
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	// 获取账户
	FindAccount(context.Context, *FindAccountRequest) (*FindAccountResponse, error)
	// 账户充值
	AccountCharge(context.Context, *AccountChargeRequest) (*AccountChargeResponse, error)
	// 获取交易唯一ID
	GetTradeUUID(context.Context, *GetTradeUUIDRequest) (*GetTradeUUIDResponse, error)
	mustEmbedUnimplementedPayBusinessServiceServer()
}

// UnimplementedPayBusinessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayBusinessServiceServer struct {
}

func (UnimplementedPayBusinessServiceServer) TradePay(context.Context, *TradePayRequest) (*TradePayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradePay not implemented")
}
func (UnimplementedPayBusinessServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedPayBusinessServiceServer) FindAccount(context.Context, *FindAccountRequest) (*FindAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAccount not implemented")
}
func (UnimplementedPayBusinessServiceServer) AccountCharge(context.Context, *AccountChargeRequest) (*AccountChargeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountCharge not implemented")
}
func (UnimplementedPayBusinessServiceServer) GetTradeUUID(context.Context, *GetTradeUUIDRequest) (*GetTradeUUIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradeUUID not implemented")
}
func (UnimplementedPayBusinessServiceServer) mustEmbedUnimplementedPayBusinessServiceServer() {}

// UnsafePayBusinessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayBusinessServiceServer will
// result in compilation errors.
type UnsafePayBusinessServiceServer interface {
	mustEmbedUnimplementedPayBusinessServiceServer()
}

func RegisterPayBusinessServiceServer(s grpc.ServiceRegistrar, srv PayBusinessServiceServer) {
	s.RegisterService(&PayBusinessService_ServiceDesc, srv)
}

func _PayBusinessService_TradePay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradePayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).TradePay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/TradePay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).TradePay(ctx, req.(*TradePayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayBusinessService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayBusinessService_FindAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).FindAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/FindAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).FindAccount(ctx, req.(*FindAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayBusinessService_AccountCharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountChargeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).AccountCharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/AccountCharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).AccountCharge(ctx, req.(*AccountChargeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayBusinessService_GetTradeUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradeUUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).GetTradeUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/GetTradeUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).GetTradeUUID(ctx, req.(*GetTradeUUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayBusinessService_ServiceDesc is the grpc.ServiceDesc for PayBusinessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayBusinessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pay_business.PayBusinessService",
	HandlerType: (*PayBusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradePay",
			Handler:    _PayBusinessService_TradePay_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _PayBusinessService_CreateAccount_Handler,
		},
		{
			MethodName: "FindAccount",
			Handler:    _PayBusinessService_FindAccount_Handler,
		},
		{
			MethodName: "AccountCharge",
			Handler:    _PayBusinessService_AccountCharge_Handler,
		},
		{
			MethodName: "GetTradeUUID",
			Handler:    _PayBusinessService_GetTradeUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall_pay_proto/pay_business/pay_business.proto",
}