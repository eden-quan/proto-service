// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.1
// - protoc             v5.26.1
// source: api/messagecenter-service/v1/services/messagecenter.messagecenter.service.v1.proto

package messagecenterservicev1

import (
	context "context"
	fmt "fmt"

	client "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/client"
	def "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/def"
	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/resources"
	fx "go.uber.org/fx"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7
const _ = fx.Version

var _ = new(fmt.Stringer)
var _ = new(def.Server)
var _ = new(client.RegisterGRPCClientFactoryType)

const (
	MessageCenterV1_SendUserMsg_FullMethodName                  = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/SendUserMsg"
	MessageCenterV1_GetUserMsgList_FullMethodName               = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/GetUserMsgList"
	MessageCenterV1_UpdateUserMsgStatus_FullMethodName          = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/UpdateUserMsgStatus"
	MessageCenterV1_ReadUserMsg_FullMethodName                  = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/ReadUserMsg"
	MessageCenterV1_ReadUserMsgByUserDataID_FullMethodName      = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/ReadUserMsgByUserDataID"
	MessageCenterV1_ExistUnreadUserMsg_FullMethodName           = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/ExistUnreadUserMsg"
	MessageCenterV1_UpdateMsgContent_FullMethodName             = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/UpdateMsgContent"
	MessageCenterV1_UpdateMsgContentByUserDataID_FullMethodName = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/UpdateMsgContentByUserDataID"
	MessageCenterV1_SetThirdPartyMsgID_FullMethodName           = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/SetThirdPartyMsgID"
	MessageCenterV1_ReadThirdPartyMsg_FullMethodName            = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/ReadThirdPartyMsg"
	MessageCenterV1_ServerStream_FullMethodName                 = "/api.messagecenter.service.messagecenterservicev1.MessageCenterV1/ServerStream"
)

// MessageCenterV1Client is the client API for MessageCenterV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageCenterV1Client interface {
	// 发送用户消息
	SendUserMsg(ctx context.Context, in *resources.SendUserMsgRequest, opts ...grpc.CallOption) (*resources.SendUserMsgResponse, error)
	// 获取用户消息
	GetUserMsgList(ctx context.Context, in *resources.GetUserMsgListRequest, opts ...grpc.CallOption) (*resources.GetUserMsgListResponse, error)
	// 更新用户消息状态
	UpdateUserMsgStatus(ctx context.Context, in *resources.UpdateUserMsgStatusRequest, opts ...grpc.CallOption) (*resources.UpdateUserMsgStatusResponse, error)
	// 已读用户消息
	ReadUserMsg(ctx context.Context, in *resources.ReadUserMsgRequest, opts ...grpc.CallOption) (*resources.ReadUserMsgResponse, error)
	ReadUserMsgByUserDataID(ctx context.Context, in *resources.ReadUserMsgByUserDataIDRequest, opts ...grpc.CallOption) (*resources.ReadUserMsgResponse, error)
	// 是否存在未读消息
	ExistUnreadUserMsg(ctx context.Context, in *resources.ExistUnreadUserMsgRequest, opts ...grpc.CallOption) (*resources.ExistUnreadUserMsgResponse, error)
	// 更新消息内容
	UpdateMsgContent(ctx context.Context, in *resources.UpdateMsgContentRequest, opts ...grpc.CallOption) (*resources.UpdateMsgContentResponse, error)
	// 更新消息内容，根据 userdataid
	UpdateMsgContentByUserDataID(ctx context.Context, in *resources.UpdateMsgContentByUserDataIDRequest, opts ...grpc.CallOption) (*resources.UpdateMsgContentResponse, error)
	// 设置第三方消息id
	SetThirdPartyMsgID(ctx context.Context, in *resources.SetThirdPartyMsgIDRequest, opts ...grpc.CallOption) (*resources.SetThirdPartyMsgIDResponse, error)
	// 读第三方消息
	ReadThirdPartyMsg(ctx context.Context, in *resources.ReadThirdPartyMsgRequest, opts ...grpc.CallOption) (*resources.ReadThirdPartyMsgResponse, error)
	// 前端的长连，用于实时推送消息给前端
	ServerStream(ctx context.Context, in *resources.ServerStreamRequest, opts ...grpc.CallOption) (MessageCenterV1_ServerStreamClient, error)
	RegisterNameForDiscover() string
}

