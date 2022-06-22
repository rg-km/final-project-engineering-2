package api

import (
	"encoding/json"
	"final-project-eng2-be/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	Nik 			  string `json:"nik"`
	TanggalLahir      string `json:"tanggal_lahir"`
	TempatLahir       string `json:"tempat_lahir"`
	KotaDomisili 	  string `json:"kota_domisili"`
}

func (a *API) getSiswaFromToken(tokenStr string) (*Siswa, error) {

	SiswaClaims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, SiswaClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is Invalid")
	}

	retrivedSiswa := SiswaClaims.SiswaData
	return &retrivedSiswa, nil
}
func (a *API) UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	var siswaInput Siswa

	err := json.NewDecoder(r.Body).Decode(&siswaInput)
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	tokenStr := c.Value
	siswaFromToken, err := a.getSiswaFromToken(tokenStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idNum, _ := strconv.Atoi(siswaFromToken.Id)
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
		Nik: 			   siswaInput.Nik,
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
	c, err := r.Cookie("token")
	encoder := json.NewEncoder(w)

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	tokenStr := c.Value
	retrivedSiswa, err := a.getSiswaFromToken(tokenStr)
	if err != nil {
		if err == jwt.ErrSignatureInvalid || err.Error() == "Token is Invalid" {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	expTime := time.Now().Add(60 * time.Minute)
	newTokenString, err := a.GenerateSiswaToken(*retrivedSiswa, expTime)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   newTokenString,
		Expires: expTime,
		Path:    "/api",
	})

	w.WriteHeader(http.StatusOK)
	encoder.Encode(retrivedSiswa)
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
			Nik: 			   s.Nik,
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
		Nik: 			   res.Nik,
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
