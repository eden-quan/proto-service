// Code generated by protoc-gen-go-fx. DO NOT EDIT.
// versions:
// 	protoc-gen-go-fx v1.32.0-devel
// 	protoc        v5.26.1
// source: api/messagecenter-service/v1/enums/messagecenter.messagecenter.enum.v1.proto

package messagecenterenumv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 消息状态
type MsgStatusEnum_MsgStatus int32

const (
	// 未知
	MsgStatusEnum_MSG_STATUS_UNKNOWN MsgStatusEnum_MsgStatus = 0
	// 未读
	MsgStatusEnum_MSG_STATUS_UNREAD MsgStatusEnum_MsgStatus = 1
	// 已读
	MsgStatusEnum_MSG_STATUS_READ MsgStatusEnum_MsgStatus = 2
	// 已撤销
	MsgStatusEnum_MSG_STATUS_REVOKE MsgStatusEnum_MsgStatus = 3
)

// Enum value maps for MsgStatusEnum_MsgStatus.
var (
	MsgStatusEnum_MsgStatus_name = map[int32]string{
		0: "MSG_STATUS_UNKNOWN",
		1: "MSG_STATUS_UNREAD",
		2: "MSG_STATUS_READ",
		3: "MSG_STATUS_REVOKE",
	}
	MsgStatusEnum_MsgStatus_value = map[string]int32{
		"MSG_STATUS_UNKNOWN": 0,
		"MSG_STATUS_UNREAD":  1,
		"MSG_STATUS_READ":    2,
		"MSG_STATUS_REVOKE":  3,
	}
)

func (x MsgStatusEnum_MsgStatus) Enum() *MsgStatusEnum_MsgStatus {
	p := new(MsgStatusEnum_MsgStatus)
	*p = x
	return p
}

func (x MsgStatusEnum_MsgStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgStatusEnum_MsgStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[0].Descriptor()
}

func (MsgStatusEnum_MsgStatus) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[0]
}

func (x MsgStatusEnum_MsgStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgStatusEnum_MsgStatus.Descriptor instead.
func (MsgStatusEnum_MsgStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{0, 0}
}

// 消息展示类型
type MsgDisplayTypeEnum_MsgDisplayType int32

const (
	// 未知
	MsgDisplayTypeEnum_MSG_DISPLAY_UNKNOWN MsgDisplayTypeEnum_MsgDisplayType = 0
	// 消息卡片
	MsgDisplayTypeEnum_MSG_DISPLAY_CARD MsgDisplayTypeEnum_MsgDisplayType = 1
	// 文本
	MsgDisplayTypeEnum_MSG_DISPLAY_TEXT MsgDisplayTypeEnum_MsgDisplayType = 2
	// markdown 文本
	MsgDisplayTypeEnum_MSG_DISPLAY_MARKDOWN MsgDisplayTypeEnum_MsgDisplayType = 3
)

// Enum value maps for MsgDisplayTypeEnum_MsgDisplayType.
var (
	MsgDisplayTypeEnum_MsgDisplayType_name = map[int32]string{
		0: "MSG_DISPLAY_UNKNOWN",
		1: "MSG_DISPLAY_CARD",
		2: "MSG_DISPLAY_TEXT",
		3: "MSG_DISPLAY_MARKDOWN",
	}
	MsgDisplayTypeEnum_MsgDisplayType_value = map[string]int32{
		"MSG_DISPLAY_UNKNOWN":  0,
		"MSG_DISPLAY_CARD":     1,
		"MSG_DISPLAY_TEXT":     2,
		"MSG_DISPLAY_MARKDOWN": 3,
	}
)

func (x MsgDisplayTypeEnum_MsgDisplayType) Enum() *MsgDisplayTypeEnum_MsgDisplayType {
	p := new(MsgDisplayTypeEnum_MsgDisplayType)
	*p = x
	return p
}

func (x MsgDisplayTypeEnum_MsgDisplayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgDisplayTypeEnum_MsgDisplayType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[1].Descriptor()
}

func (MsgDisplayTypeEnum_MsgDisplayType) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[1]
}

func (x MsgDisplayTypeEnum_MsgDisplayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgDisplayTypeEnum_MsgDisplayType.Descriptor instead.
func (MsgDisplayTypeEnum_MsgDisplayType) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{1, 0}
}

// 消息类型
type MsgTypeEnum_MsgType int32

