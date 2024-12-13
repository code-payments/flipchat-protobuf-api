// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: push/v1/push_service.proto

package pushpb

import (
	v1 "github.com/code-payments/flipchat-protobuf-api/generated/go/common/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenType int32

const (
	TokenType_UNKNOWN TokenType = 0
	// FCM registration token for an Android device
	TokenType_FCM_ANDROID TokenType = 1
	// FCM registration token or an iOS device
	TokenType_FCM_APNS TokenType = 2
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "UNKNOWN",
		1: "FCM_ANDROID",
		2: "FCM_APNS",
	}
	TokenType_value = map[string]int32{
		"UNKNOWN":     0,
		"FCM_ANDROID": 1,
		"FCM_APNS":    2,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{0}
}

type AddTokenResponse_Result int32

const (
	AddTokenResponse_OK                 AddTokenResponse_Result = 0
	AddTokenResponse_INVALID_PUSH_TOKEN AddTokenResponse_Result = 1
)

// Enum value maps for AddTokenResponse_Result.
var (
	AddTokenResponse_Result_name = map[int32]string{
		0: "OK",
		1: "INVALID_PUSH_TOKEN",
	}
	AddTokenResponse_Result_value = map[string]int32{
		"OK":                 0,
		"INVALID_PUSH_TOKEN": 1,
	}
)

func (x AddTokenResponse_Result) Enum() *AddTokenResponse_Result {
	p := new(AddTokenResponse_Result)
	*p = x
	return p
}

func (x AddTokenResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddTokenResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[1].Descriptor()
}

func (AddTokenResponse_Result) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[1]
}

func (x AddTokenResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddTokenResponse_Result.Descriptor instead.
func (AddTokenResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{1, 0}
}

type DeleteTokenResponse_Result int32

const (
	DeleteTokenResponse_OK DeleteTokenResponse_Result = 0
)

// Enum value maps for DeleteTokenResponse_Result.
var (
	DeleteTokenResponse_Result_name = map[int32]string{
		0: "OK",
	}
	DeleteTokenResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x DeleteTokenResponse_Result) Enum() *DeleteTokenResponse_Result {
	p := new(DeleteTokenResponse_Result)
	*p = x
	return p
}

func (x DeleteTokenResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteTokenResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[2].Descriptor()
}

func (DeleteTokenResponse_Result) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[2]
}

