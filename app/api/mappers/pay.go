package mappers

import (
	"github.com/katalabut/money-tell-api/app/api/models"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

func MapPay(p queries.Pay) models.Pay {
	rt := models.RepeatType(p.RepeatType)

	return models.Pay{
		Id:        p.ID.String(),
		UpdatedAt: p.UpdatedAt.Time,
		CreatedAt: p.CreatedAt.Time,

		PayRequest: models.PayRequest{
			Title:      p.Title,
			Amount:     p.Amount,
			Date:       p.Date,
			RepeatType: &rt,
			Type:       models.PayType(p.Type),
		},
	}
}
