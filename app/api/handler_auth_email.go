package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katalabut/money-tell-api/app/api/models"
)

func (a *Api) AuthEmail(c echo.Context) error {
	r := new(models.AuthEmailRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	token, err := a.processors.Auth.GenTokenByBasicLogin(c.Request().Context(), r.Email, r.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, models.TokenResponse{Token: token})
}
