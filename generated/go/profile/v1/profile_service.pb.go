// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: profile/v1/profile_service.proto

package profilepb

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

type GetProfileResponse_Result int32

const (
	GetProfileResponse_OK        GetProfileResponse_Result = 0
	GetProfileResponse_NOT_FOUND GetProfileResponse_Result = 1
)

// Enum value maps for GetProfileResponse_Result.
var (
	GetProfileResponse_Result_name = map[int32]string{
		0: "OK",
		1: "NOT_FOUND",
	}
	GetProfileResponse_Result_value = map[string]int32{
		"OK":        0,
		"NOT_FOUND": 1,
	}
)

func (x GetProfileResponse_Result) Enum() *GetProfileResponse_Result {
	p := new(GetProfileResponse_Result)
	*p = x
	return p
}

func (x GetProfileResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetProfileResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_v1_profile_service_proto_enumTypes[0].Descriptor()
}

func (GetProfileResponse_Result) Type() protoreflect.EnumType {
	return &file_profile_v1_profile_service_proto_enumTypes[0]
}

func (x GetProfileResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetProfileResponse_Result.Descriptor instead.
func (GetProfileResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{1, 0}
}

type SetDisplayNameResponse_Result int32

const (
	SetDisplayNameResponse_OK                   SetDisplayNameResponse_Result = 0
	SetDisplayNameResponse_INVALID_DISPLAY_NAME SetDisplayNameResponse_Result = 1
)

// Enum value maps for SetDisplayNameResponse_Result.
var (
	SetDisplayNameResponse_Result_name = map[int32]string{
		0: "OK",
		1: "INVALID_DISPLAY_NAME",
	}
	SetDisplayNameResponse_Result_value = map[string]int32{
		"OK":                   0,
		"INVALID_DISPLAY_NAME": 1,
	}
)

func (x SetDisplayNameResponse_Result) Enum() *SetDisplayNameResponse_Result {
	p := new(SetDisplayNameResponse_Result)
	*p = x
	return p
}

func (x SetDisplayNameResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetDisplayNameResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_v1_profile_service_proto_enumTypes[1].Descriptor()
}

func (SetDisplayNameResponse_Result) Type() protoreflect.EnumType {
	return &file_profile_v1_profile_service_proto_enumTypes[1]
}

func (x SetDisplayNameResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetDisplayNameResponse_Result.Descriptor instead.
func (SetDisplayNameResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{3, 0}
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *v1.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_profile_v1_profile_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetProfileRequest) GetUserId() *v1.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result GetProfileResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.profile.v1.GetProfileResponse_Result" json:"result,omitempty"`
	// UserProfile, if found.
	//
	// Some fields may or may not be set, depending on the scope of request
	// in the future.
	UserProfile *UserProfile `protobuf:"bytes,2,opt,name=user_profile,json=userProfile,proto3" json:"user_profile,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	mi := &file_profile_v1_profile_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetResult() GetProfileResponse_Result {
	if x != nil {
		return x.Result
	}
	return GetProfileResponse_OK
}

func (x *GetProfileResponse) GetUserProfile() *UserProfile {
	if x != nil {
		return x.UserProfile
	}
	return nil
}

type SetDisplayNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DisplayName is the new name to set.
	DisplayName string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Auth        *v1.Auth `protobuf:"bytes,10,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *SetDisplayNameRequest) Reset() {
	*x = SetDisplayNameRequest{}
	mi := &file_profile_v1_profile_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDisplayNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDisplayNameRequest) ProtoMessage() {}

func (x *SetDisplayNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDisplayNameRequest.ProtoReflect.Descriptor instead.
func (*SetDisplayNameRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{2}
}

func (x *SetDisplayNameRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SetDisplayNameRequest) GetAuth() *v1.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type SetDisplayNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result SetDisplayNameResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.profile.v1.SetDisplayNameResponse_Result" json:"result,omitempty"`
}

func (x *SetDisplayNameResponse) Reset() {
	*x = SetDisplayNameResponse{}
	mi := &file_profile_v1_profile_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDisplayNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDisplayNameResponse) ProtoMessage() {}

func (x *SetDisplayNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDisplayNameResponse.ProtoReflect.Descriptor instead.
func (*SetDisplayNameResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetDisplayNameResponse) GetResult() SetDisplayNameResponse_Result {
	if x != nil {
		return x.Result
	}
	return SetDisplayNameResponse_OK
}

var File_profile_v1_profile_service_proto protoreflect.FileDescriptor

var file_profile_v1_profile_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x66, 0x6c,
	0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x1f, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x22, 0x7d, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x40, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x53,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x2a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44,
	0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x32, 0xd3, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x66, 0x6c, 0x69,
	0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x87, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x69, 0x6e, 0x63, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x50, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0d,
	0x46, 0x43, 0x50, 0x42, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_profile_service_proto_rawDescOnce sync.Once
	file_profile_v1_profile_service_proto_rawDescData = file_profile_v1_profile_service_proto_rawDesc
)

func file_profile_v1_profile_service_proto_rawDescGZIP() []byte {
	file_profile_v1_profile_service_proto_rawDescOnce.Do(func() {
		file_profile_v1_profile_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_profile_service_proto_rawDescData)
	})
	return file_profile_v1_profile_service_proto_rawDescData
}

