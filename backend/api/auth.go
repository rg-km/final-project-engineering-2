package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type Siswa struct {
	Id                string `json:"id"` //untuk update
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	Nama              string `json:"nama" validate:"required"`
	JenjangPendidikan string `json:"jenjang_pendidikan" validate:"required"`
	Nik               string `json:"nik" validate:"required"`
	TempatLahir       string `json:"tempat_lahir" validate:"required"`
	TanggalLahir      string `json:"tanggal_lahir" validate:"required"`
	KotaDomisili      string `json:"kota_domisili" validate:"required"`
}

type LoginSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
type LoginSiswa struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterSuccessResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}
type ValidationErrorResponse struct {
	Error []string `json:"error"`
}

type LogoutSuccessResponse struct{
	Message string `json:"message"`
}

var jwtKey = []byte("key")

type Claims struct {
	// Email string `json:"email"`
	SiswaData Siswa `json:"siswa_data"`
	jwt.StandardClaims
}

func formatValidationError(err error) []string {
	var ve validator.ValidationErrors
	out := []string{}
	var msg string

	if errors.As(err, &ve) {
		out = make([]string, len(ve))
		for i, fe := range ve {
			tag := fe.Tag()
			field := fe.Field()

			switch tag {
			case "required":
				msg = fmt.Sprintf("%s is required", field)
			case "email":
				msg = fmt.Sprintf("%s is not a valid email", field)
			}
			out[i] = msg

		}
	}
	return out
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
	// var s Siswa
	var ls LoginSiswa

	err := json.NewDecoder(r.Body).Decode(&ls)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(ls)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errors := formatValidationError(err)
		json.NewEncoder(w).Encode(ValidationErrorResponse{Error: errors})
		return
	}

	res, err := api.siswaRepo.Login(ls.Email, ls.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	idStr := strconv.Itoa(int(res.Id))
	siswa := Siswa{
		Id:                idStr,
		Nama:              res.Nama,
		Email:             res.Email,
		JenjangPendidikan: res.JenjangPendidikan,
		Nik:               res.Nik,
		TempatLahir:       res.TempatLahir,
		TanggalLahir:      res.TanggalLahir,
		KotaDomisili:      res.KotaDomisili,
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

	err = validator.New().Struct(s)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		errors := formatValidationError(err)
		json.NewEncoder(w).Encode(ValidationErrorResponse{Error: errors})
		return
	}

	res, err := api.siswaRepo.Register(s.Nama, s.Password, s.Email, s.JenjangPendidikan, s.Nik, s.TempatLahir, s.TanggalLahir, s.KotaDomisili)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		// Email: res.Email,
		SiswaData: Siswa{
			Email:             res.Email,
			Nama:              res.Nama,
			JenjangPendidikan: res.JenjangPendidikan,
			Nik:               res.Nik,
			TempatLahir:       res.TempatLahir,
			TanggalLahir:      res.TanggalLahir,
			KotaDomisili:      res.KotaDomisili,
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

func (api *API) logout(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LogoutSuccessResponse{Message: "Logout success"})
}
