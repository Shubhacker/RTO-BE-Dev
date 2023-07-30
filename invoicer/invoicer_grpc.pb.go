// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: invoicer.proto

package invoicer

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

// InvoicerClient is the client API for Invoicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvoicerClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponseMessage, error)
}

type invoicerClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoicerClient(cc grpc.ClientConnInterface) InvoicerClient {
	return &invoicerClient{cc}
}

func (c *invoicerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponseMessage, error) {
	out := new(CreateResponseMessage)
	err := c.cc.Invoke(ctx, "/Invoicer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoicerServer is the server API for Invoicer service.
// All implementations must embed UnimplementedInvoicerServer
// for forward compatibility
type InvoicerServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponseMessage, error)
	mustEmbedUnimplementedInvoicerServer()
}

// UnimplementedInvoicerServer must be embedded to have forward compatible implementations.
type UnimplementedInvoicerServer struct {
}

func (UnimplementedInvoicerServer) Create(context.Context, *CreateRequest) (*CreateResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInvoicerServer) mustEmbedUnimplementedInvoicerServer() {}

// UnsafeInvoicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvoicerServer will
// result in compilation errors.
type UnsafeInvoicerServer interface {
	mustEmbedUnimplementedInvoicerServer()
}

func RegisterInvoicerServer(s grpc.ServiceRegistrar, srv InvoicerServer) {
	s.RegisterService(&Invoicer_ServiceDesc, srv)
}

func _Invoicer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoicerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Invoicer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoicerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Invoicer_ServiceDesc is the grpc.ServiceDesc for Invoicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Invoicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Invoicer",
	HandlerType: (*InvoicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoicer_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoicer.proto",
}
