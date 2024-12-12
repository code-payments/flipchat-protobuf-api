// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: iap/v1/iap_service.proto

package iappb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Iap_OnPurchaseCompleted_FullMethodName = "/flipchat.iap.v1.Iap/OnPurchaseCompleted"
)

// IapClient is the client API for Iap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IapClient interface {
	// OnPurchaseCompleted is called when an IAP has been completed
	OnPurchaseCompleted(ctx context.Context, in *OnPurchaseCompletedRequest, opts ...grpc.CallOption) (*OnPurchaseCompletedResponse, error)
}

type iapClient struct {
	cc grpc.ClientConnInterface
}

func NewIapClient(cc grpc.ClientConnInterface) IapClient {
	return &iapClient{cc}
}

func (c *iapClient) OnPurchaseCompleted(ctx context.Context, in *OnPurchaseCompletedRequest, opts ...grpc.CallOption) (*OnPurchaseCompletedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnPurchaseCompletedResponse)
	err := c.cc.Invoke(ctx, Iap_OnPurchaseCompleted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IapServer is the server API for Iap service.
// All implementations must embed UnimplementedIapServer
// for forward compatibility.
type IapServer interface {
	// OnPurchaseCompleted is called when an IAP has been completed
	OnPurchaseCompleted(context.Context, *OnPurchaseCompletedRequest) (*OnPurchaseCompletedResponse, error)
	mustEmbedUnimplementedIapServer()
}

// UnimplementedIapServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIapServer struct{}

func (UnimplementedIapServer) OnPurchaseCompleted(context.Context, *OnPurchaseCompletedRequest) (*OnPurchaseCompletedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnPurchaseCompleted not implemented")
}
func (UnimplementedIapServer) mustEmbedUnimplementedIapServer() {}
func (UnimplementedIapServer) testEmbeddedByValue()             {}

// UnsafeIapServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IapServer will
// result in compilation errors.
type UnsafeIapServer interface {
	mustEmbedUnimplementedIapServer()
}

func RegisterIapServer(s grpc.ServiceRegistrar, srv IapServer) {
	// If the following call pancis, it indicates UnimplementedIapServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Iap_ServiceDesc, srv)
}

func _Iap_OnPurchaseCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnPurchaseCompletedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IapServer).OnPurchaseCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Iap_OnPurchaseCompleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IapServer).OnPurchaseCompleted(ctx, req.(*OnPurchaseCompletedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Iap_ServiceDesc is the grpc.ServiceDesc for Iap service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Iap_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipchat.iap.v1.Iap",
	HandlerType: (*IapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnPurchaseCompleted",
			Handler:    _Iap_OnPurchaseCompleted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iap/v1/iap_service.proto",
}
