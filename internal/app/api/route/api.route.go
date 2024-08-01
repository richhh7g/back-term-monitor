package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/richhh7g/term-alarms/internal/app/api/handler"
)

type ApiImpl struct {
	router *chi.Mux
}

func NewApi(router *chi.Mux) {
	api := &ApiImpl{
		router: router,
	}

	api.RegisterRoutes()
}

func (api *ApiImpl) RegisterRoutes() {
	api.router.Route(api.v1Prefix(), func(r chi.Router) {
		r.Post("/check-term", handler.NewCheckTermHandler)
	})
}

func (api *ApiImpl) v1Prefix() string {
	return "/api/v1"
}
