package api_config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/richhh7g/back-term-monitor/pkg/environment"
)

type ServerConfig struct {
	ctx context.Context
}

func NewServerConfig(ctx context.Context) *ServerConfig {
	return &ServerConfig{
		ctx: ctx,
	}
}

func (c *ServerConfig) Configure() error {
	httpRouterConfig := NewHTTPRouterConfig(c.ctx)
	router, err := httpRouterConfig.Configure()
	if err != nil {
		return err
	}

	port := environment.Get[string]("APP_PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		return err
	}

	return err
}
