package pays

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

func (m *Manager) GetPaysByUser(ctx context.Context, userID uuid.UUID, params *models.GetPaysRequest) ([]*queries.Pay, error) {
	pays, err := m.queriesSlave.GetPaysByUserID(ctx, queries.GetPaysByUserIDParams{
		UserID:   userID,
		DateFrom: params.DateFrom,
		DateTo:   params.DateTo,
	})
	if err != nil {
		return nil, err
	}

	return m.fetchRepeatedPays(ctx, userID, params, pays)
}

func (m *Manager) AddPay(ctx context.Context, userID uuid.UUID, req *models.PayRequest) (*queries.Pay, error) {
	rt := queries.PaysRepeatTypeNone
	if req.RepeatType != nil {
		rt = queries.PaysRepeatType(*req.RepeatType)
	}

	pay, err := m.queriesMaster.PayInsert(ctx, queries.PayInsertParams{
		UserID:     userID,
		Type:       queries.PaysType(req.Type),
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
