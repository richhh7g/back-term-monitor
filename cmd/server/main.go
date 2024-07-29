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
