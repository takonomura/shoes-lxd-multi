// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/shoes-lxd-multi.proto

package shoeslxdmulti

import (
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

// type
type ResourceType int32

const (
	ResourceType_Unknown ResourceType = 0
	ResourceType_Nano    ResourceType = 1
	ResourceType_Micro   ResourceType = 2
	ResourceType_Small   ResourceType = 3
	ResourceType_Medium  ResourceType = 4
	ResourceType_Large   ResourceType = 5
	ResourceType_XLarge  ResourceType = 6
	ResourceType_XLarge2 ResourceType = 7
	ResourceType_XLarge3 ResourceType = 8
	ResourceType_XLarge4 ResourceType = 9
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "Unknown",
		1: "Nano",
		2: "Micro",
		3: "Small",
		4: "Medium",
		5: "Large",
		6: "XLarge",
		7: "XLarge2",
		8: "XLarge3",
		9: "XLarge4",
	}
	ResourceType_value = map[string]int32{
		"Unknown": 0,
		"Nano":    1,
		"Micro":   2,
		"Small":   3,
		"Medium":  4,
		"Large":   5,
		"XLarge":  6,
		"XLarge2": 7,
		"XLarge3": 8,
		"XLarge4": 9,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_shoes_lxd_multi_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_proto_shoes_lxd_multi_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_proto_shoes_lxd_multi_proto_rawDescGZIP(), []int{0}
}

// req / resp
type AddInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunnerName   string       `protobuf:"bytes,1,opt,name=runner_name,json=runnerName,proto3" json:"runner_name,omitempty"`
	SetupScript  string       `protobuf:"bytes,2,opt,name=setup_script,json=setupScript,proto3" json:"setup_script,omitempty"`
	ResourceType ResourceType `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=shoeslxdmulti.ResourceType" json:"resource_type,omitempty"`
	TargetHosts  []string     `protobuf:"bytes,4,rep,name=target_hosts,json=targetHosts,proto3" json:"target_hosts,omitempty"`
}

func (x *AddInstanceRequest) Reset() {
	*x = AddInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shoes_lxd_multi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInstanceRequest) ProtoMessage() {}

func (x *AddInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shoes_lxd_multi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInstanceRequest.ProtoReflect.Descriptor instead.
func (*AddInstanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_shoes_lxd_multi_proto_rawDescGZIP(), []int{0}
}

func (x *AddInstanceRequest) GetRunnerName() string {
	if x != nil {
		return x.RunnerName
	}
	return ""
}

func (x *AddInstanceRequest) GetSetupScript() string {
	if x != nil {
		return x.SetupScript
	}
	return ""
}

func (x *AddInstanceRequest) GetResourceType() ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return ResourceType_Unknown
}

func (x *AddInstanceRequest) GetTargetHosts() []string {
	if x != nil {
		return x.TargetHosts
	}
	return nil
}

type AddInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId   string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	ShoesType string `protobuf:"bytes,2,opt,name=shoes_type,json=shoesType,proto3" json:"shoes_type,omitempty"`
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *AddInstanceResponse) Reset() {
	*x = AddInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shoes_lxd_multi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInstanceResponse) ProtoMessage() {}

func (x *AddInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shoes_lxd_multi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInstanceResponse.ProtoReflect.Descriptor instead.
func (*AddInstanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_shoes_lxd_multi_proto_rawDescGZIP(), []int{1}
}

func (x *AddInstanceResponse) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *AddInstanceResponse) GetShoesType() string {
	if x != nil {
		return x.ShoesType
	}
	return ""
}

func (x *AddInstanceResponse) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

type DeleteInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CloudId     string   `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	TargetHosts []string `protobuf:"bytes,2,rep,name=target_hosts,json=targetHosts,proto3" json:"target_hosts,omitempty"`
}

func (x *DeleteInstanceRequest) Reset() {
	*x = DeleteInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shoes_lxd_multi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstanceRequest) ProtoMessage() {}

func (x *DeleteInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shoes_lxd_multi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstanceRequest.ProtoReflect.Descriptor instead.
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_shoes_lxd_multi_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteInstanceRequest) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *DeleteInstanceRequest) GetTargetHosts() []string {
	if x != nil {
		return x.TargetHosts
	}
	return nil
}

type DeleteInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInstanceResponse) Reset() {
	*x = DeleteInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shoes_lxd_multi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstanceResponse) ProtoMessage() {}

func (x *DeleteInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shoes_lxd_multi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstanceResponse.ProtoReflect.Descriptor instead.
func (*DeleteInstanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_shoes_lxd_multi_proto_rawDescGZIP(), []int{3}
}

var File_proto_shoes_lxd_multi_proto protoreflect.FileDescriptor

var file_proto_shoes_lxd_multi_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x2d, 0x6c, 0x78,
	0x64, 0x2d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73,
	0x68, 0x6f, 0x65, 0x73, 0x6c, 0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x22, 0xbd, 0x01, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b,
	0x2e, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x6c, 0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x6e, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x55, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x85, 0x01,
	0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x61, 0x6e, 0x6f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x10, 0x02,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x6d, 0x61, 0x6c, 0x6c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x61, 0x72, 0x67, 0x65,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x58, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x10, 0x06, 0x12, 0x0b,
	0x0a, 0x07, 0x58, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x32, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x58,
	0x4c, 0x61, 0x72, 0x67, 0x65, 0x33, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x58, 0x4c, 0x61, 0x72,
	0x67, 0x65, 0x34, 0x10, 0x09, 0x32, 0xc8, 0x01, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x65, 0x73, 0x4c,
	0x58, 0x44, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x56, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x6c, 0x78,
	0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x65,
	0x73, 0x6c, 0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x6c, 0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x6c,
	0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x68, 0x79, 0x77, 0x61, 0x69, 0x74, 0x61, 0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x2d, 0x6c, 0x78,
	0x64, 0x2d, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f,
	0x2f, 0x73, 0x68, 0x6f, 0x65, 0x73, 0x6c, 0x78, 0x64, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shoes_lxd_multi_proto_rawDescOnce sync.Once
	file_proto_shoes_lxd_multi_proto_rawDescData = file_proto_shoes_lxd_multi_proto_rawDesc
)

func file_proto_shoes_lxd_multi_proto_rawDescGZIP() []byte {
	file_proto_shoes_lxd_multi_proto_rawDescOnce.Do(func() {
		file_proto_shoes_lxd_multi_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shoes_lxd_multi_proto_rawDescData)
	})
	return file_proto_shoes_lxd_multi_proto_rawDescData
}

var file_proto_shoes_lxd_multi_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_shoes_lxd_multi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shoes_lxd_multi_proto_goTypes = []interface{}{
	(ResourceType)(0),              // 0: shoeslxdmulti.ResourceType
	(*AddInstanceRequest)(nil),     // 1: shoeslxdmulti.AddInstanceRequest
	(*AddInstanceResponse)(nil),    // 2: shoeslxdmulti.AddInstanceResponse
	(*DeleteInstanceRequest)(nil),  // 3: shoeslxdmulti.DeleteInstanceRequest
	(*DeleteInstanceResponse)(nil), // 4: shoeslxdmulti.DeleteInstanceResponse
}
var file_proto_shoes_lxd_multi_proto_depIdxs = []int32{
	0, // 0: shoeslxdmulti.AddInstanceRequest.resource_type:type_name -> shoeslxdmulti.ResourceType
	1, // 1: shoeslxdmulti.ShoesLXDMulti.AddInstance:input_type -> shoeslxdmulti.AddInstanceRequest
	3, // 2: shoeslxdmulti.ShoesLXDMulti.DeleteInstance:input_type -> shoeslxdmulti.DeleteInstanceRequest
	2, // 3: shoeslxdmulti.ShoesLXDMulti.AddInstance:output_type -> shoeslxdmulti.AddInstanceResponse
	4, // 4: shoeslxdmulti.ShoesLXDMulti.DeleteInstance:output_type -> shoeslxdmulti.DeleteInstanceResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_shoes_lxd_multi_proto_init() }
func file_proto_shoes_lxd_multi_proto_init() {
	if File_proto_shoes_lxd_multi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shoes_lxd_multi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInstanceRequest); i {
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
		file_proto_shoes_lxd_multi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInstanceResponse); i {
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
		file_proto_shoes_lxd_multi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstanceRequest); i {
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
		file_proto_shoes_lxd_multi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstanceResponse); i {
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
			RawDescriptor: file_proto_shoes_lxd_multi_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shoes_lxd_multi_proto_goTypes,
		DependencyIndexes: file_proto_shoes_lxd_multi_proto_depIdxs,
		EnumInfos:         file_proto_shoes_lxd_multi_proto_enumTypes,
		MessageInfos:      file_proto_shoes_lxd_multi_proto_msgTypes,
	}.Build()
	File_proto_shoes_lxd_multi_proto = out.File
	file_proto_shoes_lxd_multi_proto_rawDesc = nil
	file_proto_shoes_lxd_multi_proto_goTypes = nil
	file_proto_shoes_lxd_multi_proto_depIdxs = nil
}
