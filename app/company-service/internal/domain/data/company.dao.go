package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/inject"

)

type CompanyDao struct {
	db kit.MySQL
}

func NewCompanyDao(db kit.MySQL) *CompanyDao {
	return &CompanyDao{db: db}
}

func (p *CompanyDao) Query(ctx context.Context) (*CompanyData, error) {
	data := &CompanyData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewCompanyDao)
}
