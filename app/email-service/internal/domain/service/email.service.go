package services

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/domain/data"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/domain/repo"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/inject"
)

type emailRepo struct {
	dao *data.EmailDao
}

func NewEmailRepo(dao *data.EmailDao) repos.EmailRepo {

	return &emailRepo{
		dao: dao,
	}
}

func (p *emailRepo) Hello(ctx context.Context) (*data.EmailData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewEmailRepo)
}
