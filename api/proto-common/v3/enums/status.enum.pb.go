// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/proto-common/v3/enums/status.enum.proto

package common_enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusCodeEnum_StatusCode int32

const (
	StatusCodeEnum_STATUS_CODE_UNSPECIFIED StatusCodeEnum_StatusCode = 0
)

// Enum value maps for StatusCodeEnum_StatusCode.
var (
	StatusCodeEnum_StatusCode_name = map[int32]string{
		0: "STATUS_CODE_UNSPECIFIED",
	}
	StatusCodeEnum_StatusCode_value = map[string]int32{
		"STATUS_CODE_UNSPECIFIED": 0,
	}
)

func (x StatusCodeEnum_StatusCode) Enum() *StatusCodeEnum_StatusCode {
	p := new(StatusCodeEnum_StatusCode)
	*p = x
	return p
}

func (x StatusCodeEnum_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCodeEnum_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_common_v3_enums_status_enum_proto_enumTypes[0].Descriptor()
}

func (StatusCodeEnum_StatusCode) Type() protoreflect.EnumType {
	return &file_api_proto_common_v3_enums_status_enum_proto_enumTypes[0]
}

func (x StatusCodeEnum_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCodeEnum_StatusCode.Descriptor instead.
func (StatusCodeEnum_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_common_v3_enums_status_enum_proto_rawDescGZIP(), []int{0, 0}
}

type StatusCodeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusCodeEnum) Reset() {
	*x = StatusCodeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_common_v3_enums_status_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCodeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCodeEnum) ProtoMessage() {}

func (x *StatusCodeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_common_v3_enums_status_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCodeEnum.ProtoReflect.Descriptor instead.
func (*StatusCodeEnum) Descriptor() ([]byte, []int) {
	return file_api_proto_common_v3_enums_status_enum_proto_rawDescGZIP(), []int{0}
}

var file_api_proto_common_v3_enums_status_enum_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         99999999,
		Name:          "proto_common.common.status.description",
		Tag:           "bytes,99999999,opt,name=description",
		Filename:      "api/proto-common/v3/enums/status.enum.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string description = 99999999;
	E_Description = &file_api_proto_common_v3_enums_status_enum_proto_extTypes[0]
)

var File_api_proto_common_v3_enums_status_enum_proto protoreflect.FileDescriptor

var file_api_proto_common_v3_enums_status_enum_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x29, 0x0a,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x3a, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xff, 0xc1, 0xd7, 0x2f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69, 0x6e, 0x75,
	0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d, 0x71, 0x75,
	0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_common_v3_enums_status_enum_proto_rawDescOnce sync.Once
	file_api_proto_common_v3_enums_status_enum_proto_rawDescData = file_api_proto_common_v3_enums_status_enum_proto_rawDesc
)

func file_api_proto_common_v3_enums_status_enum_proto_rawDescGZIP() []byte {
	file_api_proto_common_v3_enums_status_enum_proto_rawDescOnce.Do(func() {
		file_api_proto_common_v3_enums_status_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_common_v3_enums_status_enum_proto_rawDescData)
	})
	return file_api_proto_common_v3_enums_status_enum_proto_rawDescData
}

var file_api_proto_common_v3_enums_status_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_common_v3_enums_status_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_common_v3_enums_status_enum_proto_goTypes = []interface{}{
	(StatusCodeEnum_StatusCode)(0),        // 0: proto_common.common.status.StatusCodeEnum.StatusCode
	(*StatusCodeEnum)(nil),                // 1: proto_common.common.status.StatusCodeEnum
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_api_proto_common_v3_enums_status_enum_proto_depIdxs = []int32{
	2, // 0: proto_common.common.status.description:extendee -> google.protobuf.EnumValueOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_common_v3_enums_status_enum_proto_init() }
func file_api_proto_common_v3_enums_status_enum_proto_init() {
	if File_api_proto_common_v3_enums_status_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_common_v3_enums_status_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCodeEnum); i {
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
			RawDescriptor: file_api_proto_common_v3_enums_status_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_common_v3_enums_status_enum_proto_goTypes,
		DependencyIndexes: file_api_proto_common_v3_enums_status_enum_proto_depIdxs,
		EnumInfos:         file_api_proto_common_v3_enums_status_enum_proto_enumTypes,
		MessageInfos:      file_api_proto_common_v3_enums_status_enum_proto_msgTypes,
		ExtensionInfos:    file_api_proto_common_v3_enums_status_enum_proto_extTypes,
	}.Build()
	File_api_proto_common_v3_enums_status_enum_proto = out.File
	file_api_proto_common_v3_enums_status_enum_proto_rawDesc = nil
	file_api_proto_common_v3_enums_status_enum_proto_goTypes = nil
	file_api_proto_common_v3_enums_status_enum_proto_depIdxs = nil
}
