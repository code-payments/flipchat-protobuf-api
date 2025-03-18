// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: activity/v1/activity_feed_service.proto

package activitypb

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

type GetLatestNotificationsResponse_Result int32

const (
	GetLatestNotificationsResponse_OK     GetLatestNotificationsResponse_Result = 0
	GetLatestNotificationsResponse_DENIED GetLatestNotificationsResponse_Result = 1
)

// Enum value maps for GetLatestNotificationsResponse_Result.
var (
	GetLatestNotificationsResponse_Result_name = map[int32]string{
		0: "OK",
		1: "DENIED",
	}
	GetLatestNotificationsResponse_Result_value = map[string]int32{
		"OK":     0,
		"DENIED": 1,
	}
)

func (x GetLatestNotificationsResponse_Result) Enum() *GetLatestNotificationsResponse_Result {
	p := new(GetLatestNotificationsResponse_Result)
	*p = x
	return p
}

func (x GetLatestNotificationsResponse_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetLatestNotificationsResponse_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_v1_activity_feed_service_proto_enumTypes[0].Descriptor()
}

func (GetLatestNotificationsResponse_Result) Type() protoreflect.EnumType {
	return &file_activity_v1_activity_feed_service_proto_enumTypes[0]
}

func (x GetLatestNotificationsResponse_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetLatestNotificationsResponse_Result.Descriptor instead.
func (GetLatestNotificationsResponse_Result) EnumDescriptor() ([]byte, []int) {
	return file_activity_v1_activity_feed_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetLatestNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The activity feed to fetch notifications from
	Type ActivityFeedType `protobuf:"varint,1,opt,name=type,proto3,enum=flipchat.activity.v1.ActivityFeedType" json:"type,omitempty"`
	// Maximum number of notifications to return. If <= 0, the server default is used
	MaxItems int32    `protobuf:"varint,2,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	Auth     *v1.Auth `protobuf:"bytes,10,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *GetLatestNotificationsRequest) Reset() {
	*x = GetLatestNotificationsRequest{}
	mi := &file_activity_v1_activity_feed_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestNotificationsRequest) ProtoMessage() {}

func (x *GetLatestNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_activity_feed_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetLatestNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_activity_v1_activity_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLatestNotificationsRequest) GetType() ActivityFeedType {
	if x != nil {
		return x.Type
	}
	return ActivityFeedType_UNKNOWN
}

func (x *GetLatestNotificationsRequest) GetMaxItems() int32 {
	if x != nil {
		return x.MaxItems
	}
	return 0
}

func (x *GetLatestNotificationsRequest) GetAuth() *v1.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type GetLatestNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result        GetLatestNotificationsResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=flipchat.activity.v1.GetLatestNotificationsResponse_Result" json:"result,omitempty"`
	Notifications []*Notification                       `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *GetLatestNotificationsResponse) Reset() {
	*x = GetLatestNotificationsResponse{}
	mi := &file_activity_v1_activity_feed_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestNotificationsResponse) ProtoMessage() {}

func (x *GetLatestNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activity_v1_activity_feed_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetLatestNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_activity_v1_activity_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatestNotificationsResponse) GetResult() GetLatestNotificationsResponse_Result {
	if x != nil {
		return x.Result
	}
	return GetLatestNotificationsResponse_OK
}

func (x *GetLatestNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

var File_activity_v1_activity_feed_service_proto protoreflect.FileDescriptor

var file_activity_v1_activity_feed_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x66, 0x6c, 0x69, 0x70, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x66, 0x6c, 0x69, 0x70,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x18, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x1a, 0x03, 0x18, 0x80, 0x08, 0x52, 0x08,
	0x6d, 0x61, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0xe8, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0x80, 0x08, 0x52, 0x0d,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x1c, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x01, 0x32, 0x94, 0x01, 0x0a, 0x0c,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x64, 0x12, 0x83, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x66,
	0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x8b, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x69,
	0x6e, 0x63, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x52, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x70, 0x62, 0xa2,
	0x02, 0x0e, 0x46, 0x43, 0x50, 0x42, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_activity_v1_activity_feed_service_proto_rawDescOnce sync.Once
	file_activity_v1_activity_feed_service_proto_rawDescData = file_activity_v1_activity_feed_service_proto_rawDesc
)

func file_activity_v1_activity_feed_service_proto_rawDescGZIP() []byte {
	file_activity_v1_activity_feed_service_proto_rawDescOnce.Do(func() {
		file_activity_v1_activity_feed_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_v1_activity_feed_service_proto_rawDescData)
	})
	return file_activity_v1_activity_feed_service_proto_rawDescData
}

var file_activity_v1_activity_feed_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_activity_v1_activity_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_activity_v1_activity_feed_service_proto_goTypes = []any{
	(GetLatestNotificationsResponse_Result)(0), // 0: flipchat.activity.v1.GetLatestNotificationsResponse.Result
	(*GetLatestNotificationsRequest)(nil),      // 1: flipchat.activity.v1.GetLatestNotificationsRequest
	(*GetLatestNotificationsResponse)(nil),     // 2: flipchat.activity.v1.GetLatestNotificationsResponse
	(ActivityFeedType)(0),                      // 3: flipchat.activity.v1.ActivityFeedType
	(*v1.Auth)(nil),                            // 4: flipchat.common.v1.Auth
	(*Notification)(nil),                       // 5: flipchat.activity.v1.Notification
}
var file_activity_v1_activity_feed_service_proto_depIdxs = []int32{
	3, // 0: flipchat.activity.v1.GetLatestNotificationsRequest.type:type_name -> flipchat.activity.v1.ActivityFeedType
	4, // 1: flipchat.activity.v1.GetLatestNotificationsRequest.auth:type_name -> flipchat.common.v1.Auth
	0, // 2: flipchat.activity.v1.GetLatestNotificationsResponse.result:type_name -> flipchat.activity.v1.GetLatestNotificationsResponse.Result
	5, // 3: flipchat.activity.v1.GetLatestNotificationsResponse.notifications:type_name -> flipchat.activity.v1.Notification
	1, // 4: flipchat.activity.v1.ActivityFeed.GetLatestNotifications:input_type -> flipchat.activity.v1.GetLatestNotificationsRequest
	2, // 5: flipchat.activity.v1.ActivityFeed.GetLatestNotifications:output_type -> flipchat.activity.v1.GetLatestNotificationsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_activity_v1_activity_feed_service_proto_init() }
func file_activity_v1_activity_feed_service_proto_init() {
	if File_activity_v1_activity_feed_service_proto != nil {
		return
	}
	file_activity_v1_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_activity_v1_activity_feed_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_v1_activity_feed_service_proto_goTypes,
		DependencyIndexes: file_activity_v1_activity_feed_service_proto_depIdxs,
		EnumInfos:         file_activity_v1_activity_feed_service_proto_enumTypes,
		MessageInfos:      file_activity_v1_activity_feed_service_proto_msgTypes,
	}.Build()
	File_activity_v1_activity_feed_service_proto = out.File
	file_activity_v1_activity_feed_service_proto_rawDesc = nil
	file_activity_v1_activity_feed_service_proto_goTypes = nil
	file_activity_v1_activity_feed_service_proto_depIdxs = nil
}