const (
	// 未知
	MsgTypeEnum_MSG_TYPE_UNKNOWN MsgTypeEnum_MsgType = 0
	// 测试
	MsgTypeEnum_MSG_TYPE_TEST1 MsgTypeEnum_MsgType = 1
	MsgTypeEnum_MSG_TYPE_TEST2 MsgTypeEnum_MsgType = 2
	// 创建企业
	MsgTypeEnum_MSG_TYPE_CREATE_COMPANY MsgTypeEnum_MsgType = 100
	// 邀请进企业消息
	MsgTypeEnum_MSG_TYPE_INVITE_JOIN_COMPANY MsgTypeEnum_MsgType = 101
	// 从企业移除成员消息
	MsgTypeEnum_MSG_TYPE_REMOVE_FROM_COMPANY MsgTypeEnum_MsgType = 103
	// 新成员加入通知
	MsgTypeEnum_MSG_TYPE_JOIN_COMPANY MsgTypeEnum_MsgType = 105
	// 事件消息-切换了企业
	MsgTypeEnum_MSG_TYPE_SWITCH_COMPANY MsgTypeEnum_MsgType = 108
	// 事件消息-从企业移除成员
	MsgTypeEnum_MSG_TYPE_EVENT_REMOVE_FROM_COMPANY MsgTypeEnum_MsgType = 109
	// 单点登录，被踢下线事件消息
	MsgTypeEnum_MSG_TYPE_KICK_OFFLINE MsgTypeEnum_MsgType = 750
)

// Enum value maps for MsgTypeEnum_MsgType.
var (
	MsgTypeEnum_MsgType_name = map[int32]string{
		0:   "MSG_TYPE_UNKNOWN",
		1:   "MSG_TYPE_TEST1",
		2:   "MSG_TYPE_TEST2",
		100: "MSG_TYPE_CREATE_COMPANY",
		101: "MSG_TYPE_INVITE_JOIN_COMPANY",
		103: "MSG_TYPE_REMOVE_FROM_COMPANY",
		105: "MSG_TYPE_JOIN_COMPANY",
		108: "MSG_TYPE_SWITCH_COMPANY",
		109: "MSG_TYPE_EVENT_REMOVE_FROM_COMPANY",
		750: "MSG_TYPE_KICK_OFFLINE",
	}
	MsgTypeEnum_MsgType_value = map[string]int32{
		"MSG_TYPE_UNKNOWN":                   0,
		"MSG_TYPE_TEST1":                     1,
		"MSG_TYPE_TEST2":                     2,
		"MSG_TYPE_CREATE_COMPANY":            100,
		"MSG_TYPE_INVITE_JOIN_COMPANY":       101,
		"MSG_TYPE_REMOVE_FROM_COMPANY":       103,
		"MSG_TYPE_JOIN_COMPANY":              105,
		"MSG_TYPE_SWITCH_COMPANY":            108,
		"MSG_TYPE_EVENT_REMOVE_FROM_COMPANY": 109,
		"MSG_TYPE_KICK_OFFLINE":              750,
	}
)

func (x MsgTypeEnum_MsgType) Enum() *MsgTypeEnum_MsgType {
	p := new(MsgTypeEnum_MsgType)
	*p = x
	return p
}

func (x MsgTypeEnum_MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgTypeEnum_MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[2].Descriptor()
}

func (MsgTypeEnum_MsgType) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[2]
}

func (x MsgTypeEnum_MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgTypeEnum_MsgType.Descriptor instead.
func (MsgTypeEnum_MsgType) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{2, 0}
}

// 按钮颜色
type MsgButtonColorEnum_MsgButtonColor int32

const (
	MsgButtonColorEnum_BUTTON_COLOR_UNKNOWN MsgButtonColorEnum_MsgButtonColor = 0
	MsgButtonColorEnum_BUTTON_COLOR_DEFAULT MsgButtonColorEnum_MsgButtonColor = 1
	MsgButtonColorEnum_BUTTON_COLOR_PRIMARY MsgButtonColorEnum_MsgButtonColor = 2
	MsgButtonColorEnum_BUTTON_COLOR_DANGER  MsgButtonColorEnum_MsgButtonColor = 3
)

// Enum value maps for MsgButtonColorEnum_MsgButtonColor.
var (
	MsgButtonColorEnum_MsgButtonColor_name = map[int32]string{
		0: "BUTTON_COLOR_UNKNOWN",
		1: "BUTTON_COLOR_DEFAULT",
		2: "BUTTON_COLOR_PRIMARY",
		3: "BUTTON_COLOR_DANGER",
	}
	MsgButtonColorEnum_MsgButtonColor_value = map[string]int32{
		"BUTTON_COLOR_UNKNOWN": 0,
		"BUTTON_COLOR_DEFAULT": 1,
		"BUTTON_COLOR_PRIMARY": 2,
		"BUTTON_COLOR_DANGER":  3,
	}
)

