package iam_service

/*
	本模块提供了 IAM 相关客户端的简易封装 需要使用的时候可以将对应的 GRPC/HTTPProvider 及 Service 同时注册到 Inject 模块中

	手动注入
	```go
	inject.Injection().InjectGRPCClient(NewIAMUserGRPCProvider)
	inject.Injection().Inject(NewIAMUserService)
	```

	也可以
	InjectIAMUserService(inject.Injection())

	然后就可以为自己添加所需使用的依赖
	IAMUserService 或 IAMUserServiceClient

*/

import (
	"context"

	contextutil "gitlab.lainuoniao.cn/eden-quan/go-biz-kit/context"
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	iamv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/iam-service/v1/resources"
	iamservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/iam-service/v1/services"
	errorv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/proto-common/v1/errors"
)

type Provider = func(creator interface{}) []interface{}

// NewIAMUserGRPCProvider 用于为 Injection 提供 RPC 客户端的账注入入口
func NewIAMUserGRPCProvider(creator interface{}) []interface{} {
	return iamservicev1.RegisterIAMUserServiceClientGRPCProvider(creator)
}

type IAMUserService struct {
	Client iamservicev1.IAMUserServiceClient
}

func NewIAMUserService(client iamservicev1.IAMUserServiceClient) *IAMUserService {
	return &IAMUserService{
		Client: client,
	}
}

// InjectIAMUserService 注入
func InjectIAMUserService(inj *injection.Injector) {
	inj.InjectGRPCClient(NewIAMUserGRPCProvider)
	inj.Inject(NewIAMUserService)
}

func (iam *IAMUserService) Get(ctx context.Context) (*iamv1.UserCheckTokenResponse, error) {
	token, ok := contextutil.GetAuthorizationToken(ctx)
	if !ok {
		return nil, errorv1.ERROR_UNAUTHORIZED.ToError("invalid token")
	}

	resp, err := iam.Client.CheckToken(ctx, &iamv1.UserCheckTokenRequest{
		Token: token,
	})

	return resp, errorv1.ERROR_UNAUTHORIZED.FromError(err)
}
