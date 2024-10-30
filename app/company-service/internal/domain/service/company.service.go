package services

import (
    "context"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/inject"
    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/domain/data"
    repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/domain/repo"
)

type companyRepo struct {
	dao *data.CompanyDao
}

func NewCompanyRepo(dao *data.CompanyDao) repos.CompanyRepo {

	return &companyRepo{
		dao: dao,
	}
}

func (p *companyRepo) Hello(ctx context.Context) (*data.CompanyData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewCompanyRepo)
}
