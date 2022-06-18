package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PendaftaranListErrorResponse struct {
	Error string `json:"error"`
}

type ListPendaftaran struct {
	Id                string `json:"id"`
	IdBeasiswa        string `json:"id_beasiswa"`
	IdSiswa 		  string `json:"id_siswa"`
	TanggalDaftar     string `json:"tanggal_daftar"`
	Status            string `json:"status"`
}

type PendaftaranListResponse struct{
	Pendaftaran []ListPendaftaran `json:"pendaftaran"`
}

func (a *API) getAllPendaftaran (w http.ResponseWriter, r *http.Request) {
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
			Id:                strconv.Itoa(int(p.Id)),
			IdBeasiswa:        strconv.Itoa(int(p.IdBeasiswa)),
			IdSiswa:           strconv.Itoa(int(p.IdSiswa)),
			TanggalDaftar:     p.TanggalDaftar,
			Status:           p.Status,
		})
	}

	encoder.Encode(response)

	w.WriteHeader(http.StatusOK)
}

func (a *API) getPendaftaranById (w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	pendaftaran, err := a.pendaftaranRepo.GetPendaftaranById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
		Id:                strconv.Itoa(int(pendaftaran.Id)),
		IdBeasiswa:        strconv.Itoa(int(pendaftaran.IdBeasiswa)),
		IdSiswa:           strconv.Itoa(int(pendaftaran.IdSiswa)),
		TanggalDaftar:     pendaftaran.TanggalDaftar,
		Status:           pendaftaran.Status,
	})

	encoder.Encode(response)

	w.WriteHeader(http.StatusOK)
}

func (a *API) getPendaftaranBySiswa (w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
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
			Id:                strconv.Itoa(int(p.Id)),
			IdBeasiswa:        strconv.Itoa(int(p.IdBeasiswa)),
			IdSiswa:           strconv.Itoa(int(p.IdSiswa)),
			TanggalDaftar:     p.TanggalDaftar,
			Status:           p.Status,
		})
	}

	encoder.Encode(response)

	w.WriteHeader(http.StatusOK)
}

func (a *API) getPendaftaranByBeasiswa (w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}
	response.Pendaftaran = make([]ListPendaftaran, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
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
			Id:                strconv.Itoa(int(p.Id)),
			IdBeasiswa:        strconv.Itoa(int(p.IdBeasiswa)),
			IdSiswa:           strconv.Itoa(int(p.IdSiswa)),
			TanggalDaftar:     p.TanggalDaftar,
			Status:           p.Status,
		})
	}

	encoder.Encode(response)

	w.WriteHeader(http.StatusOK)
}

func ( a *API) createPendaftaran (w http.ResponseWriter, r *http.Request) {
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

	err = a.pendaftaranRepo.CreatePendaftaran(pendaftaran.IdBeasiswa, pendaftaran.IdSiswa, pendaftaran.TanggalDaftar, pendaftaran.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	response.Pendaftaran = append(response.Pendaftaran, ListPendaftaran{
		Id:                strconv.Itoa(int(pendaftaran.Id)),
		IdBeasiswa:        strconv.Itoa(int(pendaftaran.IdBeasiswa)),
		IdSiswa:           strconv.Itoa(int(pendaftaran.IdSiswa)),
		TanggalDaftar:     pendaftaran.TanggalDaftar,
		Status:            pendaftaran.Status,
	})

	encoder.Encode(response)

	w.WriteHeader(http.StatusOK)
}
