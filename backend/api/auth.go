package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Siswa struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	Nama              string `json:"nama"`
	JenjangPendidikan string `json:"jenjang_pendidikan"`
	Nik               string `json:"nik"`
	TempatLahir       string `json:"tempat_lahir"`
	TanggalLahir      string `json:"tanggal_lahir"`
}

type LoginSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterSuccessResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	//Email string `json:"email"`
	SiswaData Siswa `json:"siswa_data"`
	jwt.StandardClaims
}

func (api *API) GenerateSiswaToken(siswa Siswa, expTime time.Time) (string, error) {
	claims := &Claims{
		SiswaData: siswa,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (api *API) login(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var s Siswa
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := api.siswaRepo.Login(s.Email, s.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	siswa := Siswa{
		Nama:              res.Nama,
		Email:             res.Email,
		JenjangPendidikan: res.JenjangPendidikan,
		Nik:               res.Nik,
		TempatLahir:       res.TempatLahir,
		TanggalLahir:      res.TanggalLahir,
	}
	expTime := time.Now().Add(60 * time.Minute)
	tokenString, err := api.GenerateSiswaToken(siswa, expTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginSuccessResponse{Email: res.Email, Token: tokenString})
}

func (api *API) register(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	encoder := json.NewEncoder(w)
	var s Siswa
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: "Failed to register"})
		return
	}
	res, err := api.siswaRepo.Register(s.Nama, s.Password, s.Email, s.JenjangPendidikan, s.Nik, s.TempatLahir, s.TanggalLahir)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		//Email: res.Email,
		SiswaData: Siswa{
			Email:             res.Email,
			Nama:              res.Nama,
			JenjangPendidikan: res.JenjangPendidikan,
			Nik:               res.Nik,
			TempatLahir:       res.TempatLahir,
			TanggalLahir:      res.TanggalLahir,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	json.NewEncoder(w).Encode(RegisterSuccessResponse{Nama: res.Nama, Email: res.Email, Token: tokenString})
}
