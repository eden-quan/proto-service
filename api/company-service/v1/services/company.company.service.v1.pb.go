// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/company-service/v1/services/company.company.service.v1.proto

package companyservicev1

import (
	reflect "reflect"

	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/company-service/v1/resources"
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

var File_api_company_service_v1_services_company_company_service_v1_proto protoreflect.FileDescriptor

var file_api_company_service_v1_services_company_company_service_v1_proto_rawDesc = []byte{
	0x0a, 0x40, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x24, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x1a, 0x32, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x42, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x9e, 0x0c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x56, 0x31,
	0x12, 0x92, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x3c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x80, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79, 0x55, 0x69, 0x64, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79, 0x55, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79,
	0x55, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x95, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x12, 0x3d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x83, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x86, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7a, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x34, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x77, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x92, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x3c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x9e, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x4f, 0x52, 0x47, 0x46, 0x6f, 0x72, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x40, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x72, 0x4f, 0x52, 0x47, 0x46, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x4f, 0x52, 0x47, 0x46, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x92, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6e,
	0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x3c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x16, 0xfa, 0xe9, 0x30,
	0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61,
	0x69, 0x6e, 0x75, 0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e,
	0x2d, 0x71, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_company_service_v1_services_company_company_service_v1_proto_goTypes = []interface{}{
	(*resources.GetEmployeeInfoByLoginRequest)(nil),     // 0: api.company.service.companyv1.GetEmployeeInfoByLoginRequest
	(*resources.GetEmployeeByUidRequest)(nil),           // 1: api.company.service.companyv1.GetEmployeeByUidRequest
	(*resources.UpdateEmployeeRequest)(nil),             // 2: api.company.service.companyv1.UpdateEmployeeRequest
	(*resources.CheckAndUpdatePrivilegeRequest)(nil),    // 3: api.company.service.companyv1.CheckAndUpdatePrivilegeRequest
	(*resources.GetPrivilegeLimitRequest)(nil),          // 4: api.company.service.companyv1.GetPrivilegeLimitRequest
	(*resources.GetPartCompanyListRequest)(nil),         // 5: api.company.service.companyv1.GetPartCompanyListRequest
	(*resources.GetCompanyByIdRequest)(nil),             // 6: api.company.service.companyv1.GetCompanyByIdRequest
	(*resources.GetUsersByIdsRequest)(nil),              // 7: api.company.service.companyv1.GetUsersByIdsRequest
	(*resources.GetSupplierCompanyByIdRequest)(nil),     // 8: api.company.service.companyv1.GetSupplierCompanyByIdRequest
	(*resources.CreateDesignerORGForPersonRequest)(nil), // 9: api.company.service.companyv1.CreateDesignerORGForPersonRequest
	(*resources.GetAndCheckLastCompanyRequest)(nil),     // 10: api.company.service.companyv1.GetAndCheckLastCompanyRequest
	(*resources.GetEmployeeInfoByLoginReply)(nil),       // 11: api.company.service.companyv1.GetEmployeeInfoByLoginReply
	(*resources.GetEmployeeByUidReply)(nil),             // 12: api.company.service.companyv1.GetEmployeeByUidReply
	(*resources.UpdateEmployeeReply)(nil),               // 13: api.company.service.companyv1.UpdateEmployeeReply
	(*resources.CheckAndUpdatePrivilegeReply)(nil),      // 14: api.company.service.companyv1.CheckAndUpdatePrivilegeReply
	(*resources.GetPrivilegeLimitReply)(nil),            // 15: api.company.service.companyv1.GetPrivilegeLimitReply
	(*resources.GetPartCompanyListReply)(nil),           // 16: api.company.service.companyv1.GetPartCompanyListReply
	(*resources.GetCompanyByIdReply)(nil),               // 17: api.company.service.companyv1.GetCompanyByIdReply
	(*resources.GetUsersByIdsReply)(nil),                // 18: api.company.service.companyv1.GetUsersByIdsReply
	(*resources.GetSupplierCompanyByIdReply)(nil),       // 19: api.company.service.companyv1.GetSupplierCompanyByIdReply
	(*resources.CreateDesignerORGForPersonReply)(nil),   // 20: api.company.service.companyv1.CreateDesignerORGForPersonReply
	(*resources.GetAndCheckLastCompanyReply)(nil),       // 21: api.company.service.companyv1.GetAndCheckLastCompanyReply
}
var file_api_company_service_v1_services_company_company_service_v1_proto_depIdxs = []int32{
	0,  // 0: api.company.service.companyservicev1.CompanyV1.GetEmployeeInfoByLogin:input_type -> api.company.service.companyv1.GetEmployeeInfoByLoginRequest
	1,  // 1: api.company.service.companyservicev1.CompanyV1.GetEmployeeByUid:input_type -> api.company.service.companyv1.GetEmployeeByUidRequest
	2,  // 2: api.company.service.companyservicev1.CompanyV1.UpdateEmployee:input_type -> api.company.service.companyv1.UpdateEmployeeRequest
	3,  // 3: api.company.service.companyservicev1.CompanyV1.CheckAndUpdatePrivilege:input_type -> api.company.service.companyv1.CheckAndUpdatePrivilegeRequest
	4,  // 4: api.company.service.companyservicev1.CompanyV1.GetPrivilegeLimit:input_type -> api.company.service.companyv1.GetPrivilegeLimitRequest
	5,  // 5: api.company.service.companyservicev1.CompanyV1.GetPartCompanyList:input_type -> api.company.service.companyv1.GetPartCompanyListRequest
	6,  // 6: api.company.service.companyservicev1.CompanyV1.GetCompanyById:input_type -> api.company.service.companyv1.GetCompanyByIdRequest
	7,  // 7: api.company.service.companyservicev1.CompanyV1.GetUsersByIds:input_type -> api.company.service.companyv1.GetUsersByIdsRequest
	8,  // 8: api.company.service.companyservicev1.CompanyV1.GetSupplierCompanyById:input_type -> api.company.service.companyv1.GetSupplierCompanyByIdRequest
	9,  // 9: api.company.service.companyservicev1.CompanyV1.CreateDesignerORGForPerson:input_type -> api.company.service.companyv1.CreateDesignerORGForPersonRequest
	10, // 10: api.company.service.companyservicev1.CompanyV1.GetAndCheckLastCompany:input_type -> api.company.service.companyv1.GetAndCheckLastCompanyRequest
	11, // 11: api.company.service.companyservicev1.CompanyV1.GetEmployeeInfoByLogin:output_type -> api.company.service.companyv1.GetEmployeeInfoByLoginReply
	12, // 12: api.company.service.companyservicev1.CompanyV1.GetEmployeeByUid:output_type -> api.company.service.companyv1.GetEmployeeByUidReply
	13, // 13: api.company.service.companyservicev1.CompanyV1.UpdateEmployee:output_type -> api.company.service.companyv1.UpdateEmployeeReply
	14, // 14: api.company.service.companyservicev1.CompanyV1.CheckAndUpdatePrivilege:output_type -> api.company.service.companyv1.CheckAndUpdatePrivilegeReply
	15, // 15: api.company.service.companyservicev1.CompanyV1.GetPrivilegeLimit:output_type -> api.company.service.companyv1.GetPrivilegeLimitReply
	16, // 16: api.company.service.companyservicev1.CompanyV1.GetPartCompanyList:output_type -> api.company.service.companyv1.GetPartCompanyListReply
	17, // 17: api.company.service.companyservicev1.CompanyV1.GetCompanyById:output_type -> api.company.service.companyv1.GetCompanyByIdReply
	18, // 18: api.company.service.companyservicev1.CompanyV1.GetUsersByIds:output_type -> api.company.service.companyv1.GetUsersByIdsReply
	19, // 19: api.company.service.companyservicev1.CompanyV1.GetSupplierCompanyById:output_type -> api.company.service.companyv1.GetSupplierCompanyByIdReply
	20, // 20: api.company.service.companyservicev1.CompanyV1.CreateDesignerORGForPerson:output_type -> api.company.service.companyv1.CreateDesignerORGForPersonReply
	21, // 21: api.company.service.companyservicev1.CompanyV1.GetAndCheckLastCompany:output_type -> api.company.service.companyv1.GetAndCheckLastCompanyReply
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_company_service_v1_services_company_company_service_v1_proto_init() }
func file_api_company_service_v1_services_company_company_service_v1_proto_init() {
	if File_api_company_service_v1_services_company_company_service_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_company_service_v1_services_company_company_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_company_service_v1_services_company_company_service_v1_proto_goTypes,
		DependencyIndexes: file_api_company_service_v1_services_company_company_service_v1_proto_depIdxs,
	}.Build()
	File_api_company_service_v1_services_company_company_service_v1_proto = out.File
	file_api_company_service_v1_services_company_company_service_v1_proto_rawDesc = nil
	file_api_company_service_v1_services_company_company_service_v1_proto_goTypes = nil
	file_api_company_service_v1_services_company_company_service_v1_proto_depIdxs = nil
}
