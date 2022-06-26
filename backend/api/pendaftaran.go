package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"final-project-eng2-be/repository"
)

type PendaftaranListErrorResponse struct {
	Error string `json:"error"`
}
type PendaftaranSuccessfulResponse struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}

type PendaftaranListResponse struct {
	Pendaftaran []repository.PendaftaranResponse `json:"pendaftaran"`
}

func (a *API) getAllPendaftaran(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}

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

	response.Pendaftaran = pendaftaran
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}

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

	response.Pendaftaran = append(response.Pendaftaran, pendaftaran)
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranBySiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}

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

	response.Pendaftaran = pendaftaran
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) getPendaftaranByBeasiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := PendaftaranListResponse{}

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

	response.Pendaftaran = pendaftaran
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) createPendaftaran(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	var pendaftaran repository.Pendaftaran
	err := json.NewDecoder(r.Body).Decode(&pendaftaran)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	if pendaftaran.Status == "" {
		pendaftaran.Status = "Menunggu Pengumuman"
	}

	SiswaData := r.Context().Value("siswa_data").(Siswa)
	idNum, _ := strconv.Atoi(SiswaData.Id)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = a.pendaftaranRepo.CreatePendaftaran(int(pendaftaran.IdBeasiswa), idNum, pendaftaran.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(PendaftaranSuccessfulResponse{
		Msg:    "Successful",
		Status: pendaftaran.Status,
	})
}

func (a *API) updatePendaftaran(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	var pendaftaran repository.Pendaftaran
	err := json.NewDecoder(r.Body).Decode(&pendaftaran)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	err = a.pendaftaranRepo.UpdatePendaftaran(int(pendaftaran.Id), pendaftaran.Status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranListErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	encoder.Encode(PendaftaranSuccessfulResponse{Msg: "Successful", Status: pendaftaran.Status})
}
