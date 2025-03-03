// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: preview/v1/preview_service.proto

package previewpb

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
	Preview_GetPreviewUrl_FullMethodName = "/flipchat.preview.v1.Preview/GetPreviewUrl"
)

// PreviewClient is the client API for Preview service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PreviewClient interface {
	GetPreviewUrl(ctx context.Context, in *GetPreviewUrlRequest, opts ...grpc.CallOption) (*GetPreviewUrlResponse, error)
}

type previewClient struct {
	cc grpc.ClientConnInterface
}

func NewPreviewClient(cc grpc.ClientConnInterface) PreviewClient {
	return &previewClient{cc}
}

func (c *previewClient) GetPreviewUrl(ctx context.Context, in *GetPreviewUrlRequest, opts ...grpc.CallOption) (*GetPreviewUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPreviewUrlResponse)
	err := c.cc.Invoke(ctx, Preview_GetPreviewUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PreviewServer is the server API for Preview service.
// All implementations must embed UnimplementedPreviewServer
// for forward compatibility.
type PreviewServer interface {
	GetPreviewUrl(context.Context, *GetPreviewUrlRequest) (*GetPreviewUrlResponse, error)
	mustEmbedUnimplementedPreviewServer()
}

// UnimplementedPreviewServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPreviewServer struct{}

func (UnimplementedPreviewServer) GetPreviewUrl(context.Context, *GetPreviewUrlRequest) (*GetPreviewUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreviewUrl not implemented")
}
func (UnimplementedPreviewServer) mustEmbedUnimplementedPreviewServer() {}
func (UnimplementedPreviewServer) testEmbeddedByValue()                 {}

// UnsafePreviewServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PreviewServer will
// result in compilation errors.
type UnsafePreviewServer interface {
	mustEmbedUnimplementedPreviewServer()
}

func RegisterPreviewServer(s grpc.ServiceRegistrar, srv PreviewServer) {
	// If the following call pancis, it indicates UnimplementedPreviewServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Preview_ServiceDesc, srv)
}

func _Preview_GetPreviewUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreviewUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PreviewServer).GetPreviewUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Preview_GetPreviewUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PreviewServer).GetPreviewUrl(ctx, req.(*GetPreviewUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Preview_ServiceDesc is the grpc.ServiceDesc for Preview service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Preview_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipchat.preview.v1.Preview",
	HandlerType: (*PreviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPreviewUrl",
			Handler:    _Preview_GetPreviewUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preview/v1/preview_service.proto",
}
