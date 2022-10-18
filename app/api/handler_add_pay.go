package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katalabut/money-tell-api/app/api/mappers"
	"github.com/katalabut/money-tell-api/app/api/models"
	"github.com/katalabut/money-tell-api/app/processors/auth"
)

func (a *Api) AddPay(c echo.Context) error {
	userID, err := auth.UserIDFromEchoCtx(c)
	if err != nil {
		return err
	}

	r := new(models.PayRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	pay, err := a.processors.Pays.AddPay(c.Request().Context(), userID, r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, mappers.MapPay(*pay))
}
