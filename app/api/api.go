package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/katalabut/money-tell-api/app/config"
	"github.com/katalabut/money-tell-api/app/services/auth"
	"github.com/katalabut/money-tell-api/app/storage"
)

type Api struct {
	cfg *config.Config

	auth     *auth.Auth
	storage  storage.Storage
	dbClient *mongo.Client
}

func Run(cfg *config.Config) error {
	log.Println("Api initialising")

	api := &Api{
		cfg: cfg,
	}

	err := api.setup()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	api.routes(r)

	log.Println("Api starting ListenAndServe")

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.HttpPort), r)
}

func (a *Api) setup() error {
	var err error
	a.dbClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(a.cfg.Mongo.Uri))
	if err != nil {
		return err
	}

	if err := a.dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	a.storage = storage.New(a.dbClient.Database(a.cfg.Mongo.DataBase))
	a.auth = auth.New(a.cfg.Auth, a.storage)

	return nil
}

func (a *Api) routes(r *chi.Mux) {
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/signup", a.signup)
		r.Post("/login", a.Login)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(a.auth.GetTokenAuth()))

			r.Use(a.auth.Authenticator)

			r.Route("/pays", func(r chi.Router) {
				r.Get("/", a.getPays)
				r.Post("/", a.addPay)
			})
		})
	})
}
