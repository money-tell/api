package entities

import "time"

type PayDate struct {
	Date  string `json:"date"`
	Items []*Pay `json:"items"`
}

type Pay struct {
	Title       string    `json:"title"`
	Price       float32   `json:"price"`
	PaymentDate time.Time `json:"payment_date"`
}
