package main

import (

    "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/cmd/export"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/conf"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/domain"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/impl/service"

)

func main() {
    injector := export.InjectBuiltIn()
	injector.DoIt(
		injection.WithGraph("./dep.dot"),
	)
}
