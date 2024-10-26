package serves

import (
	"github.com/go-kratos/kratos/v2/log"

	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	filesystemservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/filesystem-service/v1/services"
)

// fileio 服务实现示例，服务实现时无需关注 Http/gRPC, 只关注参数及返回结果
// 该类型的构造参数均由依赖注入框架提供
type fileio struct {
	// 导入默认实现，默认实现会在调用对应的接口时返回 NotImplement 错误
	filesystemservicev1.UnimplementedFileIOApiV1Server
}

// NewFileioArgs 示例在依赖较多时使用，用于将依赖集中声明到 struct 中，防止初始化函数参数过多难以维护
type NewFileioArgs struct {
	injection.In // 需要集中依赖的结构体需要添加该 Embed 字段

	Logger log.Logger // 日志库
}

func NewFileioServer(arg NewFileioArgs) (filesystemservicev1.FileIOApiV1Server, error) {
	return &fileio{}, nil
}

func init() {}
