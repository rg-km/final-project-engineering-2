package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *API) GetAllSiswa(w http.ResponseWriter, r *http.Request) {

	a.AllowOrigin(w, r)
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

func (a *API) GetSiswaByID(w http.ResponseWriter, r *http.Request) {

	a.AllowOrigin(w, r)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("error : Internal server error"))
		return
	}

	res, err := a.siswaRepo.GetSiswaByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("error : No siswa with id = %d", id)))
		return
	}

	data, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("error : Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}
