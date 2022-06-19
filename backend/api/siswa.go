package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"
)

type SiswaErrorResponse struct {
	Error string `json:"error"`
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
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(SiswaErrorResponse{Error: "Internal server error"})
		return
	}

	response, err := json.Marshal(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(SiswaErrorResponse{Error: "Internal server error"})
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
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
	w.WriteHeader(http.StatusOK)
	encoder.Encode(res)
	return
}
