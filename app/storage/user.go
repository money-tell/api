package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/katalabut/money-tell-api/app/entities"
)

func (s *StorageImpl) FindUser(ctx context.Context, email, pass string) (*entities.User, error) {
	result := &userRow{}
	filter := bson.D{{"email", email}, {"password", pass}}
	err := s.users.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return mapUserRow(result), nil
}

func (s *StorageImpl) FindUserByID(ctx context.Context, hex string) (*entities.User, error) {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	result := &userRow{}

	err = s.users.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return mapUserRow(result), nil
}

func (s *StorageImpl) InsertUser(ctx context.Context, email, pass string) (*entities.User, error) {
	row := &userRow{
		Email:    email,
		Password: pass,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err := s.users.InsertOne(ctx, row)
	if err != nil {
		return nil, err
	}

	return mapUserRow(row), nil
}
