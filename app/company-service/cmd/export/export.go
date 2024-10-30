package export

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	config "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/config/inject"
	servers "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/server/inject"
	setup "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/setup/inject"

	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/domain"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/impl/service"
    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/inject"
)

func InjectBuiltIn() *injection.Injector {
    injector := inject.Injection()

	config.InjectIns(injector)
	setup.InjectIns(injector)
	servers.InjectIns(injector)

	return injector
}