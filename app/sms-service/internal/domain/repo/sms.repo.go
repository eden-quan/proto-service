package repos

import (
	"context"

	"gitlab.lainuoniao.cn/eden-quan/proto-service/app/sms-service/internal/domain/data"
)

type SmsRepo interface {
	Hello(ctx context.Context) (*data.SmsData, error)
}