func (x MsgButtonColorEnum_MsgButtonColor) Enum() *MsgButtonColorEnum_MsgButtonColor {
	p := new(MsgButtonColorEnum_MsgButtonColor)
	*p = x
	return p
}

func (x MsgButtonColorEnum_MsgButtonColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgButtonColorEnum_MsgButtonColor) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[3].Descriptor()
}

func (MsgButtonColorEnum_MsgButtonColor) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[3]
}

func (x MsgButtonColorEnum_MsgButtonColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgButtonColorEnum_MsgButtonColor.Descriptor instead.
func (MsgButtonColorEnum_MsgButtonColor) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{3, 0}
}

// 按钮动作
type MsgButtonActionTypeEnum_MsgButtonActionType int32

const (
	MsgButtonActionTypeEnum_BUTTON_ACTION_UNKNOWN MsgButtonActionTypeEnum_MsgButtonActionType = 0
	// 新标签打开
	MsgButtonActionTypeEnum_BUTTON_ACTION_OPEN_BLANK MsgButtonActionTypeEnum_MsgButtonActionType = 1
	// 当前页面打开
	MsgButtonActionTypeEnum_BUTTON_ACTION_OPEN_CURRENT MsgButtonActionTypeEnum_MsgButtonActionType = 2
	// 直接请求
	MsgButtonActionTypeEnum_BUTTON_ACTION_OPEN_REQUEST MsgButtonActionTypeEnum_MsgButtonActionType = 3
)

// Enum value maps for MsgButtonActionTypeEnum_MsgButtonActionType.
var (
	MsgButtonActionTypeEnum_MsgButtonActionType_name = map[int32]string{
		0: "BUTTON_ACTION_UNKNOWN",
		1: "BUTTON_ACTION_OPEN_BLANK",
		2: "BUTTON_ACTION_OPEN_CURRENT",
		3: "BUTTON_ACTION_OPEN_REQUEST",
	}
	MsgButtonActionTypeEnum_MsgButtonActionType_value = map[string]int32{
		"BUTTON_ACTION_UNKNOWN":      0,
		"BUTTON_ACTION_OPEN_BLANK":   1,
		"BUTTON_ACTION_OPEN_CURRENT": 2,
		"BUTTON_ACTION_OPEN_REQUEST": 3,
	}
)

func (x MsgButtonActionTypeEnum_MsgButtonActionType) Enum() *MsgButtonActionTypeEnum_MsgButtonActionType {
	p := new(MsgButtonActionTypeEnum_MsgButtonActionType)
	*p = x
	return p
}

func (x MsgButtonActionTypeEnum_MsgButtonActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgButtonActionTypeEnum_MsgButtonActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[4].Descriptor()
}

func (MsgButtonActionTypeEnum_MsgButtonActionType) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[4]
}

func (x MsgButtonActionTypeEnum_MsgButtonActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgButtonActionTypeEnum_MsgButtonActionType.Descriptor instead.
func (MsgButtonActionTypeEnum_MsgButtonActionType) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{4, 0}
}

// 按钮样式
type MsgButtonStyleEnum_MsgButtonStyle int32

const (
	// 正常样式
	MsgButtonStyleEnum_STYLE_NORMAL MsgButtonStyleEnum_MsgButtonStyle = 0
	// 隐藏
	MsgButtonStyleEnum_STYLE_HIDDEN MsgButtonStyleEnum_MsgButtonStyle = 1
	// 禁止
	MsgButtonStyleEnum_STYLE_DISABLE MsgButtonStyleEnum_MsgButtonStyle = 2
)

// Enum value maps for MsgButtonStyleEnum_MsgButtonStyle.
var (
	MsgButtonStyleEnum_MsgButtonStyle_name = map[int32]string{
		0: "STYLE_NORMAL",
		1: "STYLE_HIDDEN",
		2: "STYLE_DISABLE",
	}
	MsgButtonStyleEnum_MsgButtonStyle_value = map[string]int32{
		"STYLE_NORMAL":  0,
		"STYLE_HIDDEN":  1,
		"STYLE_DISABLE": 2,
	}
)

func (x MsgButtonStyleEnum_MsgButtonStyle) Enum() *MsgButtonStyleEnum_MsgButtonStyle {
	p := new(MsgButtonStyleEnum_MsgButtonStyle)
	*p = x
	return p
}

func (x MsgButtonStyleEnum_MsgButtonStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgButtonStyleEnum_MsgButtonStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[5].Descriptor()
}

