// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: messaging/v1/messaging_service.proto

package messagingpb

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
	Messaging_StreamMessages_FullMethodName = "/flipchat.messaging.v1.Messaging/StreamMessages"
	Messaging_GetMessage_FullMethodName     = "/flipchat.messaging.v1.Messaging/GetMessage"
	Messaging_GetMessages_FullMethodName    = "/flipchat.messaging.v1.Messaging/GetMessages"
	Messaging_SendMessage_FullMethodName    = "/flipchat.messaging.v1.Messaging/SendMessage"
	Messaging_AdvancePointer_FullMethodName = "/flipchat.messaging.v1.Messaging/AdvancePointer"
	Messaging_NotifyIsTyping_FullMethodName = "/flipchat.messaging.v1.Messaging/NotifyIsTyping"
)

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	// StreamMessages streams all messages/message states for the requested chat.
	StreamMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamMessagesRequest, StreamMessagesResponse], error)
	// GetMessage gets a single message in a chat
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
	// GetMessages gets the set of messages for a chat using a paged and batched APIs
	GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error)
	// SendMessage sends a message to a chat.
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	// AdvancePointer advances a pointer in message history for a chat member.
	AdvancePointer(ctx context.Context, in *AdvancePointerRequest, opts ...grpc.CallOption) (*AdvancePointerResponse, error)
	// NotifyIsTypingRequest notifies a chat that the sending member is typing.
	//
	// These requests are transient, and may be dropped at any point.
	NotifyIsTyping(ctx context.Context, in *NotifyIsTypingRequest, opts ...grpc.CallOption) (*NotifyIsTypingResponse, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) StreamMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamMessagesRequest, StreamMessagesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[0], Messaging_StreamMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamMessagesRequest, StreamMessagesResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_StreamMessagesClient = grpc.BidiStreamingClient[StreamMessagesRequest, StreamMessagesResponse]

func (c *messagingClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, Messaging_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) GetMessages(ctx context.Context, in *GetMessagesRequest, opts ...grpc.CallOption) (*GetMessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMessagesResponse)
	err := c.cc.Invoke(ctx, Messaging_GetMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Messaging_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) AdvancePointer(ctx context.Context, in *AdvancePointerRequest, opts ...grpc.CallOption) (*AdvancePointerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdvancePointerResponse)
	err := c.cc.Invoke(ctx, Messaging_AdvancePointer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingClient) NotifyIsTyping(ctx context.Context, in *NotifyIsTypingRequest, opts ...grpc.CallOption) (*NotifyIsTypingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NotifyIsTypingResponse)
	err := c.cc.Invoke(ctx, Messaging_NotifyIsTyping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility.
type MessagingServer interface {
	// StreamMessages streams all messages/message states for the requested chat.
	StreamMessages(grpc.BidiStreamingServer[StreamMessagesRequest, StreamMessagesResponse]) error
	// GetMessage gets a single message in a chat
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	// GetMessages gets the set of messages for a chat using a paged and batched APIs
	GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error)
	// SendMessage sends a message to a chat.
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	// AdvancePointer advances a pointer in message history for a chat member.
	AdvancePointer(context.Context, *AdvancePointerRequest) (*AdvancePointerResponse, error)
	// NotifyIsTypingRequest notifies a chat that the sending member is typing.
	//
	// These requests are transient, and may be dropped at any point.
	NotifyIsTyping(context.Context, *NotifyIsTypingRequest) (*NotifyIsTypingResponse, error)
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessagingServer struct{}

func (UnimplementedMessagingServer) StreamMessages(grpc.BidiStreamingServer[StreamMessagesRequest, StreamMessagesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamMessages not implemented")
}
func (UnimplementedMessagingServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessagingServer) GetMessages(context.Context, *GetMessagesRequest) (*GetMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedMessagingServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedMessagingServer) AdvancePointer(context.Context, *AdvancePointerRequest) (*AdvancePointerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdvancePointer not implemented")
}
func (UnimplementedMessagingServer) NotifyIsTyping(context.Context, *NotifyIsTypingRequest) (*NotifyIsTypingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyIsTyping not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}
func (UnimplementedMessagingServer) testEmbeddedByValue()                   {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	// If the following call pancis, it indicates UnimplementedMessagingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_StreamMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).StreamMessages(&grpc.GenericServerStream[StreamMessagesRequest, StreamMessagesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Messaging_StreamMessagesServer = grpc.BidiStreamingServer[StreamMessagesRequest, StreamMessagesResponse]

func _Messaging_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessages(ctx, req.(*GetMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_AdvancePointer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvancePointerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).AdvancePointer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_AdvancePointer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).AdvancePointer(ctx, req.(*AdvancePointerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messaging_NotifyIsTyping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyIsTypingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).NotifyIsTyping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messaging_NotifyIsTyping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).NotifyIsTyping(ctx, req.(*NotifyIsTypingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipchat.messaging.v1.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _Messaging_GetMessage_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _Messaging_GetMessages_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Messaging_SendMessage_Handler,
		},
		{
			MethodName: "AdvancePointer",
			Handler:    _Messaging_AdvancePointer_Handler,
		},
		{
			MethodName: "NotifyIsTyping",
			Handler:    _Messaging_NotifyIsTyping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMessages",
			Handler:       _Messaging_StreamMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messaging/v1/messaging_service.proto",
}
