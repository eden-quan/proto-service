// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/iam-service/v1/services/iam.user.service.v1.proto

package iamservicev1

import (
	reflect "reflect"

	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/api/common/v3/services"
	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/iam-service/v1/resources"
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

var File_api_iam_service_v1_services_iam_user_service_v1_proto protoreflect.FileDescriptor

var file_api_iam_service_v1_services_iam_user_service_v1_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61,
	0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x2c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x93,
	0x08, 0x0a, 0x0e, 0x49, 0x41, 0x4d, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x2c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x71, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x92, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x41, 0x6e, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x69, 0x61,
	0x6d, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x1a, 0x12, 0xfa, 0xe9, 0x30, 0x0e, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c,
	0x61, 0x69, 0x6e, 0x75, 0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65,
	0x6e, 0x2d, 0x71, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3b, 0x69, 0x61, 0x6d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_iam_service_v1_services_iam_user_service_v1_proto_goTypes = []interface{}{
	(*resources.UserUserLoginRequest)(nil),                // 0: api.iam.service.iamv1.UserUserLoginRequest
	(*resources.UserUserLogoutRequest)(nil),               // 1: api.iam.service.iamv1.UserUserLogoutRequest
	(*resources.UserUserRegisterRequest)(nil),             // 2: api.iam.service.iamv1.UserUserRegisterRequest
	(*resources.UserCheckTokenRequest)(nil),               // 3: api.iam.service.iamv1.UserCheckTokenRequest
	(*resources.UserCheckUserRequest)(nil),                // 4: api.iam.service.iamv1.UserCheckUserRequest
	(*resources.UserRefreshTokenRequest)(nil),             // 5: api.iam.service.iamv1.UserRefreshTokenRequest
	(*resources.UserUpdateCompanyIdAndLabelRequest)(nil),  // 6: api.iam.service.iamv1.UserUpdateCompanyIdAndLabelRequest
	(*resources.UserDeleteUserRequest)(nil),               // 7: api.iam.service.iamv1.UserDeleteUserRequest
	(*resources.UserCheckLastLoginRequest)(nil),           // 8: api.iam.service.iamv1.UserCheckLastLoginRequest
	(*resources.UserUser)(nil),                            // 9: api.iam.service.iamv1.UserUser
	(*resources.UserUserLogoutReply)(nil),                 // 10: api.iam.service.iamv1.UserUserLogoutReply
	(*resources.UserCheckTokenResponse)(nil),              // 11: api.iam.service.iamv1.UserCheckTokenResponse
	(*resources.UserCheckUserResponse)(nil),               // 12: api.iam.service.iamv1.UserCheckUserResponse
	(*resources.UserRefreshTokenResponse)(nil),            // 13: api.iam.service.iamv1.UserRefreshTokenResponse
	(*resources.UserUpdateCompanyIdAndLabelResponse)(nil), // 14: api.iam.service.iamv1.UserUpdateCompanyIdAndLabelResponse
	(*resources.UserDeleteUserResponse)(nil),              // 15: api.iam.service.iamv1.UserDeleteUserResponse
	(*resources.UserCheckLastLoginResponse)(nil),          // 16: api.iam.service.iamv1.UserCheckLastLoginResponse
}
var file_api_iam_service_v1_services_iam_user_service_v1_proto_depIdxs = []int32{
	0,  // 0: api.iam.service.iamservicev1.IAMUserService.UserLogin:input_type -> api.iam.service.iamv1.UserUserLoginRequest
	1,  // 1: api.iam.service.iamservicev1.IAMUserService.UserLogout:input_type -> api.iam.service.iamv1.UserUserLogoutRequest
	2,  // 2: api.iam.service.iamservicev1.IAMUserService.UserRegister:input_type -> api.iam.service.iamv1.UserUserRegisterRequest
	3,  // 3: api.iam.service.iamservicev1.IAMUserService.CheckToken:input_type -> api.iam.service.iamv1.UserCheckTokenRequest
	4,  // 4: api.iam.service.iamservicev1.IAMUserService.CheckUser:input_type -> api.iam.service.iamv1.UserCheckUserRequest
	5,  // 5: api.iam.service.iamservicev1.IAMUserService.RefreshToken:input_type -> api.iam.service.iamv1.UserRefreshTokenRequest
	6,  // 6: api.iam.service.iamservicev1.IAMUserService.UpdateCompanyIdAndLabel:input_type -> api.iam.service.iamv1.UserUpdateCompanyIdAndLabelRequest
	7,  // 7: api.iam.service.iamservicev1.IAMUserService.DeleteUser:input_type -> api.iam.service.iamv1.UserDeleteUserRequest
	8,  // 8: api.iam.service.iamservicev1.IAMUserService.CheckLastLogin:input_type -> api.iam.service.iamv1.UserCheckLastLoginRequest
	9,  // 9: api.iam.service.iamservicev1.IAMUserService.UserLogin:output_type -> api.iam.service.iamv1.UserUser
	10, // 10: api.iam.service.iamservicev1.IAMUserService.UserLogout:output_type -> api.iam.service.iamv1.UserUserLogoutReply
	9,  // 11: api.iam.service.iamservicev1.IAMUserService.UserRegister:output_type -> api.iam.service.iamv1.UserUser
	11, // 12: api.iam.service.iamservicev1.IAMUserService.CheckToken:output_type -> api.iam.service.iamv1.UserCheckTokenResponse
	12, // 13: api.iam.service.iamservicev1.IAMUserService.CheckUser:output_type -> api.iam.service.iamv1.UserCheckUserResponse
	13, // 14: api.iam.service.iamservicev1.IAMUserService.RefreshToken:output_type -> api.iam.service.iamv1.UserRefreshTokenResponse
	14, // 15: api.iam.service.iamservicev1.IAMUserService.UpdateCompanyIdAndLabel:output_type -> api.iam.service.iamv1.UserUpdateCompanyIdAndLabelResponse
	15, // 16: api.iam.service.iamservicev1.IAMUserService.DeleteUser:output_type -> api.iam.service.iamv1.UserDeleteUserResponse
	16, // 17: api.iam.service.iamservicev1.IAMUserService.CheckLastLogin:output_type -> api.iam.service.iamv1.UserCheckLastLoginResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_iam_service_v1_services_iam_user_service_v1_proto_init() }
func file_api_iam_service_v1_services_iam_user_service_v1_proto_init() {
	if File_api_iam_service_v1_services_iam_user_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_iam_service_v1_services_iam_user_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_iam_service_v1_services_iam_user_service_v1_proto_goTypes,
		DependencyIndexes: file_api_iam_service_v1_services_iam_user_service_v1_proto_depIdxs,
	}.Build()
	File_api_iam_service_v1_services_iam_user_service_v1_proto = out.File
	file_api_iam_service_v1_services_iam_user_service_v1_proto_rawDesc = nil
	file_api_iam_service_v1_services_iam_user_service_v1_proto_goTypes = nil
	file_api_iam_service_v1_services_iam_user_service_v1_proto_depIdxs = nil
}