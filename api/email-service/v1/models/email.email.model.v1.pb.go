// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/email-service/v1/models/email.email.model.v1.proto

package emailmodelv1

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_email_service_v1_models_email_email_model_v1_proto protoreflect.FileDescriptor

var file_api_email_service_v1_models_email_email_model_v1_proto_rawDesc = []byte{
	0x0a, 0x36, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x6c,
	0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69, 0x6e, 0x75, 0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e,
	0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d, 0x71, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x3b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_email_service_v1_models_email_email_model_v1_proto_goTypes = []interface{}{}
var file_api_email_service_v1_models_email_email_model_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_email_service_v1_models_email_email_model_v1_proto_init() }
func file_api_email_service_v1_models_email_email_model_v1_proto_init() {
	if File_api_email_service_v1_models_email_email_model_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_email_service_v1_models_email_email_model_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_email_service_v1_models_email_email_model_v1_proto_goTypes,
		DependencyIndexes: file_api_email_service_v1_models_email_email_model_v1_proto_depIdxs,
	}.Build()
	File_api_email_service_v1_models_email_email_model_v1_proto = out.File
	file_api_email_service_v1_models_email_email_model_v1_proto_rawDesc = nil
	file_api_email_service_v1_models_email_email_model_v1_proto_goTypes = nil
	file_api_email_service_v1_models_email_email_model_v1_proto_depIdxs = nil
}
