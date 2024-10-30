// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             v5.26.1
// source: api/messagecenter-service/v1/services/messagecenter.messagecenter.service.v1.proto

package messagecenterservicev1

import (
	context "context"
	fmt "fmt"

	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	client "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/client"
	def "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/def"
	resources "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/resources"
	fx "go.uber.org/fx"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(fmt.Stringer)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1
const _ = fx.Version

var _ = new(def.Server)
var _ = new(client.RegisterGRPCClientFactoryType)

const OperationMessagecenterV1ConsumeMessage = "/api.messagecenter.service.messagecenterservicev1.MessagecenterV1/ConsumeMessage"
const OperationMessagecenterV1Ping = "/api.messagecenter.service.messagecenterservicev1.MessagecenterV1/Ping"
const OperationMessagecenterV1PublishMessage = "/api.messagecenter.service.messagecenterservicev1.MessagecenterV1/PublishMessage"

type MessagecenterV1HTTPServer interface {
	// ConsumeMessage 消息队列 - 消费消息示例
	ConsumeMessage(context.Context, *resources.MessagecenterConsumeReq) (*resources.MessagecenterConsumeResp, error)
	// Ping 测试 http 接口
	Ping(context.Context, *resources.MessagecenterPingReq) (*resources.MessagecenterPingResp, error)
	// PublishMessage 消息队列 - 发布消息示例
	PublishMessage(context.Context, *resources.MessagecenterPublishReq) (*resources.MessagecenterPublishResp, error)
}

type registerMessagecenterV1HTTPResult struct{}

func (*registerMessagecenterV1HTTPResult) String() string {
	return "MessagecenterV1HTTPServer"
}

func RegisterMessagecenterV1ServerHTTPProvider(newer interface{}) []interface{} {
	return []interface{}{
		// For provide dependency
		fx.Annotate(
			newer,
			fx.As(new(MessagecenterV1HTTPServer)),
		),
		// For create instance
		fx.Annotate(
			registerMessagecenterV1HTTPProviderImpl,
			fx.As(new(fmt.Stringer)),
			fx.ResultTags(`group:"http_register"`),
		),
	}
}

// registerMessagecenterV1ProviderImpl use to trigger register
func registerMessagecenterV1HTTPProviderImpl(s *http.Server, srv MessagecenterV1HTTPServer) *registerMessagecenterV1HTTPResult {
	registerMessagecenterV1HTTPServer(s, srv)
	return &registerMessagecenterV1HTTPResult{}
}

func registerMessagecenterV1HTTPServer(s *http.Server, srv MessagecenterV1HTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/messagecenterv1/ping/{message}", _MessagecenterV1_Ping0_HTTP_Handler(srv))
	r.POST("/api/v1/messagecenterv1/consume", _MessagecenterV1_ConsumeMessage0_HTTP_Handler(srv))
	r.POST("/api/v1/messagecenterv1/publish", _MessagecenterV1_PublishMessage0_HTTP_Handler(srv))
}

func _MessagecenterV1_Ping0_HTTP_Handler(srv MessagecenterV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.MessagecenterPingReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessagecenterV1Ping)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*resources.MessagecenterPingReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.MessagecenterPingResp)
		return ctx.Result(200, reply)
	}
}

func _MessagecenterV1_ConsumeMessage0_HTTP_Handler(srv MessagecenterV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.MessagecenterConsumeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessagecenterV1ConsumeMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ConsumeMessage(ctx, req.(*resources.MessagecenterConsumeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.MessagecenterConsumeResp)
		return ctx.Result(200, reply)
	}
}

func _MessagecenterV1_PublishMessage0_HTTP_Handler(srv MessagecenterV1HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in resources.MessagecenterPublishReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMessagecenterV1PublishMessage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PublishMessage(ctx, req.(*resources.MessagecenterPublishReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*resources.MessagecenterPublishResp)
		return ctx.Result(200, reply)
	}
}

