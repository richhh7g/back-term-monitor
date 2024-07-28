package main

import (
	"context"
	"fmt"
	"path"

	"github.com/eduardolat/goeasyi18n"

	smtp_client "github.com/richhh7g/term-alarms/internal/infra/data/client/email/smtp"
	mongo_client "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo"
	mongo_document "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/document"
	mongo_repository "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/repository"
	"github.com/richhh7g/term-alarms/pkg/environment"
	"github.com/richhh7g/term-alarms/pkg/localization"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	err := environment.NewEnvLoader(&environment.EnvLoaderParams{
		Type: "env",
		File: ".env",
		Path: path.Join("cmd", "server"),
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	databaseNameEnv := environment.Get[string]("MONGO_DB")
	client, err := mongo_client.NewMongoClient(ctx, &databaseNameEnv)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	alarmRepository := mongo_repository.NewAlarmRepository(client)
	alarmRepository.Create(ctx, &mongo_document.Alarm{
		ID:    primitive.ObjectID{},
		Email: "test@test.com",
		Tags:  []string{"tag1", "tag2"},
	})
	alarm, err := alarmRepository.FindOneByEmail(ctx, "test@test.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(alarm)

	smtpClient, err := smtp_client.NewSmtp(true, false)
	if err != nil {
		panic(err)
	}
	err = smtpClient.Send(ctx, &smtp_client.SendParams{
		To:      []string{"richhh7g@protonmail.com"},
		Subject: "Email de teste pelo SMTP",
		Html:    "<h1>Esse é um email de teste pelo SMTP</h1>",
	})
	if err != nil {
		panic(err)
	}

	localizationService := localization.NewLocalization(goeasyi18n.NewI18n())
	localizationService.AddLanguages(map[localization.Language]string{
		localization.EN_US: path.Join("pkg", "localization", "locale", "en_us.locale.yml"),
		localization.PT_BR: path.Join("pkg", "localization", "locale", "pt_br.locale.yml"),
	})

	fmt.Println(environment.Get[string]("APP_ENV"))
	fmt.Println(localizationService.T("error.not_found", nil))
}
