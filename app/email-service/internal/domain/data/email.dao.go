package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/inject"
)

type EmailDao struct {
	db kit.MySQL
}

func NewEmailDao(db kit.MySQL) *EmailDao {
	return &EmailDao{db: db}
}

func (p *EmailDao) Query(ctx context.Context) (*EmailData, error) {
	data := &EmailData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewEmailDao)
}
