// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/filesystem-service/v1/services/filesystem.filewrite.service.v1.proto

package filesystemservicev1

import (
	reflect "reflect"

	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/filesystem-service/v1/resources"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/api/proto-common/v3/services"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto protoreflect.FileDescriptor

var file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_rawDesc = []byte{
	0x0a, 0x48, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x32, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4a, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x47, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xac, 0x02, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x35, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x76, 0x31,
	0x2e, 0x46, 0x53, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x2e, 0x46, 0x53, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x79, 0x0a, 0x09, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x76, 0x31, 0x2e, 0x46, 0x53, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x2e, 0x46, 0x53, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x1a, 0x18, 0xfa, 0xe9, 0x30, 0x14, 0x66, 0x69, 0x6c, 0x65, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x42,
	0x65, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69, 0x6e, 0x75, 0x6f,
	0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d, 0x71, 0x75, 0x61,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_goTypes = []interface{}{
	(*resources.FSWriteFileStream)(nil),  // 0: api.filesystem.service.filewritev1.FSWriteFileStream
	(*resources.FSWriteFileRequest)(nil), // 1: api.filesystem.service.filewritev1.FSWriteFileRequest
	(*resources.FSCommonReply)(nil),      // 2: api.filesystem.service.filecommonv1.FSCommonReply
}
var file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_depIdxs = []int32{
	0, // 0: api.filesystem.service.filewriteservicev1.FileWriteServiceV1.WriteFileStream:input_type -> api.filesystem.service.filewritev1.FSWriteFileStream
	1, // 1: api.filesystem.service.filewriteservicev1.FileWriteServiceV1.WriteFile:input_type -> api.filesystem.service.filewritev1.FSWriteFileRequest
	2, // 2: api.filesystem.service.filewriteservicev1.FileWriteServiceV1.WriteFileStream:output_type -> api.filesystem.service.filecommonv1.FSCommonReply
	2, // 3: api.filesystem.service.filewriteservicev1.FileWriteServiceV1.WriteFile:output_type -> api.filesystem.service.filecommonv1.FSCommonReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_init() }
func file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_init() {
	if File_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_goTypes,
		DependencyIndexes: file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_depIdxs,
	}.Build()
	File_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto = out.File
	file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_rawDesc = nil
	file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_goTypes = nil
	file_api_filesystem_service_v1_services_filesystem_filewrite_service_v1_proto_depIdxs = nil
}
