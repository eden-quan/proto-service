// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.1
// - protoc             v5.26.1
// source: api/email-service/v1/services/email.email.service.v1.proto

package emailservicev1

import (
	context "context"
	fmt "fmt"

	client "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/client"
	def "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/def"
	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/email-service/v1/resources"
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
	EmailServiceV1_SendEmail_FullMethodName = "/api.email.service.emailservicev1.EmailServiceV1/SendEmail"
)

// EmailServiceV1Client is the client API for EmailServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceV1Client interface {
	SendEmail(ctx context.Context, in *resources.SendEmailRequest, opts ...grpc.CallOption) (*resources.SendEmailReply, error)
	RegisterNameForDiscover() string
}

type emailServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func (c *emailServiceV1Client) RegisterNameForDiscover() string {
	return "/email-service/v1"
}

func newEmailServiceV1Client(cc grpc.ClientConnInterface) EmailServiceV1Client {
	return &emailServiceV1Client{cc}
}

func registerEmailServiceV1ClientGRPCNameProvider() []string {
	return []string{"/email-service/v1", "grpc"}
}

// RegisterEmailServiceV1ClientGRPCProvider is the provider for injection framework
// creator is the factory function which use to create the EmailServiceV1Client instance/implement
// the creator function receive dependency provided by fx to create ClientInterface,
// and returns the new dependency can use by others functions
func RegisterEmailServiceV1ClientGRPCProvider(creator interface{}) []interface{} {
	return []interface{}{
		fx.Annotate(
			newEmailServiceV1Client,
			fx.As(new(EmailServiceV1Client)),
			fx.ParamTags(`name:"/email-service/v1/grpc/emailServiceV1"`),
		),
		fx.Annotate(
			creator,
			fx.As(new(grpc.ClientConnInterface)),
			fx.ParamTags(`name:"/email-service/v1/grpc/name/emailServiceV1"`),
			fx.ResultTags(`name:"/email-service/v1/grpc/emailServiceV1"`),
		),
		fx.Annotate(
			registerEmailServiceV1ClientGRPCNameProvider,
			fx.ResultTags(`name:"/email-service/v1/grpc/name/emailServiceV1"`),
		),
	}
}

type EmailServiceV1ClientGRPCFactory interface {
	New(conf *def.Server) (EmailServiceV1Client, error)
}

type emailServiceV1ClientGRPCFactoryImpl struct {
	factory client.RegisterGRPCClientFactoryType
}

func (p *emailServiceV1ClientGRPCFactoryImpl) New(conf *def.Server) (EmailServiceV1Client, error) {
	cc, err := p.factory(conf)
	if err != nil {
		return nil, fmt.Errorf("create EmailServiceV1Client failed cause %s", err)
	}
	return &emailServiceV1Client{cc: cc}, nil
}

func RegisterEmailServiceV1ClientGRPCFactoryProvider(factory client.RegisterGRPCClientFactoryType) EmailServiceV1ClientGRPCFactory {
	return &emailServiceV1ClientGRPCFactoryImpl{factory: factory}
}

func (c *emailServiceV1Client) SendEmail(ctx context.Context, in *resources.SendEmailRequest, opts ...grpc.CallOption) (*resources.SendEmailReply, error) {
	out := new(resources.SendEmailReply)
	err := c.cc.Invoke(ctx, EmailServiceV1_SendEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceV1Server is the server API for EmailServiceV1 service.
// All implementations must embed UnimplementedEmailServiceV1Server
// for forward compatibility
type EmailServiceV1Server interface {
	SendEmail(context.Context, *resources.SendEmailRequest) (*resources.SendEmailReply, error)
	mustEmbedUnimplementedEmailServiceV1Server()
}

// Generate Injection
type registerEmailServiceV1ServerGRPCResult struct{}

func (*registerEmailServiceV1ServerGRPCResult) String() string {
	return "EmailServiceV1ServerGRPCServer"
}

func RegisterEmailServiceV1ServerGRPCProvider(newer interface{}) []interface{} {
	return []interface{}{
		// For provide dependency
		fx.Annotate(
			newer,
			fx.As(new(EmailServiceV1Server)),
		),
		// For create instance
		fx.Annotate(
			registerEmailServiceV1ServerProviderImpl,
			fx.As(new(fmt.Stringer)),
			fx.ResultTags(`group:"grpc_register"`),
		),
	}
}

// registerEmailServiceV1ServerProviderImpl use to trigger register
func registerEmailServiceV1ServerProviderImpl(s grpc.ServiceRegistrar, srv EmailServiceV1Server) *registerEmailServiceV1ServerGRPCResult {
	registerEmailServiceV1Server(s, srv)
	return &registerEmailServiceV1ServerGRPCResult{}
}

// UnimplementedEmailServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceV1Server struct {
}

func (UnimplementedEmailServiceV1Server) SendEmail(context.Context, *resources.SendEmailRequest) (*resources.SendEmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedEmailServiceV1Server) mustEmbedUnimplementedEmailServiceV1Server() {}

// UnsafeEmailServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceV1Server will
// result in compilation errors.
type UnsafeEmailServiceV1Server interface {
	mustEmbedUnimplementedEmailServiceV1Server()
}

func registerEmailServiceV1Server(s grpc.ServiceRegistrar, srv EmailServiceV1Server) {
	s.RegisterService(&EmailServiceV1_ServiceDesc, srv)
}

func _EmailServiceV1_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceV1Server).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailServiceV1_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceV1Server).SendEmail(ctx, req.(*resources.SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailServiceV1_ServiceDesc is the grpc.ServiceDesc for EmailServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.email.service.emailservicev1.EmailServiceV1",
	HandlerType: (*EmailServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _EmailServiceV1_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/email-service/v1/services/email.email.service.v1.proto",
}