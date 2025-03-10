// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: account/v1/account_service.proto

package acountpb

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
	Account_Register_FullMethodName              = "/flipchat.account.v1.Account/Register"
	Account_Login_FullMethodName                 = "/flipchat.account.v1.Account/Login"
	Account_AuthorizePublicKey_FullMethodName    = "/flipchat.account.v1.Account/AuthorizePublicKey"
	Account_RevokePublicKey_FullMethodName       = "/flipchat.account.v1.Account/RevokePublicKey"
	Account_GetPaymentDestination_FullMethodName = "/flipchat.account.v1.Account/GetPaymentDestination"
	Account_GetUserFlags_FullMethodName          = "/flipchat.account.v1.Account/GetUserFlags"
)

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	// Register registers a new user, bound to the provided PublicKey.
	// If the PublicKey is already in use, the previous user account is returned.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Login retrieves the UserId (and in the future, potentially other information)
	// required for 'recovering' an account.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// AuthorizePublicKey authorizes an additional PublicKey to an account.
	AuthorizePublicKey(ctx context.Context, in *AuthorizePublicKeyRequest, opts ...grpc.CallOption) (*AuthorizePublicKeyResponse, error)
	// RevokePublicKey revokes a public key from an account.
	//
	// There must be at least one public key per account. For now, any authorized public key
	// may revoke another public key, but this may change in the future.
	RevokePublicKey(ctx context.Context, in *RevokePublicKeyRequest, opts ...grpc.CallOption) (*RevokePublicKeyResponse, error)
	// GetPaymentDestination gets the payment destination for a UserId
	GetPaymentDestination(ctx context.Context, in *GetPaymentDestinationRequest, opts ...grpc.CallOption) (*GetPaymentDestinationResponse, error)
	// GetUserFlags gets user-specific flags
	GetUserFlags(ctx context.Context, in *GetUserFlagsRequest, opts ...grpc.CallOption) (*GetUserFlagsResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Account_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Account_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) AuthorizePublicKey(ctx context.Context, in *AuthorizePublicKeyRequest, opts ...grpc.CallOption) (*AuthorizePublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizePublicKeyResponse)
	err := c.cc.Invoke(ctx, Account_AuthorizePublicKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) RevokePublicKey(ctx context.Context, in *RevokePublicKeyRequest, opts ...grpc.CallOption) (*RevokePublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokePublicKeyResponse)
	err := c.cc.Invoke(ctx, Account_RevokePublicKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetPaymentDestination(ctx context.Context, in *GetPaymentDestinationRequest, opts ...grpc.CallOption) (*GetPaymentDestinationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentDestinationResponse)
	err := c.cc.Invoke(ctx, Account_GetPaymentDestination_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserFlags(ctx context.Context, in *GetUserFlagsRequest, opts ...grpc.CallOption) (*GetUserFlagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserFlagsResponse)
	err := c.cc.Invoke(ctx, Account_GetUserFlags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility.
type AccountServer interface {
	// Register registers a new user, bound to the provided PublicKey.
	// If the PublicKey is already in use, the previous user account is returned.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Login retrieves the UserId (and in the future, potentially other information)
	// required for 'recovering' an account.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// AuthorizePublicKey authorizes an additional PublicKey to an account.
	AuthorizePublicKey(context.Context, *AuthorizePublicKeyRequest) (*AuthorizePublicKeyResponse, error)
	// RevokePublicKey revokes a public key from an account.
	//
	// There must be at least one public key per account. For now, any authorized public key
	// may revoke another public key, but this may change in the future.
	RevokePublicKey(context.Context, *RevokePublicKeyRequest) (*RevokePublicKeyResponse, error)
	// GetPaymentDestination gets the payment destination for a UserId
	GetPaymentDestination(context.Context, *GetPaymentDestinationRequest) (*GetPaymentDestinationResponse, error)
	// GetUserFlags gets user-specific flags
	GetUserFlags(context.Context, *GetUserFlagsRequest) (*GetUserFlagsResponse, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountServer struct{}

func (UnimplementedAccountServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAccountServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAccountServer) AuthorizePublicKey(context.Context, *AuthorizePublicKeyRequest) (*AuthorizePublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizePublicKey not implemented")
}
func (UnimplementedAccountServer) RevokePublicKey(context.Context, *RevokePublicKeyRequest) (*RevokePublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePublicKey not implemented")
}
func (UnimplementedAccountServer) GetPaymentDestination(context.Context, *GetPaymentDestinationRequest) (*GetPaymentDestinationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentDestination not implemented")
}
func (UnimplementedAccountServer) GetUserFlags(context.Context, *GetUserFlagsRequest) (*GetUserFlagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFlags not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}
func (UnimplementedAccountServer) testEmbeddedByValue()                 {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	// If the following call pancis, it indicates UnimplementedAccountServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_AuthorizePublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizePublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).AuthorizePublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_AuthorizePublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).AuthorizePublicKey(ctx, req.(*AuthorizePublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_RevokePublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RevokePublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_RevokePublicKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RevokePublicKey(ctx, req.(*RevokePublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetPaymentDestination_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentDestinationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetPaymentDestination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetPaymentDestination_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetPaymentDestination(ctx, req.(*GetPaymentDestinationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetUserFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetUserFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Account_GetUserFlags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetUserFlags(ctx, req.(*GetUserFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipchat.account.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Account_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "AuthorizePublicKey",
			Handler:    _Account_AuthorizePublicKey_Handler,
		},
		{
			MethodName: "RevokePublicKey",
			Handler:    _Account_RevokePublicKey_Handler,
		},
		{
			MethodName: "GetPaymentDestination",
			Handler:    _Account_GetPaymentDestination_Handler,
		},
		{
			MethodName: "GetUserFlags",
			Handler:    _Account_GetUserFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1/account_service.proto",
}
