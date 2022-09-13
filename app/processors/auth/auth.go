package auth

import (
	"context"
	"errors"

	"github.com/lestrrat-go/jwx/jwt"

	"github.com/katalabut/money-tell-api/app/entities"
	queries "github.com/katalabut/money-tell-api/app/generated/db"
)

var tokenNotValidError = errors.New("token not valid")

type Auth struct {
	tokenAuth *JWTAuth
	queries   *queries.Queries
}

func New(cfg Config, queries *queries.Queries) *Auth {
	return &Auth{
		tokenAuth: NewJwt("HS256", []byte(cfg.Secret), nil),
		queries:   queries,
	}
}

func (a *Auth) GetTokenAuth() *JWTAuth {
	return a.tokenAuth
}

func (a *Auth) BaseLogin(ctx context.Context, email string, password string) (*queries.User, error) {
	if user, err := a.queries.FindUser(ctx, queries.FindUserParams{
		Email:    email,
		Password: password,
	}); err == nil {
		return user, nil
	}

	return nil, entities.ErrUserNotFound
}

func (a *Auth) MakeToken(id int64) (string, error) {
	claims := make(map[string]interface{})
	claims["id"] = id
	SetIssuedNow(claims)

	_, tokenString, err := a.tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func UserIDFromCtx(ctx context.Context) (int64, error) {
	token, _, err := FromContext(ctx)
	if err != nil {
		return 0, err
	}

	return UserIDFromToken(token)
}

func UserIDFromToken(token jwt.Token) (int64, error) {
	if token == nil || jwt.Validate(token) != nil {
		return 0, tokenNotValidError
	}

	cUserId, ok := token.Get("id")
	if !ok {
		return 0, tokenNotValidError
	}

	userId, ok := cUserId.(float64)
	if !ok {
		return 0, tokenNotValidError
	}

	return int64(userId), nil
}
