package sms_service

import (
	"context"
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	smserrorv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/sms-service/v1/errors"
	smsv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/sms-service/v1/resources"
	smsservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/sms-service/v1/services"
)

type SMS struct {
	Client smsservicev1.SMSServiceV1Client
}

func (e *SMS) SendEmail(ctx context.Context, phone, content, templateCode string) error {
	if phone == "" || content == "" || templateCode == "" {
		return smserrorv1.SMS_ERROR_SMS_CONTENT_MISSING.ToError("invalid parameter")
	}

	resp, err := e.Client.SendSMS(ctx, &smsv1.SendSMSRequest{
		Phone:        phone,
		Content:      content,
		TemplateCode: templateCode,
	})

	if err != nil || resp.Code != 0 {
		return smserrorv1.SMS_ERROR_SMS_UNKNOWN.FromOrToError(err)
	}

	return nil
}

func NewSMS(client smsservicev1.SMSServiceV1Client) *SMS {
	return &SMS{
		Client: client,
	}
}

func InjectSMSService(inj *injection.Injector) {
	inj.InjectGRPCClient(smsservicev1.RegisterSMSServiceV1ClientGRPCProvider)
	inj.Inject(NewSMS)
}
