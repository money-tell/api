package api

import (
	"context"
)

func (a *Api) AuthEmail(ctx context.Context, r AuthEmailRequestObject) interface{} {
	user, err := a.processors.Auth.BaseLogin(ctx, r.Body.Email, r.Body.Password)
	if err != nil {
		return AuthEmail401Response{}
	}

	token, err := a.processors.Auth.MakeToken(user.ID)
	if err != nil {
		return AuthEmail401Response{}
	}

	return AuthEmail200JSONResponse{Token: token}
}
