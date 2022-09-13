package api

import (
	"context"

	"github.com/katalabut/money-tell-api/app/processors/auth"
)

func (a *Api) AddPay(ctx context.Context, request AddPayRequestObject) interface{} {
	userID, err := auth.UserIDFromCtx(ctx)
	if err != nil {
		return err
	}

	pay, err := a.processors.Pays.AddPay(ctx, userID, request)
	if err != nil {
		return err
	}

	return
	//user, err := a.processors.Auth.BaseLogin(ctx, r.Body.Email, r.Body.Password)
	//if err != nil {
	//	return AuthEmail401Response{}
	//}
	//
	//token, err := a.processors.Auth.MakeToken(user.ID)
	//if err != nil {
	//	return AuthEmail401Response{}
	//}
	//
	//return AuthEmail200JSONResponse{Token: token}
}
