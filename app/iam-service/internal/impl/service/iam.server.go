package serves

import (
	"github.com/go-kratos/kratos/v2/log"

	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	iamservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/iam-service/v1/services"
	iamconf "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/conf"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/inject"
)

// iam 服务实现示例，服务实现时无需关注 Http/gRPC, 只关注参数及返回结果
// 该类型的构造参数均由依赖注入框架提供
type iam struct {
	// 导入默认实现，默认实现会在调用对应的接口时返回 NotImplement 错误
	iamservicev1.UnimplementedIAMServiceServer
	log  *log.Helper               // log 为日志处理示例
	conf *iamconf.IamServiceConfig // conf 为配置中心示例
}

// NewIamArgs 示例在依赖较多时使用，用于将依赖集中声明到 struct 中，防止初始化函数参数过多难以维护
type NewIamArgs struct {
	injection.In // 需要集中依赖的结构体需要添加该 Embed 字段

	Logger log.Logger                // 日志库
	Conf   *iamconf.IamServiceConfig // 配置中心
}

// NewIamServer 创建一个 IamServer 的实例
// 需要注意的是如果准备同时提供 Http 及 gRPC 服务的情况下，他会在依赖注入框架中被创建出两个实例
// 如果需要两种服务使用同一个实例，请参考 iam-service.server.single.go
func NewIamServer(arg NewIamArgs) (iamservicev1.IAMServiceServer, error) {

	helper := log.NewHelper(log.With(arg.Logger, "module", "iam-service/impl/service/iam"))

	return &iam{
		conf: arg.Conf,
		log:  helper,
	}, nil
}

func init() {
	// 加入客户端注入信息, TODO: rpc 调用示例, 需要在应用中心配置后再移除注释
	// inject.Injection().InjectGRPCClient(iamservicev1.RegisterIamV1ClientGRPCProvider)
	// inject.Injection().InjectHTTPClient(iamservicev1.RegisterIamV1ClientHTTPProvider)

	// 加入客户端工厂注入信息，允许用户自己创建 RPC 的客户端, 以下为 HTTP 的示例
	// inject.Injection().Inject(iamservicev1.RegisterIamV1HTTPClientFactoryProvider)

	// 加入 Server 注入信息
	// TIPS: 因为这里将 NewIamServiceServer 作为不同的类型 Register 了两次，所以整个应用中会存在两个 IamServiceServer 实例
	// 如果是要使用同一个实例，则需要自行增加 sync.once 进行管理, 如下文中的 Singleton 示例
	annoGrpc := iamservicev1.RegisterIAMServiceServerGRPCProvider(NewIamServer)
	inject.Injection().InjectMany(annoGrpc)
}
