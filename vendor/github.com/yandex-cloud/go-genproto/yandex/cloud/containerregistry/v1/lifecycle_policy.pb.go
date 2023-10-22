// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/containerregistry/v1/lifecycle_policy.proto

package containerregistry

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LifecyclePolicy_Status int32

const (
	LifecyclePolicy_STATUS_UNSPECIFIED LifecyclePolicy_Status = 0
	// Policy is active and regularly deletes Docker images according to the established rules.
	LifecyclePolicy_ACTIVE LifecyclePolicy_Status = 1
	// Policy is disabled and does not delete Docker images in the repository.
	// Policies in this status can be used for preparing and testing rules.
	LifecyclePolicy_DISABLED LifecyclePolicy_Status = 2
)

// Enum value maps for LifecyclePolicy_Status.
var (
	LifecyclePolicy_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ACTIVE",
		2: "DISABLED",
	}
	LifecyclePolicy_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ACTIVE":             1,
		"DISABLED":           2,
	}
)

func (x LifecyclePolicy_Status) Enum() *LifecyclePolicy_Status {
	p := new(LifecyclePolicy_Status)
	*p = x
	return p
}

func (x LifecyclePolicy_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LifecyclePolicy_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_enumTypes[0].Descriptor()
}

func (LifecyclePolicy_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_enumTypes[0]
}

func (x LifecyclePolicy_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LifecyclePolicy_Status.Descriptor instead.
func (LifecyclePolicy_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescGZIP(), []int{0, 0}
}

type LifecyclePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the lifecycle policy.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the lifecycle policy.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the repository that the lifecycle policy belongs to.
	// Required. The maximum string length in characters is 50.
	RepositoryId string `protobuf:"bytes,3,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	// Description of the lifecycle policy.
	// The maximum string length in characters is 256.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Status of lifecycle policy.
	Status LifecyclePolicy_Status `protobuf:"varint,5,opt,name=status,proto3,enum=yandex.cloud.containerregistry.v1.LifecyclePolicy_Status" json:"status,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The rules of lifecycle policy.
	Rules []*LifecycleRule `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *LifecyclePolicy) Reset() {
	*x = LifecyclePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifecyclePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecyclePolicy) ProtoMessage() {}

func (x *LifecyclePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecyclePolicy.ProtoReflect.Descriptor instead.
func (*LifecyclePolicy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescGZIP(), []int{0}
}

func (x *LifecyclePolicy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LifecyclePolicy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LifecyclePolicy) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

func (x *LifecyclePolicy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LifecyclePolicy) GetStatus() LifecyclePolicy_Status {
	if x != nil {
		return x.Status
	}
	return LifecyclePolicy_STATUS_UNSPECIFIED
}

func (x *LifecyclePolicy) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LifecyclePolicy) GetRules() []*LifecycleRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type LifecycleRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description of the lifecycle policy rule.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Period of time for automatic deletion.
	// Period must be a multiple of 24 hours.
	ExpirePeriod *durationpb.Duration `protobuf:"bytes,2,opt,name=expire_period,json=expirePeriod,proto3" json:"expire_period,omitempty"`
	// Tag for specifying a filter in the form of a regular expression.
	TagRegexp string `protobuf:"bytes,3,opt,name=tag_regexp,json=tagRegexp,proto3" json:"tag_regexp,omitempty"`
	// Tag for applying the rule to Docker images without tags.
	Untagged bool `protobuf:"varint,4,opt,name=untagged,proto3" json:"untagged,omitempty"`
	// Number of Docker images (falling under the specified filter by tags) that must be left, even if the expire_period has already expired.
	RetainedTop int64 `protobuf:"varint,5,opt,name=retained_top,json=retainedTop,proto3" json:"retained_top,omitempty"`
}

func (x *LifecycleRule) Reset() {
	*x = LifecycleRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifecycleRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleRule) ProtoMessage() {}

func (x *LifecycleRule) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleRule.ProtoReflect.Descriptor instead.
func (*LifecycleRule) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescGZIP(), []int{1}
}

func (x *LifecycleRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LifecycleRule) GetExpirePeriod() *durationpb.Duration {
	if x != nil {
		return x.ExpirePeriod
	}
	return nil
}

func (x *LifecycleRule) GetTagRegexp() string {
	if x != nil {
		return x.TagRegexp
	}
	return ""
}

func (x *LifecycleRule) GetUntagged() bool {
	if x != nil {
		return x.Untagged
	}
	return false
}

func (x *LifecycleRule) GetRetainedTop() int64 {
	if x != nil {
		return x.RetainedTop
	}
	return 0
}

var File_yandex_cloud_containerregistry_v1_lifecycle_policy_proto protoreflect.FileDescriptor

var file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDesc = []byte{
	0x0a, 0x38, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03,
	0x0a, 0x0f, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x46, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x3a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0xf9,
	0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x36,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a,
	0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x09, 0xfa, 0xc7, 0x31, 0x05, 0x3e, 0x3d, 0x32, 0x34, 0x68, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x28, 0x0a, 0x0a, 0x74, 0x61, 0x67, 0x5f,
	0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8,
	0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x36, 0x52, 0x09, 0x74, 0x61, 0x67, 0x52, 0x65, 0x67, 0x65,
	0x78, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x12, 0x2a,
	0x0a, 0x0c, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x3e, 0x3d, 0x30, 0x52, 0x0b, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x54, 0x6f, 0x70, 0x42, 0x80, 0x01, 0x0a, 0x25, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescData = file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDesc
)

func file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescData)
	})
	return file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_goTypes = []interface{}{
	(LifecyclePolicy_Status)(0),   // 0: yandex.cloud.containerregistry.v1.LifecyclePolicy.Status
	(*LifecyclePolicy)(nil),       // 1: yandex.cloud.containerregistry.v1.LifecyclePolicy
	(*LifecycleRule)(nil),         // 2: yandex.cloud.containerregistry.v1.LifecycleRule
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
}
var file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.containerregistry.v1.LifecyclePolicy.status:type_name -> yandex.cloud.containerregistry.v1.LifecyclePolicy.Status
	3, // 1: yandex.cloud.containerregistry.v1.LifecyclePolicy.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: yandex.cloud.containerregistry.v1.LifecyclePolicy.rules:type_name -> yandex.cloud.containerregistry.v1.LifecycleRule
	4, // 3: yandex.cloud.containerregistry.v1.LifecycleRule.expire_period:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_init() }
func file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_init() {
	if File_yandex_cloud_containerregistry_v1_lifecycle_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifecyclePolicy); i {
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
		file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifecycleRule); i {
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
			RawDescriptor: file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_lifecycle_policy_proto = out.File
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_rawDesc = nil
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_lifecycle_policy_proto_depIdxs = nil
}
