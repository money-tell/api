package models

import (
	"time"
)

type (
	GetTransactionsRequest struct {
		DateFrom time.Time `query:"date_from" validate:"required"`
		DateTo   time.Time `query:"date_to" validate:"required"`
	}

	Transaction struct {
		Id        string    `json:"id"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`

		TransactionRequest
	}

	TransactionRequest struct {
		Title  string    `json:"title" validate:"required"`
		Amount string    `json:"amount" validate:"required,numeric"`
		Date   time.Time `json:"date" validate:"required"`

		// RepeatType Тип повторения: * daily - Каждый день * weekly - Каждую неделю * monthly - Каждый месяц
		RepeatType *RepeatType `json:"repeat_type,omitempty"`
		// Type Тип оплаты: * accrual - начисление * redemption - списание
		Type TransactionType `json:"type"`
	}

	// TransactionType Тип оплаты: * accrual - начисление * redemption - списание
	TransactionType string

	// RepeatType Тип повторения: * daily - Каждый день * weekly - Каждую неделю * monthly - Каждый месяц
	RepeatType string
)
