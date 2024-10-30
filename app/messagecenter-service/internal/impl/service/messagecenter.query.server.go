package serves

import (
    "context"

    "github.com/go-kratos/kratos/v2/log"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/inject"

	messagecenterservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/services"
)

// messagecenterSQLQuery 数据库插件实现示例, 如果
// 该类型的构造参数均由依赖注入框架提供
type messagecenterSQLQuery struct {
    // 导入默认实现，默认实现会在调用对应的接口时返回 NotImplement 错误
   messagecenterservicev1.UnimplementedMessagecenterV1SQLQueryServer
	log        *log.Helper                     // log 为日志处理示例
}

// NewMessagecenterSQLQueryServer 创建一个新的 MessagecenterSQLQueryServer 实例，
// 将该函数注册到依赖注入矿建后，该函数参数中定义的类型，会由依赖注入框架提供
func NewMessagecenterSQLQueryServer(logger log.Logger) (messagecenterservicev1.MessagecenterV1SQLQueryServer, error) {
	q := &messagecenterSQLQuery {
	    log: log.NewHelper(logger),
	}

	// 注册 Inject 类型的处理函数，这样服务在收到请求时会按照 SQL 插件中定义的顺序来调用该处理函数
	messagecenterservicev1.RegisterMessagecenterV1SQLQueryChainSQLActionInjectMethodAdjustName(q.InjectFunc)
	return q, nil
}

// InjectFunc 为注入函数的示例，从参数中可以获取对应的上下文，本次请求的参数及返回值，注入函数可以对这些信息进行修改及调整
func (q *messagecenterSQLQuery) InjectFunc(ctx context.Context, arg *messagecenterservicev1.MessagecenterV1SQLQueryChainSQLActionInjectArg) (context.Context, error) {
	q.log.Warn(arg.Resp)
	q.log.Warn(ctx)
	return ctx, nil
}

func init() {
    // 注册服务的处理器，即本示例中的 messagecenterSQLQuery, 用于提供 HTTP 服务
	annoHttp := messagecenterservicev1.RegisterMessagecenterV1SQLQueryServerHTTPProvider(NewMessagecenterSQLQueryServer)
	// 注册 SQL 插件生成的逻辑到依赖注入框架
	annoSql := messagecenterservicev1.RegisterMessagecenterV1SQLQuerySQLAction()

	inject.Injection().InjectMany(annoHttp, annoSql)
}
