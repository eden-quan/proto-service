package repos

import (
    "context"

    "gitlab.lainuoniao.cn/eden-quan/proto-service/app/company-service/internal/domain/data"
)

type CompanyRepo interface {
	Hello(ctx context.Context) (*data.CompanyData, error)
}
