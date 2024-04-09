// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: publish.proto

package publish

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	feed "pb/feed"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId  int64  `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"` // 用户id
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                       // 视频数据
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                     // 视频标题
	CoverUrl string `protobuf:"bytes,4,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`               // 视频封面
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`               // 视频分类
	Label    string `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`                     //视频标签
}

func (x *CreateVideoRequest) Reset() {
	*x = CreateVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoRequest) ProtoMessage() {}

func (x *CreateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoRequest.ProtoReflect.Descriptor instead.
func (*CreateVideoRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVideoRequest) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *CreateVideoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateVideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *CreateVideoRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateVideoRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"` // 状态码，0-成功，其他值-失败
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`    // 返回状态描述
}

func (x *CreateVideoResponse) Reset() {
	*x = CreateVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoResponse) ProtoMessage() {}

func (x *CreateVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoResponse.ProtoReflect.Descriptor instead.
func (*CreateVideoResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{1}
}

func (x *CreateVideoResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateVideoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 被请求查询的用户id
}

func (x *ListVideoRequest) Reset() {
	*x = ListVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVideoRequest) ProtoMessage() {}

func (x *ListVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVideoRequest.ProtoReflect.Descriptor instead.
func (*ListVideoRequest) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{2}
}

func (x *ListVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ListVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int64         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                           // 状态码，0-成功，其他值-失败
	Msg       string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                              // 返回状态描述
	VideoList []*feed.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"` // 视频列表
}

func (x *ListVideoResponse) Reset() {
	*x = ListVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVideoResponse) ProtoMessage() {}

func (x *ListVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVideoResponse.ProtoReflect.Descriptor instead.
func (*ListVideoResponse) Descriptor() ([]byte, []int) {
	return file_publish_proto_rawDescGZIP(), []int{3}
}

func (x *ListVideoResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListVideoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListVideoResponse) GetVideoList() []*feed.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_publish_proto protoreflect.FileDescriptor

var file_publish_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x1a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3b,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2b, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32,
	0xa2, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x1b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x19, 0x2e, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x62, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publish_proto_rawDescOnce sync.Once
	file_publish_proto_rawDescData = file_publish_proto_rawDesc
)

func file_publish_proto_rawDescGZIP() []byte {
	file_publish_proto_rawDescOnce.Do(func() {
		file_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_publish_proto_rawDescData)
	})
	return file_publish_proto_rawDescData
}

var file_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_publish_proto_goTypes = []interface{}{
	(*CreateVideoRequest)(nil),  // 0: publish.CreateVideoRequest
	(*CreateVideoResponse)(nil), // 1: publish.CreateVideoResponse
	(*ListVideoRequest)(nil),    // 2: publish.ListVideoRequest
	(*ListVideoResponse)(nil),   // 3: publish.ListVideoResponse
	(*feed.Video)(nil),          // 4: feed.Video
}
var file_publish_proto_depIdxs = []int32{
	4, // 0: publish.ListVideoResponse.video_list:type_name -> feed.Video
	0, // 1: publish.PublishService.CreateVideo:input_type -> publish.CreateVideoRequest
	2, // 2: publish.PublishService.ListVideo:input_type -> publish.ListVideoRequest
	1, // 3: publish.PublishService.CreateVideo:output_type -> publish.CreateVideoResponse
	3, // 4: publish.PublishService.ListVideo:output_type -> publish.ListVideoResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_publish_proto_init() }
func file_publish_proto_init() {
	if File_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publish_proto_goTypes,
		DependencyIndexes: file_publish_proto_depIdxs,
		MessageInfos:      file_publish_proto_msgTypes,
	}.Build()
	File_publish_proto = out.File
	file_publish_proto_rawDesc = nil
	file_publish_proto_goTypes = nil
	file_publish_proto_depIdxs = nil
}