func (MsgButtonStyleEnum_MsgButtonStyle) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[5]
}

func (x MsgButtonStyleEnum_MsgButtonStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgButtonStyleEnum_MsgButtonStyle.Descriptor instead.
func (MsgButtonStyleEnum_MsgButtonStyle) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{5, 0}
}

// 平台
type PlatformEnum_Platform int32

const (
	PlatformEnum_PLATFORM_UNKNOWN PlatformEnum_Platform = 0
	PlatformEnum_RHINOBIRD        PlatformEnum_Platform = 1
)

// Enum value maps for PlatformEnum_Platform.
var (
	PlatformEnum_Platform_name = map[int32]string{
		0: "PLATFORM_UNKNOWN",
		1: "RHINOBIRD",
	}
	PlatformEnum_Platform_value = map[string]int32{
		"PLATFORM_UNKNOWN": 0,
		"RHINOBIRD":        1,
	}
)

func (x PlatformEnum_Platform) Enum() *PlatformEnum_Platform {
	p := new(PlatformEnum_Platform)
	*p = x
	return p
}

func (x PlatformEnum_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformEnum_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[6].Descriptor()
}

func (PlatformEnum_Platform) Type() protoreflect.EnumType {
	return &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes[6]
}

func (x PlatformEnum_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformEnum_Platform.Descriptor instead.
func (PlatformEnum_Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{6, 0}
}

type MsgStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgStatusEnum) Reset() {
	*x = MsgStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgStatusEnum) ProtoMessage() {}

func (x *MsgStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgStatusEnum.ProtoReflect.Descriptor instead.
func (*MsgStatusEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{0}
}

type MsgDisplayTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDisplayTypeEnum) Reset() {
	*x = MsgDisplayTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDisplayTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDisplayTypeEnum) ProtoMessage() {}

func (x *MsgDisplayTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDisplayTypeEnum.ProtoReflect.Descriptor instead.
func (*MsgDisplayTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{1}
}

type MsgTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgTypeEnum) Reset() {
	*x = MsgTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgTypeEnum) ProtoMessage() {}

func (x *MsgTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgTypeEnum.ProtoReflect.Descriptor instead.
func (*MsgTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{2}
}

type MsgButtonColorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgButtonColorEnum) Reset() {
	*x = MsgButtonColorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgButtonColorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgButtonColorEnum) ProtoMessage() {}

func (x *MsgButtonColorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgButtonColorEnum.ProtoReflect.Descriptor instead.
func (*MsgButtonColorEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{3}
}

type MsgButtonActionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgButtonActionTypeEnum) Reset() {
	*x = MsgButtonActionTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgButtonActionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgButtonActionTypeEnum) ProtoMessage() {}

func (x *MsgButtonActionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgButtonActionTypeEnum.ProtoReflect.Descriptor instead.
func (*MsgButtonActionTypeEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{4}
}

type MsgButtonStyleEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgButtonStyleEnum) Reset() {
	*x = MsgButtonStyleEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgButtonStyleEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgButtonStyleEnum) ProtoMessage() {}

func (x *MsgButtonStyleEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgButtonStyleEnum.ProtoReflect.Descriptor instead.
func (*MsgButtonStyleEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{5}
}

type PlatformEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlatformEnum) Reset() {
	*x = PlatformEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformEnum) ProtoMessage() {}

