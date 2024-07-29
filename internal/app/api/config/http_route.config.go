package api_config

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/richhh7g/term-alarms/internal/app/api/route"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type HTTPRouter struct {
	ctx context.Context
}

func NewHTTPRouterConfig(ctx context.Context) *HTTPRouter {
	return &HTTPRouter{
		ctx: ctx,
	}
}

func (c *HTTPRouter) Configure() (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Use(middleware.DefaultLogger, middleware.RequestID, middleware.RealIP, middleware.Recoverer)

	route.NewApi(router)
	router.Get("/docs/*", httpSwagger.Handler())

	return router, nil
}
