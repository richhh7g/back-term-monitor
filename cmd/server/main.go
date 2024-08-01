package main

import (
	"context"
	"path"

	_ "github.com/richhh7g/back-term-monitor/docs"

	api_config "github.com/richhh7g/back-term-monitor/internal/app/api/config"
	"github.com/richhh7g/back-term-monitor/pkg/environment"
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

// @title Term Monitor
// @version 1.0.0
// @description API do [Term Monitor](https://github.com/richhh7g/back-term-monitor) para monitorar concorrentes que usam termos de marca em resultados patrocinados do Google.
// @contact.name Richhh7g
// @contact.url https://github.com/richhh7g
// @contact.email richhh7g@protonmail.com
// @license.name MIT
// @license.url https://github.com/richhh7g/back-term-monitor/blob/main/LICENSE
// @BasePath /api
func main() {
	ctx := context.Background()

	serverConfig := api_config.NewServerConfig(ctx)
	err := serverConfig.Configure()
	if err != nil {
		panic(err)
	}
}
