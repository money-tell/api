package storage

import "github.com/katalabut/money-tell-api/internal/entities"

var pays = make([]entities.Pay, 0)

func (s *StorageImpl) GetPays() ([]entities.Pay, error) {
	return pays, nil
}

func (s *StorageImpl) AddPay(title string, price float32) entities.Pay {
	p := entities.Pay{
		Title: title,
		Price: price,
	}

	pays = append(pays, p)

	return p
}