func (x *PlatformEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformEnum.ProtoReflect.Descriptor instead.
func (*PlatformEnum) Descriptor() ([]byte, []int) {
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP(), []int{6}
}

var File_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto protoreflect.FileDescriptor

var file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x22, 0x77, 0x0a,
	0x0d, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x66,
	0x0a, 0x09, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x4d,
	0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x53,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x4d, 0x53, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45,
	0x56, 0x4f, 0x4b, 0x45, 0x10, 0x03, 0x22, 0x85, 0x01, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x6f, 0x0a,
	0x0e, 0x4d, 0x73, 0x67, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x17, 0x0a, 0x13, 0x4d, 0x53, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f,
	0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x4d, 0x53, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x54, 0x45,
	0x58, 0x54, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x53, 0x47, 0x5f, 0x44, 0x49, 0x53, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x22, 0xb4,
	0x02, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xa4,
	0x02, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x53,
	0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x53,
	0x54, 0x31, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x45, 0x53, 0x54, 0x32, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x53, 0x47, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50,
	0x41, 0x4e, 0x59, 0x10, 0x64, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x53, 0x47, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10, 0x67, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x53, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41,
	0x4e, 0x59, 0x10, 0x69, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10,
	0x6c, 0x12, 0x26, 0x0a, 0x22, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e, 0x59, 0x10, 0x6d, 0x12, 0x1a, 0x0a, 0x15, 0x4d, 0x53, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0xee, 0x05, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x77, 0x0a, 0x0e,
	0x4d, 0x73, 0x67, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x14, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x55, 0x54, 0x54,
	0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54,
	0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c,
	0x4f, 0x52, 0x5f, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13,
	0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4c, 0x4f, 0x52, 0x5f, 0x44, 0x41, 0x4e,
	0x47, 0x45, 0x52, 0x10, 0x03, 0x22, 0xaa, 0x01, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x74,
	0x74, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x8e, 0x01, 0x0a, 0x13, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x55, 0x54,
	0x54, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x42, 0x4c, 0x41, 0x4e, 0x4b,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e, 0x54,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x55, 0x54, 0x54, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x03, 0x22, 0x5d, 0x0a, 0x12, 0x4d, 0x73, 0x67, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x47, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54,
	0x59, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x02, 0x22, 0x3f, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x2f, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a,
	0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x48, 0x49, 0x4e, 0x4f, 0x42, 0x49, 0x52, 0x44,
	0x10, 0x01, 0x42, 0x65, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6c, 0x61, 0x69,
	0x6e, 0x75, 0x6f, 0x6e, 0x69, 0x61, 0x6f, 0x2e, 0x63, 0x6e, 0x2f, 0x65, 0x64, 0x65, 0x6e, 0x2d,
	0x71, 0x75, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescOnce sync.Once
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescData = file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDesc
)

func file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescGZIP() []byte {
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescOnce.Do(func() {
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescData)
	})
	return file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDescData
}

var file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_goTypes = []interface{}{
	(MsgStatusEnum_MsgStatus)(0),                     // 0: api.messagecenter.service.messagecenterenumv1.MsgStatusEnum.MsgStatus
	(MsgDisplayTypeEnum_MsgDisplayType)(0),           // 1: api.messagecenter.service.messagecenterenumv1.MsgDisplayTypeEnum.MsgDisplayType
	(MsgTypeEnum_MsgType)(0),                         // 2: api.messagecenter.service.messagecenterenumv1.MsgTypeEnum.MsgType
	(MsgButtonColorEnum_MsgButtonColor)(0),           // 3: api.messagecenter.service.messagecenterenumv1.MsgButtonColorEnum.MsgButtonColor
	(MsgButtonActionTypeEnum_MsgButtonActionType)(0), // 4: api.messagecenter.service.messagecenterenumv1.MsgButtonActionTypeEnum.MsgButtonActionType
	(MsgButtonStyleEnum_MsgButtonStyle)(0),           // 5: api.messagecenter.service.messagecenterenumv1.MsgButtonStyleEnum.MsgButtonStyle
	(PlatformEnum_Platform)(0),                       // 6: api.messagecenter.service.messagecenterenumv1.PlatformEnum.Platform
	(*MsgStatusEnum)(nil),                            // 7: api.messagecenter.service.messagecenterenumv1.MsgStatusEnum
	(*MsgDisplayTypeEnum)(nil),                       // 8: api.messagecenter.service.messagecenterenumv1.MsgDisplayTypeEnum
	(*MsgTypeEnum)(nil),                              // 9: api.messagecenter.service.messagecenterenumv1.MsgTypeEnum
	(*MsgButtonColorEnum)(nil),                       // 10: api.messagecenter.service.messagecenterenumv1.MsgButtonColorEnum
	(*MsgButtonActionTypeEnum)(nil),                  // 11: api.messagecenter.service.messagecenterenumv1.MsgButtonActionTypeEnum
	(*MsgButtonStyleEnum)(nil),                       // 12: api.messagecenter.service.messagecenterenumv1.MsgButtonStyleEnum
	(*PlatformEnum)(nil),                             // 13: api.messagecenter.service.messagecenterenumv1.PlatformEnum
}
var file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_init() }
func file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_init() {
	if File_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgStatusEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDisplayTypeEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgTypeEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgButtonColorEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgButtonActionTypeEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgButtonStyleEnum); i {
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
		file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformEnum); i {
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
			RawDescriptor: file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_goTypes,
		DependencyIndexes: file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_depIdxs,
		EnumInfos:         file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_enumTypes,
		MessageInfos:      file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_msgTypes,
	}.Build()
	File_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto = out.File
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_rawDesc = nil
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_goTypes = nil
	file_api_messagecenter_service_v1_enums_messagecenter_messagecenter_enum_v1_proto_depIdxs = nil
}
