package main

import (
	"context"
	"path"

	_ "github.com/richhh7g/term-alarms/docs"

	api_config "github.com/richhh7g/term-alarms/internal/app/api/config"
	mongo_client "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo"
	"github.com/richhh7g/term-alarms/pkg/environment"
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

// @title Term Alarms
// @version 1.0.0
// @description API do Term Alarms para monitorar concorrentes que usam termos de marca em resultados patrocinados do Google.
// @contact.name Richhh7g
// @contact.url https://github.com/richhh7g
// @contact.email richhh7g@protonmail.com
// @license.name MIT
// @license.url https://github.com/richhh7g/term-alarms/blob/main/LICENSE
// @BasePath /
func main() {
	ctx := context.Background()

	databaseNameEnv := environment.Get[string]("MONGO_DB")
	client, err := mongo_client.NewMongoClient(ctx, &databaseNameEnv)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	serverConfig := api_config.NewServerConfig(ctx)
	err = serverConfig.Configure()
	if err != nil {
		panic(err)
	}
}
