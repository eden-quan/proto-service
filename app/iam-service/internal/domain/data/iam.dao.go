package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/inject"
)

type IamDao struct {
	db kit.MySQL
}

func NewIamDao(db kit.MySQL) *IamDao {
	return &IamDao{db: db}
}

func (p *IamDao) Query(ctx context.Context) (*IamData, error) {
	data := &IamData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewIamDao)
}