type MessagecenterV1HTTPClient interface {
	ConsumeMessage(ctx context.Context, req *resources.MessagecenterConsumeReq, opts ...http.CallOption) (rsp *resources.MessagecenterConsumeResp, err error)
	Ping(ctx context.Context, req *resources.MessagecenterPingReq, opts ...http.CallOption) (rsp *resources.MessagecenterPingResp, err error)
	PublishMessage(ctx context.Context, req *resources.MessagecenterPublishReq, opts ...http.CallOption) (rsp *resources.MessagecenterPublishResp, err error)
	RegisterNameForDiscover() string
}

func registerMessagecenterV1ClientHTTPNameProvider() []string {
	return []string{"messagecenter-service/v1", "http"}
}

func RegisterMessagecenterV1ClientHTTPProvider(creator interface{}) []interface{} {
	return []interface{}{
		fx.Annotate(
			newMessagecenterV1HTTPClient,
			fx.As(new(MessagecenterV1HTTPClient)),
			fx.ParamTags(`name:"messagecenter-service/v1/http/messagecenterV1"`),
		),
		fx.Annotate(
			creator,
			// fx.As(new(*http.Client)),
			fx.ParamTags(`name:"messagecenter-service/v1/http/name/messagecenterV1"`),
			fx.ResultTags(`name:"messagecenter-service/v1/http/messagecenterV1"`),
		),
		fx.Annotate(
			registerMessagecenterV1ClientHTTPNameProvider,
			fx.ResultTags(`name:"messagecenter-service/v1/http/name/messagecenterV1"`),
		),
	}
}

type MessagecenterV1HTTPClientFactory interface {
	New(conf *def.Server) (MessagecenterV1HTTPClient, error)
}

type _MessagecenterV1HTTPClientFactoryImpl struct {
	factory client.RegisterHTTPClientFactoryType
}

func (p *_MessagecenterV1HTTPClientFactoryImpl) New(conf *def.Server) (MessagecenterV1HTTPClient, error) {
	cc, err := p.factory(conf)
	if err != nil {
		return nil, fmt.Errorf("create MessagecenterV1HTTPClient failed cause %s", err)
	}

	return &_MessagecenterV1HTTPClientImpl{cc: cc}, nil
}

func RegisterMessagecenterV1HTTPClientFactoryProvider(factory client.RegisterHTTPClientFactoryType) MessagecenterV1HTTPClientFactory {
	return &_MessagecenterV1HTTPClientFactoryImpl{factory: factory}
}

type _MessagecenterV1HTTPClientImpl struct {
	cc *http.Client
}

func newMessagecenterV1HTTPClient(client *http.Client) MessagecenterV1HTTPClient {
	return &_MessagecenterV1HTTPClientImpl{client}
}

func (c *_MessagecenterV1HTTPClientImpl) RegisterNameForDiscover() string {
	return "messagecenter-service/v1"
}

func (c *_MessagecenterV1HTTPClientImpl) ConsumeMessage(ctx context.Context, in *resources.MessagecenterConsumeReq, opts ...http.CallOption) (*resources.MessagecenterConsumeResp, error) {
	var out resources.MessagecenterConsumeResp
	pattern := "/api/v1/messagecenterv1/consume"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessagecenterV1ConsumeMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *_MessagecenterV1HTTPClientImpl) Ping(ctx context.Context, in *resources.MessagecenterPingReq, opts ...http.CallOption) (*resources.MessagecenterPingResp, error) {
	var out resources.MessagecenterPingResp
	pattern := "/api/v1/messagecenterv1/ping/{message}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessagecenterV1Ping))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *_MessagecenterV1HTTPClientImpl) PublishMessage(ctx context.Context, in *resources.MessagecenterPublishReq, opts ...http.CallOption) (*resources.MessagecenterPublishResp, error) {
	var out resources.MessagecenterPublishResp
	pattern := "/api/v1/messagecenterv1/publish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMessagecenterV1PublishMessage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}