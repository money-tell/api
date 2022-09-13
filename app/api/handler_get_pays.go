package api

import (
	"context"
	"strconv"

	"github.com/katalabut/money-tell-api/app/processors/auth"
)

func (a *Api) GetPays(ctx context.Context, r GetPaysRequestObject) interface{} {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return err
	}

	pays, err := a.processors.Pays.GetPaysByUser(ctx, userID)
	if err != nil {
		return err
	}

	responses := make(GetPays200JSONResponse, 0, len(pays))
	for _, pay := range pays {
		responses = append(responses, Pay{
			Id:    strconv.FormatInt(pay.ID, 10),
			Title: pay.Title,
		})
	}

	return responses
}
