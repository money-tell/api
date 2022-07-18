package api

import (
	"net/http"
)

type addPayRequest struct {
	request
	Title string  `json:"title" validate:"required"`
	Price float32 `json:"price" validate:"required"`
}

func (a *Api) addPay(w http.ResponseWriter, r *http.Request) {
	payReq := &addPayRequest{}

	if err := bindRequest(r, payReq); err != nil {
		response(w, r, http.StatusBadRequest, "Empty body")
		return
	}

	pay, err := a.storage.AddPay(payReq.Title, payReq.Price)
	if err != nil {
		responseError(w, r, http.StatusInternalServerError, err, "insert pay error")
	}

	response(w, r, http.StatusOK, pay)
}
