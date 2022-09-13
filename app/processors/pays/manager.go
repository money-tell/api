package pays

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/katalabut/money-tell-api/app/api"
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

func (m *Manager) GetPaysByUser(ctx context.Context, userID int64) ([]*queries.Pay, error) {
	return m.queriesSlave.GetPaysByUserID(ctx, userID)
}

func (m *Manager) AddPay(ctx context.Context, userID int64, req api.AddPayRequestObject) (*queries.Pay, error) {
	repeatType := queries.NullPaysRepeatType{}
	if req.Body.RepeatType != nil {
		err := repeatType.Scan(*req.Body.RepeatType)
		if err != nil {
			logrus.Errorf("AddPay: error parse decoding spec: %s", err)
		}
	}

	pay, err := m.queriesMaster.PayInsert(ctx, queries.PayInsertParams{
		UserID:     userID,
		Type:       queries.PaysType(req.Body.Type),
		Title:      req.Body.Title,
		Amount:     req.Body.Amount,
		Date:       req.Body.Date,
		RepeatType: repeatType,
	})
	if err != nil {
		return nil, err
	}

	return pay, nil
}
