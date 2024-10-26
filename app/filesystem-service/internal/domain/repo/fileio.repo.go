package repos

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/filesystem-service/internal/domain/data"
)

type FileioRepo interface {
	Hello(ctx context.Context) (*data.FileioData, error)
}
