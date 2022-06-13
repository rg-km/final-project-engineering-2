package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Siswa struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Nama string `json:"nama"`
	JenjangPendidikan string `json:"jenjang_pendidikan"`
	Nik string `json:"nik"`
	TempatLahir string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
}

type LoginSuccessResponse struct {
	Email string `json:"email"`
	Token    string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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
	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email: res.Email,
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
		Name: "token",
		Value: tokenString,
		Expires: expTime,
	})
	
	json.NewEncoder(w).Encode(LoginSuccessResponse{Email: res.Email, Token: tokenString})
}

	
