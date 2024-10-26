package main

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/cmd/export"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/conf"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/domain"
	_ "gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/impl/service"
)

func main() {
	injector := export.InjectBuiltIn()
	injector.DoIt(
		injection.WithGraph("./dep.dot"),
	)
}
