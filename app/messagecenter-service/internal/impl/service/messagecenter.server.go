package serves

import (
    "context"
    "fmt"

    "github.com/go-kratos/kratos/v2/log"

    kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
    "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
    contextutil "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/context"
    "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/message"


    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/inject"
    messagecenterv1err "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/errors"
	messagecenterv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/resources"
	messagecenterservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/services"
    messagecenterconf "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/conf"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/domain/data"
)

// messagecenter 服务实现示例，服务实现时无需关注 Http/gRPC, 只关注参数及返回结果
// 该类型的构造参数均由依赖注入框架提供
type messagecenter struct {
    // 导入默认实现，默认实现会在调用对应的接口时返回 NotImplement 错误
   messagecenterservicev1.UnimplementedMessagecenterV1Server

	mq         kit.MessageQueue               // mq 用于发送及消费消息队列的消息
	dao        *data.MessagecenterDao // dao 为数据库查询示例
	log        *log.Helper                     // log 为日志处理示例

	conf       *messagecenterconf.MessagecenterServiceConfig // conf 为配置中心示例
	consume    message.Consumer // consume 为消费消息队列的示例
	producer   message.Producer // producer 为发布消息队列消息的示例
	// httpClient messagecenterservicev1.MessagecenterV1HTTPClient // TODO: rpc 调用示例, 需要在应用中心配置后再移除注释
}


// NewMessagecenterArgs 示例在依赖较多时使用，用于将依赖集中声明到 struct 中，防止初始化函数参数过多难以维护
type NewMessagecenterArgs struct {
	injection.In // 需要集中依赖的结构体需要添加该 Embed 字段

	Logger        log.Logger                        // 日志库
	Dao           *data.MessagecenterDao    // 数据访问库
	MQ            kit.MessageQueue                  // 消息队列
	Conf          *messagecenterconf.MessagecenterServiceConfig              // 配置中心
	// HttpMessagecenterClient messagecenterservicev1.MessagecenterV1HTTPClient // HTTP Client TODO: rpc 调用示例, 需要在应用中心配置后再移除注释
	// HttpMessagecenterClientFactory messagecenterservicev1.MessagecenterV1HTTPClientFactory // HTTP Factory 示例，允许用户通过 Factory 自主创建 RPC 客户端
}

// NewMessagecenterServer 创建一个 MessagecenterServer 的实例
// 需要注意的是如果准备同时提供 Http 及 gRPC 服务的情况下，他会在依赖注入框架中被创建出两个实例
// 如果需要两种服务使用同一个实例，请参考 messagecenter-service.server.single.go
func NewMessagecenterServer(arg NewMessagecenterArgs) (messagecenterservicev1.MessagecenterV1Server, error) {

	helper := log.NewHelper(log.With(arg.Logger, "module", "messagecenter-service/impl/service/messagecenter"))

	// 可通过 Factory 自主构建客户端
	// client, err := arg.HttpMessagecenterClientFactory.New(&def.Server{})
	// helper.Warn(err) // err occur cause we don't provide the Http Config

    var pro message.Producer = nil
    var con message.Consumer = nil
    var err error = nil

    // 只有开启了 MQ 配置，才会尝试创建消费者及生产者
    if arg.MQ.Get() != nil {
        // 根据配置中心的消息队列配置获取 Consumer
        con, err = arg.MQ.Get().Consumer(
            message.NewConsumerConfig(
                &arg.Conf.DemoQueue,    // 获取配置中心的队列配置
                &arg.Conf.DemoConsumer, // 获取配置中心的消费者配置
            ),
        )

        if err != nil {
            helper.Infof("con(%s): %s\n", con, err)
            return nil, err
        }

        // 根据配置中心的消息队列的 Exchange 及 Queue 创建一个生产者
        pro, err = arg.MQ.Get().Producer(
            message.NewProducerConfig(
                &arg.Conf.DemoExchange, // 获取配置中心的交换机配置
                &arg.Conf.DemoQueue,    // 获取配置中心的队列配置
            ),
        )

        if err != nil {
            helper.Error(err)
            return nil, err
        }
	}

	return &messagecenter{
		conf:       arg.Conf,
		mq:         arg.MQ,
		dao:        arg.Dao,
		log:        helper,
		consume:    con,
		producer:   pro,
		// client:     arg.MessagecenterClient,
		// httpClient: arg.HttpMessagecenterClient,
	}, nil
}

