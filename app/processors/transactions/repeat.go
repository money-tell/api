package transactions

import (
	"context"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/ulule/deepcopier"

	"github.com/katalabut/money-tell-api/app/api/models"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

func (m *Manager) fetchRepeated(ctx context.Context, userID uuid.UUID, params *models.GetTransactionsRequest, pays []*queries.Transaction) ([]*queries.Transaction, error) {
	from := params.DateFrom
	to := params.DateTo

	dow := []int32{1, 2, 3, 4, 5, 6, 7}
	if dbd := int(to.Sub(from).Hours()) / 24; dbd < 7 {
		dow = make([]int32, 0, dbd)
		for i := 0; i < dbd; i++ {
			d := from.AddDate(0, 0, i)
			dow = append(dow, int32(d.Weekday()))
		}
	}

	repeated, err := m.queriesSlave.GetRepeatedTransactionsByUserID(ctx, queries.GetRepeatedTransactionsByUserIDParams{
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

	clones := cloneTransactions(repeated, from, to)

	return mergeAndSortPay(pays, clones), nil
}

func mergeAndSortPay(one, two []*queries.Transaction) []*queries.Transaction {
	pays := make([]*queries.Transaction, 0, len(one)+len(two))
	for _, p := range one {
		pays = append(pays, p)
	}
	for _, c := range two {
		pays = append(pays, c)
	}

	sort.Slice(pays, func(i, j int) bool { return pays[i].Date.Before(pays[j].Date) })

	return pays
}

func cloneTransactions(repeated []*queries.Transaction, from time.Time, to time.Time) []*queries.Transaction {
	var clones []*queries.Transaction
	for _, pay := range repeated {
		subDate := pay.Date
		if subDate.Before(from) {
			subDate = from
			pay.Date = pay.Date.AddDate(0, 0, int(subDate.Sub(pay.Date).Hours()/24))
		}

		difference := to.Sub(subDate)
		switch pay.RepeatType {
		case queries.TransactionsRepeatTypeDaily:
			dayBetween := int(difference.Hours() / 24)
			clones = append(clones, clonePayWithStep(*pay, dayBetween, 1, 0, 0)...)
		case queries.TransactionsRepeatTypeWeekly:
			weekBetween := int(difference.Hours() / 24 / 7)
			clones = append(clones, clonePayWithStep(*pay, weekBetween, 7, 0, 0)...)
		case queries.TransactionsRepeatTypeMonthly:
			monthBetween := int(difference.Hours() / 24 / 30)
			clones = append(clones, clonePayWithStep(*pay, monthBetween, 0, 1, 0)...)
		case queries.TransactionsRepeatTypeYearly:
			yearBetween := int(difference.Hours() / 24 / 365)
			clones = append(clones, clonePayWithStep(*pay, yearBetween, 0, 0, 1)...)
		}
	}

	return clones
}

func clonePayWithStep(txn queries.Transaction, count, dayStep, monthStep, yearStep int) []*queries.Transaction {
	txns := make([]*queries.Transaction, 0, count)
	for i := 0; i < count+1; i++ {
		t := &queries.Transaction{}
		err := deepcopier.Copy(txn).To(t)
		if err != nil {
			logrus.Error(err)
			continue
		}

		t.Date = txn.Date.AddDate(i*yearStep, i*monthStep, i*dayStep)
		txns = append(txns, t)
	}

	return txns
}
