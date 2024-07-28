package resend_client

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/richhh7g/term-alarms/pkg/environment"
)

type Resend interface {
	Send(ctx context.Context, params *SendEmailRequest) (*SendEmailResponse, error)
}
type ResendImpl struct {
	driver *resty.Client
}

func NewResend() Resend {
	client := resty.New()
	client.BaseURL = environment.Get[string]("RESEND_BASE_URL")
	client.Header.Add("Content-Type", "application/json")
	bearerToken := fmt.Sprintf("Bearer %s", environment.Get[string]("RESEND_API_KEY"))
	client.Header.Add("Authorization", bearerToken)

	return &ResendImpl{
		driver: client,
	}
}

func (c *ResendImpl) Send(ctx context.Context, params *SendEmailRequest) (*SendEmailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var response SendEmailResponse

	if params.From == "" {
		mailAddress := mail.Address{
			Name:    environment.Get[string]("MAIL_NAME"),
			Address: environment.Get[string]("MAIL_SENDER"),
		}

		params.From = mailAddress.String()
	}

	resp, err := c.driver.R().
		SetContext(ctx).
		SetBody(params).
		SetResult(&response).
		Post("/emails")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode())
	}

	return &response, nil
}
