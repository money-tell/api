package mappers

import (
	"github.com/katalabut/money-tell-api/app/api"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

func MapAddPayResponse(pay *queries.Pay) *api.AddPay200JSONResponse {
	return &api.AddPay200JSONResponse{
		Token: "",
	}
}
