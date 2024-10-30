package repos

import (
    "context"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/messagecenter-service/internal/domain/data"
)

type MessagecenterRepo interface {
	Hello(ctx context.Context) (*data.MessagecenterData, error)
}
