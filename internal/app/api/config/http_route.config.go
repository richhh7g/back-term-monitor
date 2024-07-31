package api_config

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router.Use(corsMiddleware, middleware.DefaultLogger, middleware.RequestID, middleware.RealIP, middleware.Recoverer)

	route.NewApi(router)
	router.Get("/docs/*", httpSwagger.Handler())

	return router, nil
}
