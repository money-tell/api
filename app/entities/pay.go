package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type PayDate struct {
	Date  string `json:"date"`
	Items []*Pay `json:"items"`
}

const (
	PayTypeAccrual    PayType = "accrual"
	PayTypeRedemption PayType = "redemption"
)

type (
	PayType string

	Pay struct {
		UserID    int64
		Type      PayType
		Title     string
		Amount    decimal.Decimal
		Date      time.Time
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
