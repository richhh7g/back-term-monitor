package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/richhh7g/term-alarms/internal/app/api/response"
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
		r.Get("/example", api.exampleHandler)
	})
}

func (api *ApiImpl) v1Prefix() string {
	return "/api/v1"
}

func (api *ApiImpl) exampleHandler(w http.ResponseWriter, r *http.Request) {
	response.NewSuccess(map[string]string{"hello": "world"}, 200).Send(w)
}
