package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/katalabut/money-tell-api/app/entities"
)

func (s *StorageImpl) GetPays(ctx context.Context) ([]*entities.PayDate, error) {
	pipline := bson.A{
		bson.D{
			{"$group",
				bson.D{
					{"_id",
						bson.D{
							{"$dateToString",
								bson.D{
									{"format", "%Y-%m-%d"},
									{"date", "$payment_date"},
								},
							},
						},
					},
					{"items", bson.D{{"$push", "$$ROOT"}}},
				},
			},
		},
	}

	cursor, err := s.pays.Aggregate(ctx, pipline)
	if err != nil {
		return nil, err
	}

	var results []*payDateRow
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return mapPayDateRow(results), nil
}

func (s *StorageImpl) AddPay(title string, price float32) (*entities.Pay, error) {
	pay := &payRow{
		Title:       title,
		Price:       price,
		PaymentDate: time.Now(),
	}

	_, err := s.pays.InsertOne(context.TODO(), pay)

	if err != nil {
		return nil, err
	}

	return mapPay(pay), nil
}
