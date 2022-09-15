package pays

import (
	"context"

	"github.com/katalabut/money-tell-api/app/api"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

func (m *Manager) prepareRepeatedPays(ctx context.Context, userID int64, params api.GetPaysParams) ([]*queries.Pay, error) {
	from := params.DateFrom.Time
	to := params.DateTo.Time

	dow := []int32{1, 2, 3, 4, 5, 6, 7}
	if dc := int(to.Sub(from).Hours()) / 24; dc < 7 {
		dow = make([]int32, 0, dc)
		for i := 0; i < dc; i++ {
			d := from.AddDate(0, 0, i)
			dow = append(dow, int32(d.Weekday()))
		}
	}

	_, err := m.queriesSlave.GetRepeatedPaysByUserID(ctx, queries.GetRepeatedPaysByUserIDParams{
		UserID:         userID,
		DaysOfWeek:     dow,
		MonthlyDayFrom: int32(from.Day()),
		MonthlyDayTo:   int32(to.Day()),
		YearlyDayFrom:  int32(from.YearDay()),
		YearlyDayTo:    int32(to.YearDay()),
	})
	if err != nil {
		return nil, err
	}

	return nil, err
}
