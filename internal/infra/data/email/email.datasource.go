package email_datasource

import (
	"context"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
	resend_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/email/resend"
	email_template "github.com/richhh7g/back-term-monitor/internal/infra/template/email"
	"github.com/richhh7g/back-term-monitor/pkg/localization"
)

type SendPotentialCompetitorsParams struct {
	Email string
	Terms []*competitor_model.CompetitorTermBaseModel
}

type Email interface {
	SendPotentialCompetitors(ctx context.Context, params *SendPotentialCompetitorsParams) error
}

type EmailImpl struct {
	client   resend_client.Resend
	localize localization.Localization
}

func NewEmailDataSource(client resend_client.Resend, localize localization.Localization) Email {
	return &EmailImpl{
		client:   client,
		localize: localize,
	}
}

func (d *EmailImpl) SendPotentialCompetitors(ctx context.Context, input *SendPotentialCompetitorsParams) error {
	subject := d.localize.T("email.send_potential_competitors.subject", nil)

	emailHtml, err := email_template.ParseTemplate(email_template.PotentialCompetitorsTemplate, input.Terms)
	if err != nil {
		return err
	}

	_, err = d.client.Send(ctx, &resend_client.SendEmailRequest{
		To:      []string{input.Email},
		Html:    *emailHtml,
		Subject: subject,
	})
	if err != nil {
		return err
	}

	return nil
}