type messageCenterV1Client struct {
	cc grpc.ClientConnInterface
}

func (c *messageCenterV1Client) RegisterNameForDiscover() string {
	return "/messagecenter-service/v1"
}

func newMessageCenterV1Client(cc grpc.ClientConnInterface) MessageCenterV1Client {
	return &messageCenterV1Client{cc}
}

func registerMessageCenterV1ClientGRPCNameProvider() []string {
	return []string{"/messagecenter-service/v1", "grpc"}
}

// RegisterMessageCenterV1ClientGRPCProvider is the provider for injection framework
// creator is the factory function which use to create the MessageCenterV1Client instance/implement
// the creator function receive dependency provided by fx to create ClientInterface,
// and returns the new dependency can use by others functions
func RegisterMessageCenterV1ClientGRPCProvider(creator interface{}) []interface{} {
	return []interface{}{
		fx.Annotate(
			newMessageCenterV1Client,
			fx.As(new(MessageCenterV1Client)),
			fx.ParamTags(`name:"/messagecenter-service/v1/grpc/messageCenterV1"`),
		),
		fx.Annotate(
			creator,
			fx.As(new(grpc.ClientConnInterface)),
			fx.ParamTags(`name:"/messagecenter-service/v1/grpc/name/messageCenterV1"`),
			fx.ResultTags(`name:"/messagecenter-service/v1/grpc/messageCenterV1"`),
		),
		fx.Annotate(
			registerMessageCenterV1ClientGRPCNameProvider,
			fx.ResultTags(`name:"/messagecenter-service/v1/grpc/name/messageCenterV1"`),
		),
	}
}

type MessageCenterV1ClientGRPCFactory interface {
	New(conf *def.Server) (MessageCenterV1Client, error)
}

type messageCenterV1ClientGRPCFactoryImpl struct {
	factory client.RegisterGRPCClientFactoryType
}

func (p *messageCenterV1ClientGRPCFactoryImpl) New(conf *def.Server) (MessageCenterV1Client, error) {
	cc, err := p.factory(conf)
	if err != nil {
		return nil, fmt.Errorf("create MessageCenterV1Client failed cause %s", err)
	}
	return &messageCenterV1Client{cc: cc}, nil
}

func RegisterMessageCenterV1ClientGRPCFactoryProvider(factory client.RegisterGRPCClientFactoryType) MessageCenterV1ClientGRPCFactory {
	return &messageCenterV1ClientGRPCFactoryImpl{factory: factory}
}