var file_profile_v1_profile_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_profile_v1_profile_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_profile_v1_profile_service_proto_goTypes = []any{
	(GetProfileResponse_Result)(0),     // 0: flipchat.profile.v1.GetProfileResponse.Result
	(SetDisplayNameResponse_Result)(0), // 1: flipchat.profile.v1.SetDisplayNameResponse.Result
	(*GetProfileRequest)(nil),          // 2: flipchat.profile.v1.GetProfileRequest
	(*GetProfileResponse)(nil),         // 3: flipchat.profile.v1.GetProfileResponse
	(*SetDisplayNameRequest)(nil),      // 4: flipchat.profile.v1.SetDisplayNameRequest
	(*SetDisplayNameResponse)(nil),     // 5: flipchat.profile.v1.SetDisplayNameResponse
	(*v1.UserId)(nil),                  // 6: flipchat.common.v1.UserId
	(*UserProfile)(nil),                // 7: flipchat.profile.v1.UserProfile
	(*v1.Auth)(nil),                    // 8: flipchat.common.v1.Auth
}
var file_profile_v1_profile_service_proto_depIdxs = []int32{
	6, // 0: flipchat.profile.v1.GetProfileRequest.user_id:type_name -> flipchat.common.v1.UserId
	0, // 1: flipchat.profile.v1.GetProfileResponse.result:type_name -> flipchat.profile.v1.GetProfileResponse.Result
	7, // 2: flipchat.profile.v1.GetProfileResponse.user_profile:type_name -> flipchat.profile.v1.UserProfile
	8, // 3: flipchat.profile.v1.SetDisplayNameRequest.auth:type_name -> flipchat.common.v1.Auth
	1, // 4: flipchat.profile.v1.SetDisplayNameResponse.result:type_name -> flipchat.profile.v1.SetDisplayNameResponse.Result
	2, // 5: flipchat.profile.v1.Profile.GetProfile:input_type -> flipchat.profile.v1.GetProfileRequest
	4, // 6: flipchat.profile.v1.Profile.SetDisplayName:input_type -> flipchat.profile.v1.SetDisplayNameRequest
	3, // 7: flipchat.profile.v1.Profile.GetProfile:output_type -> flipchat.profile.v1.GetProfileResponse
	5, // 8: flipchat.profile.v1.Profile.SetDisplayName:output_type -> flipchat.profile.v1.SetDisplayNameResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_profile_v1_profile_service_proto_init() }
func file_profile_v1_profile_service_proto_init() {
	if File_profile_v1_profile_service_proto != nil {
		return
	}
	file_profile_v1_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_v1_profile_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_v1_profile_service_proto_goTypes,
		DependencyIndexes: file_profile_v1_profile_service_proto_depIdxs,
		EnumInfos:         file_profile_v1_profile_service_proto_enumTypes,
		MessageInfos:      file_profile_v1_profile_service_proto_msgTypes,
	}.Build()
	File_profile_v1_profile_service_proto = out.File
	file_profile_v1_profile_service_proto_rawDesc = nil
	file_profile_v1_profile_service_proto_goTypes = nil
	file_profile_v1_profile_service_proto_depIdxs = nil
}
