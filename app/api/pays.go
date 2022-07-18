package api

import (
	"net/http"
)

func (a *Api) getPays(w http.ResponseWriter, r *http.Request) {
	pays, err := a.storage.GetPays(r.Context())
	if err != nil {
		responseError(w, r, http.StatusInternalServerError, err, "bad request")
		return
	}

	response(w, r, http.StatusOK, pays)
}
