package serves

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	smsservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/sms-service/v1/services"
)

// sms 服务实现示例，服务实现时无需关注 Http/gRPC, 只关注参数及返回结果
// 该类型的构造参数均由依赖注入框架提供
type sms struct {
	// 导入默认实现，默认实现会在调用对应的接口时返回 NotImplement 错误
	smsservicev1.UnimplementedSMSServiceV1Server
}

// NewSmsArgs 示例在依赖较多时使用，用于将依赖集中声明到 struct 中，防止初始化函数参数过多难以维护
type NewSmsArgs struct {
	injection.In // 需要集中依赖的结构体需要添加该 Embed 字段
}

func NewSmsServer(arg NewSmsArgs) (smsservicev1.SMSServiceV1Server, error) {
	return &sms{}, nil
}

func init() {
}
