// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/iam-service/v1/services/iam.auth.service.v1.proto

package iamservicev1

import (
	reflect "reflect"

	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/iam-service/v1/resources"
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

var File_api_iam_service_v1_services_iam_auth_service_v1_proto protoreflect.FileDescriptor

var file_api_iam_service_v1_services_iam_auth_service_v1_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x32, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x02, 0x0a, 0x0e, 0x49, 0x41, 0x4d, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c,
	0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x08,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x12, 0xfa, 0xe9, 0x30,
	0x0e, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x42,
	0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69, 0x6e, 0x75, 0x6f,
	0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d, 0x71, 0x75, 0x61,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x69, 0x61, 0x6d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_iam_service_v1_services_iam_auth_service_v1_proto_goTypes = []interface{}{
	(*resources.CheckServiceAuthRequest)(nil),  // 0: api.iam.service.iamv1.CheckServiceAuthRequest
	(*resources.CheckUserAuthRequest)(nil),     // 1: api.iam.service.iamv1.CheckUserAuthRequest
	(*resources.ValidateRequest)(nil),          // 2: api.iam.service.iamv1.ValidateRequest
	(*resources.CheckServiceAuthResponse)(nil), // 3: api.iam.service.iamv1.CheckServiceAuthResponse
	(*resources.CheckUserAuthResponse)(nil),    // 4: api.iam.service.iamv1.CheckUserAuthResponse
	(*resources.ValidateResponse)(nil),         // 5: api.iam.service.iamv1.ValidateResponse
}
var file_api_iam_service_v1_services_iam_auth_service_v1_proto_depIdxs = []int32{
	0, // 0: api.iam.service.iamservicev1.IAMAuthService.CheckServiceAuth:input_type -> api.iam.service.iamv1.CheckServiceAuthRequest
	1, // 1: api.iam.service.iamservicev1.IAMAuthService.CheckUserAuth:input_type -> api.iam.service.iamv1.CheckUserAuthRequest
	2, // 2: api.iam.service.iamservicev1.IAMAuthService.Validate:input_type -> api.iam.service.iamv1.ValidateRequest
	3, // 3: api.iam.service.iamservicev1.IAMAuthService.CheckServiceAuth:output_type -> api.iam.service.iamv1.CheckServiceAuthResponse
	4, // 4: api.iam.service.iamservicev1.IAMAuthService.CheckUserAuth:output_type -> api.iam.service.iamv1.CheckUserAuthResponse
	5, // 5: api.iam.service.iamservicev1.IAMAuthService.Validate:output_type -> api.iam.service.iamv1.ValidateResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_iam_service_v1_services_iam_auth_service_v1_proto_init() }
func file_api_iam_service_v1_services_iam_auth_service_v1_proto_init() {
	if File_api_iam_service_v1_services_iam_auth_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_iam_service_v1_services_iam_auth_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_iam_service_v1_services_iam_auth_service_v1_proto_goTypes,
		DependencyIndexes: file_api_iam_service_v1_services_iam_auth_service_v1_proto_depIdxs,
	}.Build()
	File_api_iam_service_v1_services_iam_auth_service_v1_proto = out.File
	file_api_iam_service_v1_services_iam_auth_service_v1_proto_rawDesc = nil
	file_api_iam_service_v1_services_iam_auth_service_v1_proto_goTypes = nil
	file_api_iam_service_v1_services_iam_auth_service_v1_proto_depIdxs = nil
}
