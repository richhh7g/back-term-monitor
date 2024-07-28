package smtp_client

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"

	"github.com/richhh7g/term-alarms/pkg/environment"
)

type Smtp interface {
	Send(ctx context.Context, params *SendParams) error
}

type SmtpImpl struct {
	driver *smtp.Client
}

func NewSmtp(authentication bool, tlsConnection bool) (Smtp, error) {
	smtpPort := environment.Get[string]("SMTP_PORT")
	smtpHost := environment.Get[string]("SMTP_HOST")

	serverName := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	conn, err := net.Dial("tcp", serverName)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return nil, err
	}

	if tlsConnection {
		err := client.StartTLS(&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		})
		if err != nil {
			return nil, err
		}
	}

	if authentication {
		smtpUser := environment.Get[string]("SMTP_USER")
		smtpPass := environment.Get[string]("SMTP_PASS")

		auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
		client.Auth(auth)
	}

	return &SmtpImpl{
		driver: client,
	}, nil
}

func (c *SmtpImpl) Send(ctx context.Context, params *SendParams) error {
	defer c.driver.Close()

	if params.From == "" {
		params.From = environment.Get[string]("MAIL_SENDER")
	}

	if err := c.driver.Mail(params.From); err != nil {
		log.Panic(err)
	}

	for _, addr := range params.To {
		if err := c.driver.Rcpt(addr); err != nil {
			return err
		}
	}

	writerMessage, err := c.driver.Data()
	if err != nil {
		return err
	}
	defer writerMessage.Close()

	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(params.To, ", ")))
	msg.WriteString(fmt.Sprintf("From: %s\r\n", params.From))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", params.Subject))

	if params.Html != "" {
		msg.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
		msg.WriteString("\r\n")
		msg.WriteString(params.Html)
	} else if params.Text != "" {
		msg.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
		msg.WriteString("\r\n")
		msg.WriteString(params.Text)
	} else {
		return fmt.Errorf("no email content provided")
	}

	if _, err = writerMessage.Write(msg.Bytes()); err != nil {
		return err
	}

	defer c.driver.Quit()

	return nil
}
