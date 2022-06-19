package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PendaftaranListErrorResponse struct {
	Error string `json:"error"`
}
type PendaftaranSuccessfulResponse struct {
	Msg string `json:"msg"`
}

type ListPendaftaran struct {
	Id            int    `json:"id"`
	IdBeasiswa    int    `json:"id_beasiswa"`
	IdSiswa       int    `json:"id_siswa"`
	TanggalDaftar string `json:"tanggal_daftar"`
	Status        string `json:"status"`
}

type PendaftaranListResponse struct {
	Pendaftaran []ListPendaftaran `json:"pendaftaran"`
}

func (a *API) getAllPendaftaran(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	pendaftaran, err := a.pendaftaranRepo.GetPendaftaranAll()

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}

	for _, p := range pendaftaran {
		response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
			Id:            int(p.Id),
			IdBeasiswa:    int(p.IdBeasiswa),
			IdSiswa:       int(p.IdSiswa),
			TanggalDaftar: p.TanggalDaftar,
			Status:        p.Status,
		})
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: "id query required"})
		return
	}

	pendaftaran, err := a.pendaftaranRepo.GetPendaftaranById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
		Id:            int(pendaftaran.Id),
		IdBeasiswa:    int(pendaftaran.IdBeasiswa),
		IdSiswa:       int(pendaftaran.IdSiswa),
		TanggalDaftar: pendaftaran.TanggalDaftar,
		Status:        pendaftaran.Status,
	})

	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranBySiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: "id query is required"})
		return
	}

	pendaftaran, err := a.pendaftaranRepo.GetBySiswa(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	for _, p := range pendaftaran {
		response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
			Id:            int(p.Id),
			IdBeasiswa:    int(p.IdBeasiswa),
			IdSiswa:       int(p.IdSiswa),
			TanggalDaftar: p.TanggalDaftar,
			Status:        p.Status,
		})
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranByBeasiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: "id query is required"})
		return
	}

	pendaftaran, err := a.pendaftaranRepo.GetByBeasiswa(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	for _, p := range pendaftaran {
		response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
			Id:            int(p.Id),
			IdBeasiswa:    int(p.IdBeasiswa),
			IdSiswa:       int(p.IdSiswa),
			TanggalDaftar: p.TanggalDaftar,
			Status:        p.Status,
		})
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)

}

func (a *API) createPendaftaran(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	var pendaftaran ListPendaftaran
	err := json.NewDecoder(r.Body).Decode(&pendaftaran)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	err = a.pendaftaranRepo.CreatePendaftaran(pendaftaran.IdBeasiswa, pendaftaran.IdSiswa, pendaftaran.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
		Id:            pendaftaran.Id,
		IdBeasiswa:    pendaftaran.IdBeasiswa,
		IdSiswa:       pendaftaran.IdSiswa,
		TanggalDaftar: pendaftaran.TanggalDaftar,
		Status:        pendaftaran.Status,
	})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(PendaftaranSuccessfulResponse{Msg: "Successful"})

}