func (x DeleteTokenResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteTokenResponse_Result.Descriptor instead.
func (DeleteTokenResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{3, 0}
}

type DeleteTokensResponse_Result int32

const (
	DeleteTokensResponse_OK DeleteTokensResponse_Result = 0
)

// Enum value maps for DeleteTokensResponse_Result.
var (
	DeleteTokensResponse_Result_name = map[int32]string{
		0: "OK",
	}
	DeleteTokensResponse_Result_value = map[string]int32{
		"OK": 0,
	}
)

func (x DeleteTokensResponse_Result) Enum() *DeleteTokensResponse_Result {
	p := new(DeleteTokensResponse_Result)
	*p = x
	return p
}

func (x DeleteTokensResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteTokensResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_push_v1_push_service_proto_enumTypes[3].Descriptor()
}

func (DeleteTokensResponse_Result) Type() protoreflect.EnumType {
	return &file_push_v1_push_service_proto_enumTypes[3]
}

func (x DeleteTokensResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteTokensResponse_Result.Descriptor instead.
func (DeleteTokensResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{5, 0}
}

type AddTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType  TokenType        `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=flipchat.push.v1.TokenType" json:"token_type,omitempty"`
	PushToken  string           `protobuf:"bytes,2,opt,name=push_token,json=pushToken,proto3" json:"push_token,omitempty"`
	AppInstall *v1.AppInstallId `protobuf:"bytes,3,opt,name=app_install,json=appInstall,proto3" json:"app_install,omitempty"`
	Auth       *v1.Auth         `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *AddTokenRequest) Reset() {
	*x = AddTokenRequest{}
	mi := &file_push_v1_push_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTokenRequest) ProtoMessage() {}

func (x *AddTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTokenRequest.ProtoReflect.Descriptor instead.
func (*AddTokenRequest) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddTokenRequest) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_UNKNOWN
}

func (x *AddTokenRequest) GetPushToken() string {
	if x != nil {
		return x.PushToken
	}
	return ""
}

func (x *AddTokenRequest) GetAppInstall() *v1.AppInstallId {
	if x != nil {
		return x.AppInstall
	}
	return nil
}

func (x *AddTokenRequest) GetAuth() *v1.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type AddTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result AddTokenResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.push.v1.AddTokenResponse_Result" json:"result,omitempty"`
}

func (x *AddTokenResponse) Reset() {
	*x = AddTokenResponse{}
	mi := &file_push_v1_push_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTokenResponse) ProtoMessage() {}

func (x *AddTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTokenResponse.ProtoReflect.Descriptor instead.
func (*AddTokenResponse) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddTokenResponse) GetResult() AddTokenResponse_Result {
	if x != nil {
		return x.Result
	}
	return AddTokenResponse_OK
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenType TokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=flipchat.push.v1.TokenType" json:"token_type,omitempty"`
	PushToken string    `protobuf:"bytes,2,opt,name=push_token,json=pushToken,proto3" json:"push_token,omitempty"`
	Auth      *v1.Auth  `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	mi := &file_push_v1_push_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTokenRequest) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_UNKNOWN
}

func (x *DeleteTokenRequest) GetPushToken() string {
	if x != nil {
		return x.PushToken
	}
	return ""
}

func (x *DeleteTokenRequest) GetAuth() *v1.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type DeleteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result DeleteTokenResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.push.v1.DeleteTokenResponse_Result" json:"result,omitempty"`
}

func (x *DeleteTokenResponse) Reset() {
	*x = DeleteTokenResponse{}
	mi := &file_push_v1_push_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenResponse) ProtoMessage() {}

func (x *DeleteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTokenResponse) GetResult() DeleteTokenResponse_Result {
	if x != nil {
		return x.Result
	}
	return DeleteTokenResponse_OK
}

type DeleteTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInstall *v1.AppInstallId `protobuf:"bytes,1,opt,name=app_install,json=appInstall,proto3" json:"app_install,omitempty"`
	Auth       *v1.Auth         `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *DeleteTokensRequest) Reset() {
	*x = DeleteTokensRequest{}
	mi := &file_push_v1_push_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokensRequest) ProtoMessage() {}

func (x *DeleteTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokensRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokensRequest) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTokensRequest) GetAppInstall() *v1.AppInstallId {
	if x != nil {
		return x.AppInstall
	}
	return nil
}

func (x *DeleteTokensRequest) GetAuth() *v1.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type DeleteTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result DeleteTokensResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.push.v1.DeleteTokensResponse_Result" json:"result,omitempty"`
}

func (x *DeleteTokensResponse) Reset() {
	*x = DeleteTokensResponse{}
	mi := &file_push_v1_push_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokensResponse) ProtoMessage() {}

func (x *DeleteTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_push_v1_push_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokensResponse.ProtoReflect.Descriptor instead.
func (*DeleteTokensResponse) Descriptor() ([]byte, []int) {
	return file_push_v1_push_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTokensResponse) GetResult() DeleteTokensResponse_Result {
	if x != nil {
		return x.Result
	}
	return DeleteTokensResponse_OK
}

var File_push_v1_push_service_proto protoreflect.FileDescriptor

var file_push_v1_push_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf5, 0x01, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x18, 0x01, 0x18, 0x02,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x70,
	0x75, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x20, 0x52, 0x09, 0x70, 0x75, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x52, 0x0a, 0x61,
	0x70, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x2c, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x7f, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x55, 0x53, 0x48,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x46, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x82, 0x01, 0x04, 0x18, 0x01, 0x18, 0x02, 0x52, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x20, 0x52, 0x09, 0x70, 0x75, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x6d, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x10, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x6f, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x10, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x2a, 0x37, 0x0a,
	0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x43, 0x4d, 0x5f, 0x41,
	0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x43, 0x4d, 0x5f,
	0x41, 0x50, 0x4e, 0x53, 0x10, 0x02, 0x32, 0x94, 0x02, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12,
	0x51, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x24, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x25,
	0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7a, 0x0a,
	0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69, 0x6e, 0x63, 0x2e, 0x66, 0x6c, 0x69,
	0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x2d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70,
	0x63, 0x68, 0x61, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x73, 0x68, 0x70, 0x62, 0xa2, 0x02, 0x09,
	0x46, 0x50, 0x42, 0x50, 0x75, 0x73, 0x68, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_push_v1_push_service_proto_rawDescOnce sync.Once
	file_push_v1_push_service_proto_rawDescData = file_push_v1_push_service_proto_rawDesc
)

