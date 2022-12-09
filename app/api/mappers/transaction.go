package mappers

import (
	"github.com/katalabut/money-tell-api/app/api/models"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

func MapTxn(t queries.Transaction) models.Transaction {
	rt := models.RepeatType(t.RepeatType)

	return models.Transaction{
		Id:        t.ID.String(),
		UpdatedAt: t.UpdatedAt.Time,
		CreatedAt: t.CreatedAt.Time,

		TransactionRequest: models.TransactionRequest{
			Title:      t.Title,
			Amount:     t.Amount,
			Date:       t.Date,
			RepeatType: &rt,
			Type:       models.TransactionType(t.Type),
		},
	}
}
