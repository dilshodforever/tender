// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tender.proto

package genprotos

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

// TenderServiceClient is the client API for TenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenderServiceClient interface {
	CreateTender(ctx context.Context, in *CreateTenderRequest, opts ...grpc.CallOption) (*TenderResponse, error)
	TenderAward(ctx context.Context, in *CreatTenderAwardRequest, opts ...grpc.CallOption) (*TenderResponse, error)
	DeleteTender(ctx context.Context, in *TenderIdRequest, opts ...grpc.CallOption) (*TenderResponse, error)
	ListTenders(ctx context.Context, in *ListTendersRequest, opts ...grpc.CallOption) (*ListTendersResponse, error)
	UpdateTender(ctx context.Context, in *UpdateTenderRequest, opts ...grpc.CallOption) (*TenderResponse, error)
	ListUserTenders(ctx context.Context, in *TenderIdRequest, opts ...grpc.CallOption) (*ListTendersResponse, error)
}

type tenderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenderServiceClient(cc grpc.ClientConnInterface) TenderServiceClient {
	return &tenderServiceClient{cc}
}

func (c *tenderServiceClient) CreateTender(ctx context.Context, in *CreateTenderRequest, opts ...grpc.CallOption) (*TenderResponse, error) {
	out := new(TenderResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/CreateTender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenderServiceClient) TenderAward(ctx context.Context, in *CreatTenderAwardRequest, opts ...grpc.CallOption) (*TenderResponse, error) {
	out := new(TenderResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/TenderAward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenderServiceClient) DeleteTender(ctx context.Context, in *TenderIdRequest, opts ...grpc.CallOption) (*TenderResponse, error) {
	out := new(TenderResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/DeleteTender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenderServiceClient) ListTenders(ctx context.Context, in *ListTendersRequest, opts ...grpc.CallOption) (*ListTendersResponse, error) {
	out := new(ListTendersResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/ListTenders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenderServiceClient) UpdateTender(ctx context.Context, in *UpdateTenderRequest, opts ...grpc.CallOption) (*TenderResponse, error) {
	out := new(TenderResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/UpdateTender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenderServiceClient) ListUserTenders(ctx context.Context, in *TenderIdRequest, opts ...grpc.CallOption) (*ListTendersResponse, error) {
	out := new(ListTendersResponse)
	err := c.cc.Invoke(ctx, "/protos.TenderService/ListUserTenders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenderServiceServer is the server API for TenderService service.
// All implementations must embed UnimplementedTenderServiceServer
// for forward compatibility
type TenderServiceServer interface {
	CreateTender(context.Context, *CreateTenderRequest) (*TenderResponse, error)
	TenderAward(context.Context, *CreatTenderAwardRequest) (*TenderResponse, error)
	DeleteTender(context.Context, *TenderIdRequest) (*TenderResponse, error)
	ListTenders(context.Context, *ListTendersRequest) (*ListTendersResponse, error)
	UpdateTender(context.Context, *UpdateTenderRequest) (*TenderResponse, error)
	ListUserTenders(context.Context, *TenderIdRequest) (*ListTendersResponse, error)
	mustEmbedUnimplementedTenderServiceServer()
}

// UnimplementedTenderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTenderServiceServer struct {
}

func (UnimplementedTenderServiceServer) CreateTender(context.Context, *CreateTenderRequest) (*TenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTender not implemented")
}
func (UnimplementedTenderServiceServer) TenderAward(context.Context, *CreatTenderAwardRequest) (*TenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TenderAward not implemented")
}
func (UnimplementedTenderServiceServer) DeleteTender(context.Context, *TenderIdRequest) (*TenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTender not implemented")
}
func (UnimplementedTenderServiceServer) ListTenders(context.Context, *ListTendersRequest) (*ListTendersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenders not implemented")
}
func (UnimplementedTenderServiceServer) UpdateTender(context.Context, *UpdateTenderRequest) (*TenderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTender not implemented")
}
func (UnimplementedTenderServiceServer) ListUserTenders(context.Context, *TenderIdRequest) (*ListTendersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTenders not implemented")
}
func (UnimplementedTenderServiceServer) mustEmbedUnimplementedTenderServiceServer() {}

// UnsafeTenderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenderServiceServer will
// result in compilation errors.
type UnsafeTenderServiceServer interface {
	mustEmbedUnimplementedTenderServiceServer()
}

func RegisterTenderServiceServer(s grpc.ServiceRegistrar, srv TenderServiceServer) {
	s.RegisterService(&TenderService_ServiceDesc, srv)
}

func _TenderService_CreateTender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).CreateTender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/CreateTender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).CreateTender(ctx, req.(*CreateTenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenderService_TenderAward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatTenderAwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).TenderAward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/TenderAward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).TenderAward(ctx, req.(*CreatTenderAwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenderService_DeleteTender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).DeleteTender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/DeleteTender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).DeleteTender(ctx, req.(*TenderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenderService_ListTenders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTendersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).ListTenders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/ListTenders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).ListTenders(ctx, req.(*ListTendersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenderService_UpdateTender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).UpdateTender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/UpdateTender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).UpdateTender(ctx, req.(*UpdateTenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenderService_ListUserTenders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenderIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenderServiceServer).ListUserTenders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.TenderService/ListUserTenders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenderServiceServer).ListUserTenders(ctx, req.(*TenderIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenderService_ServiceDesc is the grpc.ServiceDesc for TenderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TenderService",
	HandlerType: (*TenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTender",
			Handler:    _TenderService_CreateTender_Handler,
		},
		{
			MethodName: "TenderAward",
			Handler:    _TenderService_TenderAward_Handler,
		},
		{
			MethodName: "DeleteTender",
			Handler:    _TenderService_DeleteTender_Handler,
		},
		{
			MethodName: "ListTenders",
			Handler:    _TenderService_ListTenders_Handler,
		},
		{
			MethodName: "UpdateTender",
			Handler:    _TenderService_UpdateTender_Handler,
		},
		{
			MethodName: "ListUserTenders",
			Handler:    _TenderService_ListUserTenders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tender.proto",
}
