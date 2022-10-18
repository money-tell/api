package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katalabut/money-tell-api/app/api/mappers"
	"github.com/katalabut/money-tell-api/app/api/models"
	"github.com/katalabut/money-tell-api/app/processors/auth"
)

//func (a *Api) GetPays(ctx context.Context, r genApi.GetPaysRequestObject) interface{} {
/*userID, err := auth.UserIDFromCtx(ctx)
if err != nil {
	return err
}

pays, err := a.processors.Pays.GetPaysByUser(ctx, userID, r.Params)
if err != nil {
	return err
}

responses := make(genApi.GetPays200JSONResponse, 0, len(pays))
for _, pay := range pays {
	responses = append(responses, genApi.Pay{
		Id:    strconv.FormatInt(pay.ID, 10),
		Title: pay.Title,
	})
}*/

//	return nil
//}

func (a *Api) GetPays(c echo.Context) error {
	userID, err := auth.UserIDFromEchoCtx(c)
	if err != nil {
		return err
	}

	r := new(models.GetPaysRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	pays, err := a.processors.Pays.GetPaysByUser(c.Request().Context(), userID, r)
	if err != nil {
		return err
	}

	responses := make([]models.Pay, 0, len(pays))
	for _, pay := range pays {
		responses = append(responses, mappers.MapPay(*pay))
	}

	return c.JSON(http.StatusOK, responses)
}
