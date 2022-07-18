package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type signupRequest struct {
	request
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (a *Api) signup(w http.ResponseWriter, r *http.Request) {
	data := &signupRequest{}
	if err := bindRequest(r, data); err != nil {
		responseError(w, r, http.StatusBadRequest, err, "Login or password do not matched")
		return
	}

	render.Status(r, http.StatusCreated)
}
