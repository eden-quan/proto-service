package services

import (
    "context"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/inject"
    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/domain/data"
    repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/domain/repo"
)

type messagecenterRepo struct {
	dao *data.MessagecenterDao
}

func NewMessagecenterRepo(dao *data.MessagecenterDao) repos.MessagecenterRepo {

	return &messagecenterRepo{
		dao: dao,
	}
}

func (p *messagecenterRepo) Hello(ctx context.Context) (*data.MessagecenterData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewMessagecenterRepo)
}