func file_push_v1_push_service_proto_rawDescGZIP() []byte {
	file_push_v1_push_service_proto_rawDescOnce.Do(func() {
		file_push_v1_push_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_v1_push_service_proto_rawDescData)
	})
	return file_push_v1_push_service_proto_rawDescData
}

var file_push_v1_push_service_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_push_v1_push_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_push_v1_push_service_proto_goTypes = []any{
	(TokenType)(0),                   // 0: flipchat.push.v1.TokenType
	(AddTokenResponse_Result)(0),     // 1: flipchat.push.v1.AddTokenResponse.Result
	(DeleteTokenResponse_Result)(0),  // 2: flipchat.push.v1.DeleteTokenResponse.Result
	(DeleteTokensResponse_Result)(0), // 3: flipchat.push.v1.DeleteTokensResponse.Result
	(*AddTokenRequest)(nil),          // 4: flipchat.push.v1.AddTokenRequest
	(*AddTokenResponse)(nil),         // 5: flipchat.push.v1.AddTokenResponse
	(*DeleteTokenRequest)(nil),       // 6: flipchat.push.v1.DeleteTokenRequest
	(*DeleteTokenResponse)(nil),      // 7: flipchat.push.v1.DeleteTokenResponse
	(*DeleteTokensRequest)(nil),      // 8: flipchat.push.v1.DeleteTokensRequest
	(*DeleteTokensResponse)(nil),     // 9: flipchat.push.v1.DeleteTokensResponse
	(*v1.AppInstallId)(nil),          // 10: flipchat.common.v1.AppInstallId
	(*v1.Auth)(nil),                  // 11: flipchat.common.v1.Auth
}
var file_push_v1_push_service_proto_depIdxs = []int32{
	0,  // 0: flipchat.push.v1.AddTokenRequest.token_type:type_name -> flipchat.push.v1.TokenType
	10, // 1: flipchat.push.v1.AddTokenRequest.app_install:type_name -> flipchat.common.v1.AppInstallId
	11, // 2: flipchat.push.v1.AddTokenRequest.auth:type_name -> flipchat.common.v1.Auth
	1,  // 3: flipchat.push.v1.AddTokenResponse.result:type_name -> flipchat.push.v1.AddTokenResponse.Result
	0,  // 4: flipchat.push.v1.DeleteTokenRequest.token_type:type_name -> flipchat.push.v1.TokenType
	11, // 5: flipchat.push.v1.DeleteTokenRequest.auth:type_name -> flipchat.common.v1.Auth
	2,  // 6: flipchat.push.v1.DeleteTokenResponse.result:type_name -> flipchat.push.v1.DeleteTokenResponse.Result
	10, // 7: flipchat.push.v1.DeleteTokensRequest.app_install:type_name -> flipchat.common.v1.AppInstallId
	11, // 8: flipchat.push.v1.DeleteTokensRequest.auth:type_name -> flipchat.common.v1.Auth
	3,  // 9: flipchat.push.v1.DeleteTokensResponse.result:type_name -> flipchat.push.v1.DeleteTokensResponse.Result
	4,  // 10: flipchat.push.v1.Push.AddToken:input_type -> flipchat.push.v1.AddTokenRequest
	6,  // 11: flipchat.push.v1.Push.DeleteToken:input_type -> flipchat.push.v1.DeleteTokenRequest
	8,  // 12: flipchat.push.v1.Push.DeleteTokens:input_type -> flipchat.push.v1.DeleteTokensRequest
	5,  // 13: flipchat.push.v1.Push.AddToken:output_type -> flipchat.push.v1.AddTokenResponse
	7,  // 14: flipchat.push.v1.Push.DeleteToken:output_type -> flipchat.push.v1.DeleteTokenResponse
	9,  // 15: flipchat.push.v1.Push.DeleteTokens:output_type -> flipchat.push.v1.DeleteTokensResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_push_v1_push_service_proto_init() }
func file_push_v1_push_service_proto_init() {
	if File_push_v1_push_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_push_v1_push_service_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_push_v1_push_service_proto_goTypes,
		DependencyIndexes: file_push_v1_push_service_proto_depIdxs,
		EnumInfos:         file_push_v1_push_service_proto_enumTypes,
		MessageInfos:      file_push_v1_push_service_proto_msgTypes,
	}.Build()
	File_push_v1_push_service_proto = out.File
	file_push_v1_push_service_proto_rawDesc = nil
	file_push_v1_push_service_proto_goTypes = nil
	file_push_v1_push_service_proto_depIdxs = nil
}
