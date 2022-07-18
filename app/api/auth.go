package api

import (
	"net/http"
)

// https://coursehunters.online/t/reliable-webservers-with-go-part-3/4877

type loginRequest struct {
	request
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (a *Api) Login(w http.ResponseWriter, r *http.Request) {
	req := &loginRequest{}

	if err := bindRequest(r, req); err != nil {
		response(w, r, http.StatusBadRequest, "Empty body")
		return
	}

	user, err := a.auth.BaseLogin(r.Context(), req.Email, req.Password)
	if err != nil {
		responseError(w, r, http.StatusBadRequest, err, "login unsuccessful")
		return
	}

	token, err := a.auth.MakeToken(user.ID)
	if err != nil {
		responseError(w, r, http.StatusBadRequest, err, "login unsuccessful")
		return
	}

	tokenResp := &tokenResponse{Token: token}
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Value: token,
	})

	response(w, r, http.StatusOK, tokenResp)
}
