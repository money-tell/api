package transactions

import (
	"context"

	"github.com/google/uuid"

	"github.com/katalabut/money-tell-api/app/api/models"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

type Manager struct {
	queriesMaster *queries.Queries
	queriesSlave  *queries.Queries
}

func New(master, slave *queries.Queries) *Manager {
	return &Manager{
		queriesMaster: master,
		queriesSlave:  slave,
	}
}

func (m *Manager) GetByUser(ctx context.Context, userID uuid.UUID, params *models.GetTransactionsRequest) ([]*queries.Transaction, error) {
	pays, err := m.queriesSlave.GetTransactionsByUserID(ctx, queries.GetTransactionsByUserIDParams{
		UserID:   userID,
		DateFrom: params.DateFrom,
		DateTo:   params.DateTo,
	})
	if err != nil {
		return nil, err
	}

	return m.fetchRepeated(ctx, userID, params, pays)
}

func (m *Manager) Add(ctx context.Context, userID uuid.UUID, req *models.TransactionRequest) (*queries.Transaction, error) {
	rt := queries.TransactionsRepeatTypeNone
	if req.RepeatType != nil {
		rt = queries.TransactionsRepeatType(*req.RepeatType)
	}

	pay, err := m.queriesMaster.TransactionsInsert(ctx, queries.TransactionsInsertParams{
		UserID:     userID,
		Type:       queries.TransactionsType(req.Type),
		Title:      req.Title,
		Amount:     req.Amount,
		Date:       req.Date,
		RepeatType: rt,
	})
	if err != nil {
		return nil, err
	}

	return pay, nil
}
