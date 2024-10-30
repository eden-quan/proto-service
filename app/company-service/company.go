package filesystem_service

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	companyservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/company-service/v1/services"
)

// InjectCompanyService 注入 CompanyService 等依赖到 injector 中
func InjectCompanyService(inj *injection.Injector) {
	inj.InjectGRPCClient(companyservicev1.RegisterCompanyV1ClientGRPCProvider)
}
