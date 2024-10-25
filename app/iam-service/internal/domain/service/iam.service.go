package services

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/domain/data"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/domain/repo"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/inject"
)

type iamRepo struct {
	dao *data.IamDao
}

func NewIamRepo(dao *data.IamDao) repos.IamRepo {

	return &iamRepo{
		dao: dao,
	}
}

func (p *iamRepo) Hello(ctx context.Context) (*data.IamData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewIamRepo)
}
