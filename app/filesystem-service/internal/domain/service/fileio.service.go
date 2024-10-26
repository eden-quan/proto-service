package services

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/domain/data"
	repos "gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/domain/repo"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/inject"
)

type fileioRepo struct {
	dao *data.FileioDao
}

func NewFileioRepo(dao *data.FileioDao) repos.FileioRepo {

	return &fileioRepo{
		dao: dao,
	}
}

func (p *fileioRepo) Hello(ctx context.Context) (*data.FileioData, error) {
	return p.dao.Query(ctx)
}

func init() {
	inject.Injection().Inject(NewFileioRepo)
}
