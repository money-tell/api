package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/katalabut/money-tell-api/internal/storage"
)

func (a *Api) Configure() error {
	a.storage = storage.New()

	return nil
}

func (a *Api) SetupMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
}

func (a *Api) SetupHandlers(r *chi.Mux) {
	r.Get("/pays", a.getPays)
	r.Post("/pays", a.addPay)
}
