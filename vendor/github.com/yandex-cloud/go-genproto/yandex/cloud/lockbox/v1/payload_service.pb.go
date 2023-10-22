// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/lockbox/v1/payload_service.proto

package lockbox

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPayloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the secret.
	SecretId string `protobuf:"bytes,1,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	// Optional ID of the version.
	VersionId string `protobuf:"bytes,2,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
}

func (x *GetPayloadRequest) Reset() {
	*x = GetPayloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayloadRequest) ProtoMessage() {}

func (x *GetPayloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayloadRequest.ProtoReflect.Descriptor instead.
func (*GetPayloadRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPayloadRequest) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *GetPayloadRequest) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

var File_yandex_cloud_lockbox_v1_payload_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f,
	0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0x97, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x6b,
	0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData = file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc
)

func file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData)
	})
	return file_yandex_cloud_lockbox_v1_payload_service_proto_rawDescData
}

var file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes = []interface{}{
	(*GetPayloadRequest)(nil), // 0: yandex.cloud.lockbox.v1.GetPayloadRequest
	(*Payload)(nil),           // 1: yandex.cloud.lockbox.v1.Payload
}
var file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.lockbox.v1.PayloadService.Get:input_type -> yandex.cloud.lockbox.v1.GetPayloadRequest
	1, // 1: yandex.cloud.lockbox.v1.PayloadService.Get:output_type -> yandex.cloud.lockbox.v1.Payload
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_lockbox_v1_payload_service_proto_init() }
func file_yandex_cloud_lockbox_v1_payload_service_proto_init() {
	if File_yandex_cloud_lockbox_v1_payload_service_proto != nil {
		return
	}
	file_yandex_cloud_lockbox_v1_payload_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayloadRequest); i {
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
			RawDescriptor: file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_lockbox_v1_payload_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_lockbox_v1_payload_service_proto = out.File
	file_yandex_cloud_lockbox_v1_payload_service_proto_rawDesc = nil
	file_yandex_cloud_lockbox_v1_payload_service_proto_goTypes = nil
	file_yandex_cloud_lockbox_v1_payload_service_proto_depIdxs = nil
}
