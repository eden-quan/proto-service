package messagecenter_service

import (
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	messagecenterservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/messagecenter-service/v1/services"
)

func InjectMessageCenterService(inj *injection.Injector) {
	inj.InjectGRPCClient(messagecenterservicev1.RegisterMessageCenterV1ClientGRPCProvider)
}
