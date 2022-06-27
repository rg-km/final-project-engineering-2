package api

import (
	"encoding/json"
	"final-project-eng2-be/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type SiswaErrorResponse struct {
	Error string `json:"error"`
}
type SiswaSuccessfulResponse struct {
	Siswa []ListSiswa `json:"siswa"`
}

type ListSiswa struct {
	Id                string `json:"id"`
	Nama              string `json:"nama"`
	Email             string `json:"email"`
	JenjangPendidikan string `json:"jenjang_pendidikan"`
	Nik               string `json:"nik"`
	TanggalLahir      string `json:"tanggal_lahir"`
	TempatLahir       string `json:"tempat_lahir"`
	KotaDomisili      string `json:"kota_domisili"`
}

func (a *API) UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	var siswaInput Siswa

	err := json.NewDecoder(r.Body).Decode(&siswaInput)
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	SiswaData := r.Context().Value("siswa_data").(Siswa)
	idNum, _ := strconv.Atoi(SiswaData.Id)
	siswaCurrent, err := a.siswaRepo.GetSiswaByID(idNum)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	siswaNew := repository.Siswa{
		Id:                int64(idNum),
		Nama:              siswaInput.Nama,
		Email:             siswaInput.Email,
		TanggalLahir:      siswaInput.TanggalLahir,
		TempatLahir:       siswaInput.TempatLahir,
		JenjangPendidikan: siswaInput.JenjangPendidikan,
		Nik:               siswaInput.Nik,
		KotaDomisili:      siswaInput.KotaDomisili,
	}

	if siswaInput.Nama == "" {
		siswaNew.Nama = siswaCurrent.Nama
	}
	if siswaInput.Email == "" {
		siswaNew.Email = siswaCurrent.Email
	}
	if siswaInput.TanggalLahir == "" {
		siswaNew.TanggalLahir = siswaCurrent.TanggalLahir
	}
	if siswaInput.TempatLahir == "" {
		siswaNew.TempatLahir = siswaCurrent.TempatLahir
	}
	if siswaInput.JenjangPendidikan == "" {
		siswaNew.JenjangPendidikan = siswaCurrent.JenjangPendidikan
	}
	if siswaInput.KotaDomisili == "" {
		siswaNew.KotaDomisili = siswaCurrent.KotaDomisili
	}
	if siswaInput.Nik == "" {
		siswaNew.Nik = siswaCurrent.Nik
	}
	err = a.siswaRepo.UpdateSiswa(siswaNew)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(siswaNew)
}

func (a *API) GetSiswaByToken(w http.ResponseWriter, r *http.Request) {
	a.AllowOrigin(w, r)
	w.Header().Set("Content-Type", "application/json")
	SiswaData := r.Context().Value("siswa_data").(Siswa)
	encoder := json.NewEncoder(w)

	expTime := time.Now().Add(60 * time.Minute)
	newTokenString, err := a.GenerateSiswaToken(SiswaData, expTime)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(map[string]interface{}{
		"siswa": SiswaData,
		"token": newTokenString,
	})
	return
}
func (a *API) GetAllSiswa(w http.ResponseWriter, r *http.Request) {

	a.AllowOrigin(w, r)
	w.Header().Set("Content-Type", "application/json")
	response := SiswaSuccessfulResponse{}
	response.Siswa = make([]ListSiswa, 0)
	siswa, err := a.siswaRepo.GetAll()
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(SiswaErrorResponse{Error: "Internal server error"})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(SiswaErrorResponse{Error: "Internal server error"})
		return
	}

	for _, s := range siswa {
		response.Siswa = append(response.Siswa, ListSiswa{
			Id:                strconv.Itoa(int(s.Id)),
			Nama:              s.Nama,
			Nik:               s.Nik,
			Email:             s.Email,
			TanggalLahir:      s.TanggalLahir,
			TempatLahir:       s.TempatLahir,
			JenjangPendidikan: s.JenjangPendidikan,
			KotaDomisili:      s.KotaDomisili,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
}

func (a *API) GetSiswaByID(w http.ResponseWriter, r *http.Request) {

	a.AllowOrigin(w, r)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	encoder := json.NewEncoder(w)
	response := SiswaSuccessfulResponse{}
	response.Siswa = make([]ListSiswa, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		encoder.Encode(SiswaErrorResponse{Error: "Internal server error"})
		return
	}

	res, err := a.siswaRepo.GetSiswaByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		encoder.Encode(SiswaErrorResponse{Error: fmt.Sprintf("No siswa with id = %d", id)})
		return
	}

	response.Siswa = append(response.Siswa, ListSiswa{
		Id:                strconv.Itoa(int(res.Id)),
		Nama:              res.Nama,
		Nik:               res.Nik,
		Email:             res.Email,
		TanggalLahir:      res.TanggalLahir,
		TempatLahir:       res.TempatLahir,
		JenjangPendidikan: res.JenjangPendidikan,
		KotaDomisili:      res.KotaDomisili,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(response)
	return
}
