package api

import (
	"encoding/json"
	"net/http"
)

func (a *API) GetAllSiswa(w http.ResponseWriter, r *http.Request) {
	result, err := a.siswaRepo.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error : Internal server error"))
		return
	}

	response, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error : Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
