package conf

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config"
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/def"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/inject"
)

// IamServiceConfig 在全局基础配置的前提下增加了自定义配置信息 Hello,
// 并且在此之上，通过 `conf_service` 增加了对基础配置自定义的能力
type IamServiceConfig struct {
	*def.Configuration `conf_service:"iam-service"`
	// Server            *def.Server        `conf_path:"/registry/iam-service/v1/config"` // TODO: 配置完配置中心后移除注释, 实现自定义服务器监听配置
	DemoExchange      def.ExchangeConfig `conf_path:"/message/exchange/ping"`
	DemoQueue         def.QueueConfig    `conf_path:"/message/queue/ping"`
	DemoConsumer      def.ConsumeConfig  `conf_path:"/message/consumer/ping"`
}

// NewIamServiceConfig 提供新的自定义配置信息
func NewIamServiceConfig(repo config.ConfigureWatcherRepo) (*IamServiceConfig, error) {
	iam := &IamServiceConfig{
	    Configuration: &def.Configuration{
	        Server: &def.Server{},
	    },
	    // Server: &def.Server{}, // TODO: 配置完配置中心后移除注释, 实现自定义服务器监听配置
	}

	err := repo.LoadAndStart(iam)
	return iam, err
}

// NewConfigReplacer 允许应用使用自定义配置信息覆盖默认的全局配置
func NewConfigReplacer(iam *IamServiceConfig) *def.Configuration {
    iam.Configuration.Server = iam.Server
	return iam.Configuration
}

func init() {
	inject.Injection().Inject(NewIamServiceConfig)
	inject.Injection().Inject(NewConfigReplacer)
}
