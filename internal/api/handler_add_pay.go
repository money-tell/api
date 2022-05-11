package api

import (
	"encoding/json"
	"net/http"
)

type addPayRequest struct {
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

func (a *Api) addPay(w http.ResponseWriter, r *http.Request) {
	payReq := &addPayRequest{}

	err := json.NewDecoder(r.Body).Decode(&payReq)
	if err != nil {
		respondWithJSON(w, 400, "Empty body")
		return
	}

	pay := a.storage.AddPay(payReq.Title, payReq.Price)

	respondWithJSON(w, 200, pay)
}
