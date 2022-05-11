package api

import (
	"net/http"
)

func (a *Api) getPays(w http.ResponseWriter, r *http.Request) {
	pays, err := a.storage.GetPays()
	if err != nil {
		respondWithJSON(w, 500, "err to get")
		return
	}

	respondWithJSON(w, 200, pays)
}
