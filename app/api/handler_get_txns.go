package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katalabut/money-tell-api/app/api/mappers"
	"github.com/katalabut/money-tell-api/app/api/models"
	"github.com/katalabut/money-tell-api/app/processors/auth"
)

func (a *Api) GetTransactions(c echo.Context) error {
	userID, err := auth.UserIDFromEchoCtx(c)
	if err != nil {
		return err
	}

	r := new(models.GetTransactionsRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	pays, err := a.processors.Txn.GetByUser(c.Request().Context(), userID, r)
	if err != nil {
		return err
	}

	responses := make([]models.Transaction, 0, len(pays))
	for _, pay := range pays {
		responses = append(responses, mappers.MapTxn(*pay))
	}

	return c.JSON(http.StatusOK, responses)
}
