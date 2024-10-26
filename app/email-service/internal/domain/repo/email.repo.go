package repos

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/email-service/internal/domain/data"
)

type EmailRepo interface {
	Hello(ctx context.Context) (*data.EmailData, error)
}
