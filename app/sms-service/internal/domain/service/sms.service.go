package services

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/domain/data"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/domain/repo"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/inject"
)

type smsRepo struct {
	dao *data.SmsDao
}

func NewSmsRepo(dao *data.SmsDao) repos.SmsRepo {

	return &smsRepo{
		dao: dao,
	}
}

func (p *smsRepo) Hello(ctx context.Context) (*data.SmsData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewSmsRepo)
}
