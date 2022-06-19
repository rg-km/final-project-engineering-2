package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
)

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
	SiswaClaims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, SiswaClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: "Invalid Token"})
		return
	}

	retrivedSiswa := SiswaClaims.SiswaData

	expTime := time.Now().Add(60 * time.Minute)
	newTokenString, err := a.GenerateSiswaToken(retrivedSiswa, expTime)

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
