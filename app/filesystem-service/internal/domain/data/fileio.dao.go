package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/inject"
)

type FileioDao struct {
	db kit.MySQL
}

func NewFileioDao(db kit.MySQL) *FileioDao {
	return &FileioDao{db: db}
}

func (p *FileioDao) Query(ctx context.Context) (*FileioData, error) {
	data := &FileioData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewFileioDao)
}
