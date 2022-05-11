package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/katalabut/money-tell-api/internal/config"
	"github.com/katalabut/money-tell-api/internal/storage"
)

type Api struct {
	storage storage.Storage
}

func ServeApi(cfg *config.Config) error {
	api := &Api{}

	err := api.Configure()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	api.SetupMiddlewares(r)
	api.SetupHandlers(r)

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), r)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
