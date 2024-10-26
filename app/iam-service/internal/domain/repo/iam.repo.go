package repos

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/iam-service/internal/domain/data"
)

type IamRepo interface {
	Hello(ctx context.Context) (*data.IamData, error)
}