func (s *messagecenter) Pong(ctx context.Context, req *messagecenterv1.MessagecenterPingReq) (*messagecenterv1.MessagecenterPingResp, error) {

    // 可以从 Context 中获取用户信息
	token, _ := contextutil.GetAuthorizationToken(ctx)
	s.log.WithContext(ctx).Warn("token: ", token)

	var err error = nil
	if req.GetMessage() == "error" {
	    // 错误处理示例, 可以使用自定义来包装原生的错误，获取堆栈及错误码等信息
	    err = messagecenterv1err.MESSAGECENTER_ERROR_MESSAGECENTER_CONTENT_ERROR.FromErrorf(
	        fmt.Errorf("oh no"),
	        "content error",
	    )
	}

	return &messagecenterv1.MessagecenterPingResp{
		Data:  "",
	}, err
}

// Ping 为日志库 / 错误处理等基础示例
func (s *messagecenter) Ping(ctx context.Context, in *messagecenterv1.MessagecenterPingReq) (*messagecenterv1.MessagecenterPingResp, error) {


    // 可从上下文中获取用户 Token 信息
	token, _ := contextutil.GetAuthorizationToken(ctx)

    // 日志示例, 将 Token 通过日志打印
	s.log.WithContext(ctx).Warn("token: ", token)

    // TODO: rpc 调用示例, 需要在应用中心配置后再移除注释
	// resp, err := s.client.Pong(ctx, in)
	// s.log.WithContext(ctx).Info("sdk call: ", resp.String(), err)

	return &messagecenterv1.MessagecenterPingResp{
		Data: "Received Message : " + in.GetMessage(),
	}, nil

}

// PublishMessage 为消息队列生产消息的示例
func (s *messagecenter) PublishMessage(ctx context.Context, in *messagecenterv1.MessagecenterPublishReq) (*messagecenterv1.MessagecenterPublishResp, error) {

	// 需要发送的消息格式
	type tmp struct {
		Id int
		Name string
	}

	count := 0
	for i, m := 0, 10; i < m; i++ {
		err := s.producer.Push(ctx, tmp{Id: i, Name: in.Msg})
		if err != nil {
		    return nil, messagecenterv1err.MESSAGECENTER_ERROR_MESSAGECENTER_CONTENT_ERROR.FromError(err)
		}

        count++
	}

	return &messagecenterv1.MessagecenterPublishResp {
	    N: int32(count),
	}, nil
}

// ConsumeMessage 为消费消息队列的示例
func (s *messagecenter) ConsumeMessage(ctx context.Context, in *messagecenterv1.MessagecenterConsumeReq) (*messagecenterv1.MessagecenterConsumeResp, error) {

    // 示例 - 消费 N 条消息
	count := int32(0)
	for i, m := 0, int(in.N); i < m; i++ {
		_, err := s.consume.Pull(ctx)
		if err != nil {
			s.log.Error(err)
			break
		}
		count += 1
	}

    // 示例 - 消费所有消息，一般会在后台的 Goroutine 中使用
	// for msg := range con.Puller() {
	// 	 s.log.Debug(msg)
	// }

	return &messagecenterv1.MessagecenterConsumeResp {
	    N: count,
	}, nil
}

// DataQuery 为数据库查询的示例
func (s *messagecenter) DataQuery(ctx context.Context)  {

	token, _ := contextutil.GetAuthorizationToken(ctx)
	s.log.WithContext(ctx).Warn("token: ", token)

	dataFromDb, err := s.dao.Query(ctx)
	fmt.Print(dataFromDb, err)
}

func init() {
	// 加入客户端注入信息, TODO: rpc 调用示例, 需要在应用中心配置后再移除注释
	// inject.Injection().InjectGRPCClient(messagecenterservicev1.RegisterMessagecenterV1ClientGRPCProvider)
	// inject.Injection().InjectHTTPClient(messagecenterservicev1.RegisterMessagecenterV1ClientHTTPProvider)

	// 加入客户端工厂注入信息，允许用户自己创建 RPC 的客户端, 以下为 HTTP 的示例
	// inject.Injection().Inject(messagecenterservicev1.RegisterMessagecenterV1HTTPClientFactoryProvider)

	// 加入 Server 注入信息
	// TIPS: 因为这里将 NewMessagecenterServiceServer 作为不同的类型 Register 了两次，所以整个应用中会存在两个 MessagecenterServiceServer 实例
	// 如果是要使用同一个实例，则需要自行增加 sync.once 进行管理, 如下文中的 Singleton 示例
	// annoHttp := messagecenterservicev1.RegisterMessagecenterV1HTTPServerProvider(NewMessagecenterServiceServer)
	// annoGrpc := messagecenterservicev1.RegisterMessagecenterV1ServerGRPCProvider(NewMessagecenterServiceServer)

	// singleton version
	annoHttp := messagecenterservicev1.RegisterMessagecenterV1ServerHTTPProvider(NewSingletonMessagecenterServer)
	annoGrpc := messagecenterservicev1.RegisterMessagecenterV1ServerGRPCProvider(NewSingletonMessagecenterServer)

	// 注入 SQL Inject 的实现函数

	inject.Injection().InjectMany(annoHttp, annoGrpc)
}
