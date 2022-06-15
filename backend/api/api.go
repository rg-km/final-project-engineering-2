package api

import (
	"final-project-eng2-be/repository"
	"fmt"
	"net/http"
)

type API struct {
	mux       *http.ServeMux
	siswaRepo repository.SiswaRepository
}

func NewApi(siswaRepo repository.SiswaRepository) *API {
	mux := http.NewServeMux()

	api := &API{
		mux:       mux,
		siswaRepo: siswaRepo,
	}

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))
	mux.Handle("/api/siswa/all", api.GET(http.HandlerFunc(api.GetAllSiswa)))
	mux.Handle("/api/siswa", api.GET(http.HandlerFunc(api.GetSiswaByID)))
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", api.Handler())
}
