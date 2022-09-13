//go:generate oapi-codegen --config ../../oapi/cfg.yaml ../../oapi/oapi.yaml

package api

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	"github.com/katalabut/money-tell-api/app/config"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
	"github.com/katalabut/money-tell-api/app/processors"
	"github.com/katalabut/money-tell-api/app/processors/auth"
	"github.com/katalabut/money-tell-api/app/processors/pays"
	"github.com/katalabut/money-tell-api/app/system/postgres"
)

type Api struct {
	processors *processors.Container
}

func Run(cfg *config.Config) error {
	logrus.Info("Api initialising")

	api, err := configure(cfg)
	if err != nil {
		return err
	}

	spec, err := GetSwagger()
	if err != nil {
		return err
	}

	validator := oapiMiddleware.OapiRequestValidatorWithOptions(spec,
		&oapiMiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: auth.NewAuthenticator(api.processors.Auth.GetTokenAuth()),
			},
		})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(validator)

	s := NewStrictHandler(api, nil)
	RegisterHandlersWithBaseURL(e, s, "/v1")

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.HttpPort)))

	return nil
}

func configure(cfg *config.Config) (*Api, error) {
	api := &Api{}

	dbConn, err := postgres.New(cfg.Postgres)
	if err != nil {
		return nil, err
	}

	queriesMaster := queries.New(dbConn.Master())
	queriesSlave := queries.New(dbConn.Slave())

	api.processors = processors.New(
		pays.New(queriesMaster, queriesSlave),
		auth.New(cfg.Auth, queriesSlave),
	)

	return api, nil
}
