package email_service

import (
	"context"
	"gitlab.lainuoniao.cn/eden-quan/go-biz-kit/injection"
	emailerrorv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/email-service/v1/errors"
	emailv1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/email-service/v1/resources"
	emailservicev1 "gitlab.lainuoniao.cn/eden-quan/proto-service/api/email-service/v1/services"
)

type Email struct {
	Client emailservicev1.EmailServiceV1Client
}

func (e *Email) SendEmail(ctx context.Context, from, to, subject, body string) error {
	if from == "" || to == "" || subject == "" || body == "" {
		return emailerrorv1.EMAIL_ERROR_EMAIL_CONTENT_ERROR.ToError("invalid parameter")
	}

	resp, err := e.Client.SendEmail(ctx, &emailv1.SendEmailRequest{
		From:    from,
		To:      to,
		Subject: subject,
		Body:    body,
	})
	if err != nil || resp.Code != 0 {
		return emailerrorv1.EMAIL_ERROR_EMAIL_UNKNOWN.FromOrToError(err)
	}

	return nil
}

func NewEmail(client emailservicev1.EmailServiceV1Client) *Email {
	return &Email{
		Client: client,
	}
}

func InjectEmailService(inj *injection.Injector) {
	inj.InjectGRPCClient(emailservicev1.RegisterEmailServiceV1ClientGRPCProvider)
	inj.Inject(NewEmail)
}
