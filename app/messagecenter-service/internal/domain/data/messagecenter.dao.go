package data

import (
	"context"

	kit "gitlab.lainuoniao.cn/eden-quan/go-biz-kit"
    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/inject"

)

type MessagecenterDao struct {
	db kit.MySQL
}

func NewMessagecenterDao(db kit.MySQL) *MessagecenterDao {
	return &MessagecenterDao{db: db}
}

func (p *MessagecenterDao) Query(ctx context.Context) (*MessagecenterData, error) {
	data := &MessagecenterData{}
	err := p.db.Get().GetContext(ctx, data, `SELECT 'hello world' as name;`)
	return data, err
}

func init() {
	inject.Injection().Inject(NewMessagecenterDao)
}
