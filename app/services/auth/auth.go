package auth

import (
	"context"

	"github.com/go-chi/jwtauth/v5"

	"github.com/katalabut/money-tell-api/app/entities"
	"github.com/katalabut/money-tell-api/app/storage"
)

type Auth struct {
	tokenAuth *jwtauth.JWTAuth
	storage   storage.Storage
}

func New(cfg Config, storage storage.Storage) *Auth {
	return &Auth{
		tokenAuth: jwtauth.New("HS256", []byte(cfg.Secret), nil),
		storage:   storage,
	}
}

func (a *Auth) GetTokenAuth() *jwtauth.JWTAuth {
	return a.tokenAuth
}

func (a *Auth) BaseLogin(ctx context.Context, email string, password string) (*entities.User, error) {
	if user, err := a.storage.FindUser(ctx, email, password); err == nil {
		return user, nil
	}

	return nil, entities.ErrUserNotFound
}

func (a *Auth) MakeToken(id string) (string, error) {
	claims := make(map[string]interface{})
	claims["id"] = id
	jwtauth.SetIssuedNow(claims)

	_, tokenString, err := a.tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
