// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/proto-common/v3/errors/common.error.v3.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

// ERROR .
type ERROR int32

const (
	ERROR_ERROR_UNSPECIFIED    ERROR = 0
	ERROR_BAD_REQUEST          ERROR = 10000001  // 参数校验错误
	ERROR_INTERNAL             ERROR = 10000002  // 未处理错误
	ERROR_UNAUTHORIZED         ERROR = 10010003  // 认证错误
	ERROR_CONTENT_MISSING      ERROR = 100000001 // oh missing
	ERROR_CONTENT_ERROR        ERROR = 100000002 // oh error
	ERROR_QUERY_DEST           ERROR = 10020001  // query db error from sql-fx plugin
	ERROR_QUERY_NOT_IMPLEMENT  ERROR = 10020002  // query db's injection is not implemented
	ERROR_QUERY_ARGS_NOT_FOUND ERROR = 10020003  // query args not found
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:         "ERROR_UNSPECIFIED",
		10000001:  "BAD_REQUEST",
		10000002:  "INTERNAL",
		10010003:  "UNAUTHORIZED",
		100000001: "CONTENT_MISSING",
		100000002: "CONTENT_ERROR",
		10020001:  "QUERY_DEST",
		10020002:  "QUERY_NOT_IMPLEMENT",
		10020003:  "QUERY_ARGS_NOT_FOUND",
	}
	ERROR_value = map[string]int32{
		"ERROR_UNSPECIFIED":    0,
		"BAD_REQUEST":          10000001,
		"INTERNAL":             10000002,
		"UNAUTHORIZED":         10010003,
		"CONTENT_MISSING":      100000001,
		"CONTENT_ERROR":        100000002,
		"QUERY_DEST":           10020001,
		"QUERY_NOT_IMPLEMENT":  10020002,
		"QUERY_ARGS_NOT_FOUND": 10020003,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_common_v3_errors_common_error_v3_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_proto_common_v3_errors_common_error_v3_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_common_v3_errors_common_error_v3_proto_rawDescGZIP(), []int{0}
}

var File_api_proto_common_v3_errors_common_error_v3_proto protoreflect.FileDescriptor

var file_api_proto_common_v3_errors_common_error_v3_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x88, 0x02, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x81, 0xad, 0xe2, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x15, 0x0a,
	0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x82, 0xad, 0xe2, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x19, 0x0a, 0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52,
	0x49, 0x5a, 0x45, 0x44, 0x10, 0x93, 0xfb, 0xe2, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12,
	0x1c, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x81, 0xc2, 0xd7, 0x2f, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x1a, 0x0a,
	0x0d, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x82,
	0xc2, 0xd7, 0x2f, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x17, 0x0a, 0x0a, 0x51, 0x55, 0x45,
	0x52, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x54, 0x10, 0xa1, 0xc9, 0xe3, 0x04, 0x1a, 0x04, 0xa8, 0x45,
	0xf4, 0x03, 0x12, 0x20, 0x0a, 0x13, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0xa2, 0xc9, 0xe3, 0x04, 0x1a, 0x04,
	0xa8, 0x45, 0xf4, 0x03, 0x12, 0x21, 0x0a, 0x14, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x41, 0x52,
	0x47, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xa3, 0xc9, 0xe3,
	0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xc8, 0x01, 0x42, 0x50, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69, 0x6e, 0x75, 0x6f, 0x6e, 0x69,
	0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d, 0x71, 0x75, 0x61, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_common_v3_errors_common_error_v3_proto_rawDescOnce sync.Once
	file_api_proto_common_v3_errors_common_error_v3_proto_rawDescData = file_api_proto_common_v3_errors_common_error_v3_proto_rawDesc
)

func file_api_proto_common_v3_errors_common_error_v3_proto_rawDescGZIP() []byte {
	file_api_proto_common_v3_errors_common_error_v3_proto_rawDescOnce.Do(func() {
		file_api_proto_common_v3_errors_common_error_v3_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_common_v3_errors_common_error_v3_proto_rawDescData)
	})
	return file_api_proto_common_v3_errors_common_error_v3_proto_rawDescData
}

var file_api_proto_common_v3_errors_common_error_v3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_common_v3_errors_common_error_v3_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: proto_common.common.error.ERROR
}
var file_api_proto_common_v3_errors_common_error_v3_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_common_v3_errors_common_error_v3_proto_init() }
func file_api_proto_common_v3_errors_common_error_v3_proto_init() {
	if File_api_proto_common_v3_errors_common_error_v3_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_common_v3_errors_common_error_v3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_common_v3_errors_common_error_v3_proto_goTypes,
		DependencyIndexes: file_api_proto_common_v3_errors_common_error_v3_proto_depIdxs,
		EnumInfos:         file_api_proto_common_v3_errors_common_error_v3_proto_enumTypes,
	}.Build()
	File_api_proto_common_v3_errors_common_error_v3_proto = out.File
	file_api_proto_common_v3_errors_common_error_v3_proto_rawDesc = nil
	file_api_proto_common_v3_errors_common_error_v3_proto_goTypes = nil
	file_api_proto_common_v3_errors_common_error_v3_proto_depIdxs = nil
}