func (c *messageCenterV1Client) SendUserMsg(ctx context.Context, in *resources.SendUserMsgRequest, opts ...grpc.CallOption) (*resources.SendUserMsgResponse, error) {
	out := new(resources.SendUserMsgResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_SendUserMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) GetUserMsgList(ctx context.Context, in *resources.GetUserMsgListRequest, opts ...grpc.CallOption) (*resources.GetUserMsgListResponse, error) {
	out := new(resources.GetUserMsgListResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_GetUserMsgList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) UpdateUserMsgStatus(ctx context.Context, in *resources.UpdateUserMsgStatusRequest, opts ...grpc.CallOption) (*resources.UpdateUserMsgStatusResponse, error) {
	out := new(resources.UpdateUserMsgStatusResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_UpdateUserMsgStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) ReadUserMsg(ctx context.Context, in *resources.ReadUserMsgRequest, opts ...grpc.CallOption) (*resources.ReadUserMsgResponse, error) {
	out := new(resources.ReadUserMsgResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_ReadUserMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) ReadUserMsgByUserDataID(ctx context.Context, in *resources.ReadUserMsgByUserDataIDRequest, opts ...grpc.CallOption) (*resources.ReadUserMsgResponse, error) {
	out := new(resources.ReadUserMsgResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_ReadUserMsgByUserDataID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) ExistUnreadUserMsg(ctx context.Context, in *resources.ExistUnreadUserMsgRequest, opts ...grpc.CallOption) (*resources.ExistUnreadUserMsgResponse, error) {
	out := new(resources.ExistUnreadUserMsgResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_ExistUnreadUserMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) UpdateMsgContent(ctx context.Context, in *resources.UpdateMsgContentRequest, opts ...grpc.CallOption) (*resources.UpdateMsgContentResponse, error) {
	out := new(resources.UpdateMsgContentResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_UpdateMsgContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) UpdateMsgContentByUserDataID(ctx context.Context, in *resources.UpdateMsgContentByUserDataIDRequest, opts ...grpc.CallOption) (*resources.UpdateMsgContentResponse, error) {
	out := new(resources.UpdateMsgContentResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_UpdateMsgContentByUserDataID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) SetThirdPartyMsgID(ctx context.Context, in *resources.SetThirdPartyMsgIDRequest, opts ...grpc.CallOption) (*resources.SetThirdPartyMsgIDResponse, error) {
	out := new(resources.SetThirdPartyMsgIDResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_SetThirdPartyMsgID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) ReadThirdPartyMsg(ctx context.Context, in *resources.ReadThirdPartyMsgRequest, opts ...grpc.CallOption) (*resources.ReadThirdPartyMsgResponse, error) {
	out := new(resources.ReadThirdPartyMsgResponse)
	err := c.cc.Invoke(ctx, MessageCenterV1_ReadThirdPartyMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageCenterV1Client) ServerStream(ctx context.Context, in *resources.ServerStreamRequest, opts ...grpc.CallOption) (MessageCenterV1_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageCenterV1_ServiceDesc.Streams[0], MessageCenterV1_ServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messageCenterV1ServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageCenterV1_ServerStreamClient interface {
	Recv() (*resources.ServerStreamMsg, error)
	grpc.ClientStream
}

type messageCenterV1ServerStreamClient struct {
	grpc.ClientStream
}

func (x *messageCenterV1ServerStreamClient) Recv() (*resources.ServerStreamMsg, error) {
	m := new(resources.ServerStreamMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageCenterV1Server is the server API for MessageCenterV1 service.
// All implementations must embed UnimplementedMessageCenterV1Server
// for forward compatibility
type MessageCenterV1Server interface {
	// 发送用户消息
	SendUserMsg(context.Context, *resources.SendUserMsgRequest) (*resources.SendUserMsgResponse, error)
	// 获取用户消息
	GetUserMsgList(context.Context, *resources.GetUserMsgListRequest) (*resources.GetUserMsgListResponse, error)
	// 更新用户消息状态
	UpdateUserMsgStatus(context.Context, *resources.UpdateUserMsgStatusRequest) (*resources.UpdateUserMsgStatusResponse, error)
	// 已读用户消息
	ReadUserMsg(context.Context, *resources.ReadUserMsgRequest) (*resources.ReadUserMsgResponse, error)
	ReadUserMsgByUserDataID(context.Context, *resources.ReadUserMsgByUserDataIDRequest) (*resources.ReadUserMsgResponse, error)
	// 是否存在未读消息
	ExistUnreadUserMsg(context.Context, *resources.ExistUnreadUserMsgRequest) (*resources.ExistUnreadUserMsgResponse, error)
	// 更新消息内容
	UpdateMsgContent(context.Context, *resources.UpdateMsgContentRequest) (*resources.UpdateMsgContentResponse, error)
	// 更新消息内容，根据 userdataid
	UpdateMsgContentByUserDataID(context.Context, *resources.UpdateMsgContentByUserDataIDRequest) (*resources.UpdateMsgContentResponse, error)
	// 设置第三方消息id
	SetThirdPartyMsgID(context.Context, *resources.SetThirdPartyMsgIDRequest) (*resources.SetThirdPartyMsgIDResponse, error)
	// 读第三方消息
	ReadThirdPartyMsg(context.Context, *resources.ReadThirdPartyMsgRequest) (*resources.ReadThirdPartyMsgResponse, error)
	// 前端的长连，用于实时推送消息给前端
	ServerStream(*resources.ServerStreamRequest, MessageCenterV1_ServerStreamServer) error
	mustEmbedUnimplementedMessageCenterV1Server()
}

// Generate Injection
type registerMessageCenterV1ServerGRPCResult struct{}

func (*registerMessageCenterV1ServerGRPCResult) String() string {
	return "MessageCenterV1ServerGRPCServer"
}

func RegisterMessageCenterV1ServerGRPCProvider(newer interface{}) []interface{} {
	return []interface{}{
		// For provide dependency
		fx.Annotate(
			newer,
			fx.As(new(MessageCenterV1Server)),
		),
		// For create instance
		fx.Annotate(
			registerMessageCenterV1ServerProviderImpl,
			fx.As(new(fmt.Stringer)),
			fx.ResultTags(`group:"grpc_register"`),
		),
	}
}

// registerMessageCenterV1ServerProviderImpl use to trigger register
func registerMessageCenterV1ServerProviderImpl(s grpc.ServiceRegistrar, srv MessageCenterV1Server) *registerMessageCenterV1ServerGRPCResult {
	registerMessageCenterV1Server(s, srv)
	return &registerMessageCenterV1ServerGRPCResult{}
}

// UnimplementedMessageCenterV1Server must be embedded to have forward compatible implementations.
type UnimplementedMessageCenterV1Server struct {
}

func (UnimplementedMessageCenterV1Server) SendUserMsg(context.Context, *resources.SendUserMsgRequest) (*resources.SendUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUserMsg not implemented")
}
func (UnimplementedMessageCenterV1Server) GetUserMsgList(context.Context, *resources.GetUserMsgListRequest) (*resources.GetUserMsgListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMsgList not implemented")
}
func (UnimplementedMessageCenterV1Server) UpdateUserMsgStatus(context.Context, *resources.UpdateUserMsgStatusRequest) (*resources.UpdateUserMsgStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMsgStatus not implemented")
}
func (UnimplementedMessageCenterV1Server) ReadUserMsg(context.Context, *resources.ReadUserMsgRequest) (*resources.ReadUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUserMsg not implemented")
}
func (UnimplementedMessageCenterV1Server) ReadUserMsgByUserDataID(context.Context, *resources.ReadUserMsgByUserDataIDRequest) (*resources.ReadUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadUserMsgByUserDataID not implemented")
}
func (UnimplementedMessageCenterV1Server) ExistUnreadUserMsg(context.Context, *resources.ExistUnreadUserMsgRequest) (*resources.ExistUnreadUserMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUnreadUserMsg not implemented")
}
func (UnimplementedMessageCenterV1Server) UpdateMsgContent(context.Context, *resources.UpdateMsgContentRequest) (*resources.UpdateMsgContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMsgContent not implemented")
}
func (UnimplementedMessageCenterV1Server) UpdateMsgContentByUserDataID(context.Context, *resources.UpdateMsgContentByUserDataIDRequest) (*resources.UpdateMsgContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMsgContentByUserDataID not implemented")
}
func (UnimplementedMessageCenterV1Server) SetThirdPartyMsgID(context.Context, *resources.SetThirdPartyMsgIDRequest) (*resources.SetThirdPartyMsgIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetThirdPartyMsgID not implemented")
}
func (UnimplementedMessageCenterV1Server) ReadThirdPartyMsg(context.Context, *resources.ReadThirdPartyMsgRequest) (*resources.ReadThirdPartyMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadThirdPartyMsg not implemented")
}
func (UnimplementedMessageCenterV1Server) ServerStream(*resources.ServerStreamRequest, MessageCenterV1_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedMessageCenterV1Server) mustEmbedUnimplementedMessageCenterV1Server() {}

// UnsafeMessageCenterV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageCenterV1Server will
// result in compilation errors.
type UnsafeMessageCenterV1Server interface {
	mustEmbedUnimplementedMessageCenterV1Server()
}

func registerMessageCenterV1Server(s grpc.ServiceRegistrar, srv MessageCenterV1Server) {
	s.RegisterService(&MessageCenterV1_ServiceDesc, srv)
}

func _MessageCenterV1_SendUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SendUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).SendUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_SendUserMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).SendUserMsg(ctx, req.(*resources.SendUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_GetUserMsgList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.GetUserMsgListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).GetUserMsgList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_GetUserMsgList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).GetUserMsgList(ctx, req.(*resources.GetUserMsgListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_UpdateUserMsgStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateUserMsgStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).UpdateUserMsgStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_UpdateUserMsgStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).UpdateUserMsgStatus(ctx, req.(*resources.UpdateUserMsgStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_ReadUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).ReadUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_ReadUserMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).ReadUserMsg(ctx, req.(*resources.ReadUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_ReadUserMsgByUserDataID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadUserMsgByUserDataIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).ReadUserMsgByUserDataID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_ReadUserMsgByUserDataID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).ReadUserMsgByUserDataID(ctx, req.(*resources.ReadUserMsgByUserDataIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_ExistUnreadUserMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ExistUnreadUserMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).ExistUnreadUserMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_ExistUnreadUserMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).ExistUnreadUserMsg(ctx, req.(*resources.ExistUnreadUserMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_UpdateMsgContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateMsgContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).UpdateMsgContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_UpdateMsgContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).UpdateMsgContent(ctx, req.(*resources.UpdateMsgContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_UpdateMsgContentByUserDataID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.UpdateMsgContentByUserDataIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).UpdateMsgContentByUserDataID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_UpdateMsgContentByUserDataID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).UpdateMsgContentByUserDataID(ctx, req.(*resources.UpdateMsgContentByUserDataIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_SetThirdPartyMsgID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SetThirdPartyMsgIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).SetThirdPartyMsgID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_SetThirdPartyMsgID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).SetThirdPartyMsgID(ctx, req.(*resources.SetThirdPartyMsgIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_ReadThirdPartyMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.ReadThirdPartyMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageCenterV1Server).ReadThirdPartyMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MessageCenterV1_ReadThirdPartyMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageCenterV1Server).ReadThirdPartyMsg(ctx, req.(*resources.ReadThirdPartyMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageCenterV1_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(resources.ServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageCenterV1Server).ServerStream(m, &messageCenterV1ServerStreamServer{stream})
}

type MessageCenterV1_ServerStreamServer interface {
	Send(*resources.ServerStreamMsg) error
	grpc.ServerStream
}

type messageCenterV1ServerStreamServer struct {
	grpc.ServerStream
}

func (x *messageCenterV1ServerStreamServer) Send(m *resources.ServerStreamMsg) error {
	return x.ServerStream.SendMsg(m)
}

// MessageCenterV1_ServiceDesc is the grpc.ServiceDesc for MessageCenterV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageCenterV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.messagecenter.service.messagecenterservicev1.MessageCenterV1",
	HandlerType: (*MessageCenterV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUserMsg",
			Handler:    _MessageCenterV1_SendUserMsg_Handler,
		},
		{
			MethodName: "GetUserMsgList",
			Handler:    _MessageCenterV1_GetUserMsgList_Handler,
		},
		{
			MethodName: "UpdateUserMsgStatus",
			Handler:    _MessageCenterV1_UpdateUserMsgStatus_Handler,
		},
		{
			MethodName: "ReadUserMsg",
			Handler:    _MessageCenterV1_ReadUserMsg_Handler,
		},
		{
			MethodName: "ReadUserMsgByUserDataID",
			Handler:    _MessageCenterV1_ReadUserMsgByUserDataID_Handler,
		},
		{
			MethodName: "ExistUnreadUserMsg",
			Handler:    _MessageCenterV1_ExistUnreadUserMsg_Handler,
		},
		{
			MethodName: "UpdateMsgContent",
			Handler:    _MessageCenterV1_UpdateMsgContent_Handler,
		},
		{
			MethodName: "UpdateMsgContentByUserDataID",
			Handler:    _MessageCenterV1_UpdateMsgContentByUserDataID_Handler,
		},
		{
			MethodName: "SetThirdPartyMsgID",
			Handler:    _MessageCenterV1_SetThirdPartyMsgID_Handler,
		},
		{
			MethodName: "ReadThirdPartyMsg",
			Handler:    _MessageCenterV1_ReadThirdPartyMsg_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _MessageCenterV1_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/messagecenter-service/v1/services/messagecenter.messagecenter.service.v1.proto",
}
