// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/merchant/merchant.proto

package merchant

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

// MerchantClient is the client API for Merchant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantClient interface {
	CreateMerchant(ctx context.Context, in *CreateMerchantRequest, opts ...grpc.CallOption) (*CreateMerchantReply, error)
	UpdateMerchant(ctx context.Context, in *UpdateMerchantRequest, opts ...grpc.CallOption) (*UpdateMerchantReply, error)
	DeleteMerchant(ctx context.Context, in *DeleteMerchantRequest, opts ...grpc.CallOption) (*DeleteMerchantReply, error)
	GetMerchant(ctx context.Context, in *GetMerchantRequest, opts ...grpc.CallOption) (*GetMerchantReply, error)
	ListMerchant(ctx context.Context, in *ListMerchantRequest, opts ...grpc.CallOption) (*ListMerchantReply, error)
}

type merchantClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantClient(cc grpc.ClientConnInterface) MerchantClient {
	return &merchantClient{cc}
}

func (c *merchantClient) CreateMerchant(ctx context.Context, in *CreateMerchantRequest, opts ...grpc.CallOption) (*CreateMerchantReply, error) {
	out := new(CreateMerchantReply)
	err := c.cc.Invoke(ctx, "/api.merchant.Merchant/CreateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) UpdateMerchant(ctx context.Context, in *UpdateMerchantRequest, opts ...grpc.CallOption) (*UpdateMerchantReply, error) {
	out := new(UpdateMerchantReply)
	err := c.cc.Invoke(ctx, "/api.merchant.Merchant/UpdateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) DeleteMerchant(ctx context.Context, in *DeleteMerchantRequest, opts ...grpc.CallOption) (*DeleteMerchantReply, error) {
	out := new(DeleteMerchantReply)
	err := c.cc.Invoke(ctx, "/api.merchant.Merchant/DeleteMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) GetMerchant(ctx context.Context, in *GetMerchantRequest, opts ...grpc.CallOption) (*GetMerchantReply, error) {
	out := new(GetMerchantReply)
	err := c.cc.Invoke(ctx, "/api.merchant.Merchant/GetMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantClient) ListMerchant(ctx context.Context, in *ListMerchantRequest, opts ...grpc.CallOption) (*ListMerchantReply, error) {
	out := new(ListMerchantReply)
	err := c.cc.Invoke(ctx, "/api.merchant.Merchant/ListMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServer is the server API for Merchant service.
// All implementations must embed UnimplementedMerchantServer
// for forward compatibility
type MerchantServer interface {
	CreateMerchant(context.Context, *CreateMerchantRequest) (*CreateMerchantReply, error)
	UpdateMerchant(context.Context, *UpdateMerchantRequest) (*UpdateMerchantReply, error)
	DeleteMerchant(context.Context, *DeleteMerchantRequest) (*DeleteMerchantReply, error)
	GetMerchant(context.Context, *GetMerchantRequest) (*GetMerchantReply, error)
	ListMerchant(context.Context, *ListMerchantRequest) (*ListMerchantReply, error)
	mustEmbedUnimplementedMerchantServer()
}

// UnimplementedMerchantServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServer struct {
}

func (UnimplementedMerchantServer) CreateMerchant(context.Context, *CreateMerchantRequest) (*CreateMerchantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMerchant not implemented")
}
func (UnimplementedMerchantServer) UpdateMerchant(context.Context, *UpdateMerchantRequest) (*UpdateMerchantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (UnimplementedMerchantServer) DeleteMerchant(context.Context, *DeleteMerchantRequest) (*DeleteMerchantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchant not implemented")
}
func (UnimplementedMerchantServer) GetMerchant(context.Context, *GetMerchantRequest) (*GetMerchantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchant not implemented")
}
func (UnimplementedMerchantServer) ListMerchant(context.Context, *ListMerchantRequest) (*ListMerchantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchant not implemented")
}
func (UnimplementedMerchantServer) mustEmbedUnimplementedMerchantServer() {}

// UnsafeMerchantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServer will
// result in compilation errors.
type UnsafeMerchantServer interface {
	mustEmbedUnimplementedMerchantServer()
}

func RegisterMerchantServer(s grpc.ServiceRegistrar, srv MerchantServer) {
	s.RegisterService(&Merchant_ServiceDesc, srv)
}

func _Merchant_CreateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).CreateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.merchant.Merchant/CreateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).CreateMerchant(ctx, req.(*CreateMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.merchant.Merchant/UpdateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).UpdateMerchant(ctx, req.(*UpdateMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_DeleteMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).DeleteMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.merchant.Merchant/DeleteMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).DeleteMerchant(ctx, req.(*DeleteMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_GetMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).GetMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.merchant.Merchant/GetMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).GetMerchant(ctx, req.(*GetMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Merchant_ListMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServer).ListMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.merchant.Merchant/ListMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServer).ListMerchant(ctx, req.(*ListMerchantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Merchant_ServiceDesc is the grpc.ServiceDesc for Merchant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Merchant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.merchant.Merchant",
	HandlerType: (*MerchantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMerchant",
			Handler:    _Merchant_CreateMerchant_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _Merchant_UpdateMerchant_Handler,
		},
		{
			MethodName: "DeleteMerchant",
			Handler:    _Merchant_DeleteMerchant_Handler,
		},
		{
			MethodName: "GetMerchant",
			Handler:    _Merchant_GetMerchant_Handler,
		},
		{
			MethodName: "ListMerchant",
			Handler:    _Merchant_ListMerchant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/merchant/merchant.proto",
}