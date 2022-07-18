package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type payDateRow struct {
	Date  string    `bson:"_id,omitempty"`
	Items []*payRow `bson:"items,omitempty"`
}

type payRow struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Price       float32            `bson:"price,omitempty"`
	PaymentDate time.Time          `bson:"payment_date,omitempty"`
}

type userRow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	CreateAt time.Time          `bson:"create_at"`
	UpdateAt time.Time          `bson:"update_at"`
}
