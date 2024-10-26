package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/inject"
)

type SmsDao struct {
	db kit.MySQL
}

func NewSmsDao(db kit.MySQL) *SmsDao {
	return &SmsDao{db: db}
}

func (p *SmsDao) Query(ctx context.Context) (*SmsData, error) {
	data := &SmsData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewSmsDao)
}
