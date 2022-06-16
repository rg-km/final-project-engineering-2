package api

import (
	"fmt"
	"net/http"

	"final-project-eng2-be/repository"
)

type API struct {
	mux          *http.ServeMux
	siswaRepo    repository.SiswaRepository
	beasiswaRepo repository.BeasiswaRepository
}

func NewApi(siswaRepo repository.SiswaRepository, beasiswaRepo repository.BeasiswaRepository) *API {
	mux := http.NewServeMux()

	api := &API{
		mux:          mux,
		siswaRepo:    siswaRepo,
		beasiswaRepo: beasiswaRepo,
	}

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))

	mux.Handle("/api/beasiswa", api.GET(http.HandlerFunc(api.getBeasiswa)))
	mux.Handle("/api/beasiswa/", api.GET(http.HandlerFunc(api.getBeasiswaById)))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", api.Handler())
}
