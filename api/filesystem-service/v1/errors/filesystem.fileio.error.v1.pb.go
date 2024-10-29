// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/filesystem-service/v1/errors/filesystem.fileio.error.v1.proto

package filesystemerrorv1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ERROR .
type FILEIO_ERROR int32

const (
	FILEIO_ERROR_OK               FILEIO_ERROR = 0
	FILEIO_ERROR_OBJECT_NOT_EXIST FILEIO_ERROR = 100
	FILEIO_ERROR_OBJECT_EXIST     FILEIO_ERROR = 101
	FILEIO_ERROR_ACCESS_DENIED    FILEIO_ERROR = 102
	FILEIO_ERROR_INTERNAL         FILEIO_ERROR = 103
)

// Enum value maps for FILEIO_ERROR.
var (
	FILEIO_ERROR_name = map[int32]string{
		0:   "OK",
		100: "OBJECT_NOT_EXIST",
		101: "OBJECT_EXIST",
		102: "ACCESS_DENIED",
		103: "INTERNAL",
	}
	FILEIO_ERROR_value = map[string]int32{
		"OK":               0,
		"OBJECT_NOT_EXIST": 100,
		"OBJECT_EXIST":     101,
		"ACCESS_DENIED":    102,
		"INTERNAL":         103,
	}
)

func (x FILEIO_ERROR) Enum() *FILEIO_ERROR {
	p := new(FILEIO_ERROR)
	*p = x
	return p
}

func (x FILEIO_ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FILEIO_ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_enumTypes[0].Descriptor()
}

func (FILEIO_ERROR) Type() protoreflect.EnumType {
	return &file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_enumTypes[0]
}

func (x FILEIO_ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FILEIO_ERROR.Descriptor instead.
func (FILEIO_ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto protoreflect.FileDescriptor

var file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDesc = []byte{
	0x0a, 0x41, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x69, 0x6f, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x24, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x69, 0x6f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x83,
	0x01, 0x0a, 0x0c, 0x46, 0x49, 0x4c, 0x45, 0x49, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12,
	0x0c, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x1a, 0x0a,
	0x10, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53,
	0x54, 0x10, 0x64, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x16, 0x0a, 0x0c, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x65, 0x1a, 0x04, 0xa8, 0x45, 0xc8,
	0x01, 0x12, 0x17, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49,
	0x45, 0x44, 0x10, 0x66, 0x1a, 0x04, 0xa8, 0x45, 0xc8, 0x01, 0x12, 0x12, 0x0a, 0x08, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x67, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04,
	0xa0, 0x45, 0xf4, 0x03, 0x42, 0x61, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c,
	0x61, 0x69, 0x6e, 0x75, 0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65,
	0x6e, 0x2d, 0x71, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescOnce sync.Once
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescData = file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDesc
)

func file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescGZIP() []byte {
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescOnce.Do(func() {
		file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescData)
	})
	return file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDescData
}

var file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_goTypes = []interface{}{
	(FILEIO_ERROR)(0), // 0: api.filesystem.service.fileioerrorv1.FILEIO_ERROR
}
var file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_init() }
func file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_init() {
	if File_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_goTypes,
		DependencyIndexes: file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_depIdxs,
		EnumInfos:         file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_enumTypes,
	}.Build()
	File_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto = out.File
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_rawDesc = nil
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_goTypes = nil
	file_api_filesystem_service_v1_errors_filesystem_fileio_error_v1_proto_depIdxs = nil
}
