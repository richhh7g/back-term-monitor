package main

import (
	"context"
	"fmt"
	"path"

	"github.com/eduardolat/goeasyi18n"
	mongo_client "github.com/richhh7g/term-alarms/infra/data/client/mongo"
	"github.com/richhh7g/term-alarms/pkg/environment"
	"github.com/richhh7g/term-alarms/pkg/localization"
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

	localizationService := localization.NewLocalization(goeasyi18n.NewI18n())
	localizationService.AddLanguages(map[localization.Language]string{
		localization.EN_US: path.Join("pkg", "localization", "locale", "en_us.locale.yml"),
		localization.PT_BR: path.Join("pkg", "localization", "locale", "pt_br.locale.yml"),
	})

	fmt.Println(environment.Get[string]("APP_ENV"))
	fmt.Println(localizationService.T("error.not_found", nil))
}
