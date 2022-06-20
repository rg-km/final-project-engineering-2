package api

import (
	"fmt"
	"net/http"

	"final-project-eng2-be/repository"
)

type API struct {
	mux             *http.ServeMux
	siswaRepo       repository.SiswaRepository
	beasiswaRepo    repository.BeasiswaRepository
	pendaftaranRepo repository.PendaftaranRepository
}

func NewApi(siswaRepo repository.SiswaRepository, beasiswaRepo repository.BeasiswaRepository, pendaftaranRepo repository.PendaftaranRepository) *API {
	mux := http.NewServeMux()

	api := &API{
		mux:             mux,
		siswaRepo:       siswaRepo,
		beasiswaRepo:    beasiswaRepo,
		pendaftaranRepo: pendaftaranRepo,
	}

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))

	mux.Handle("/api/beasiswa/all", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getBeasiswa))))
	mux.Handle("/api/beasiswa", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getBeasiswaById))))

	mux.Handle("/api/siswa/all", api.GET(http.HandlerFunc(api.GetAllSiswa)))
	mux.Handle("/api/siswa", api.GET(http.HandlerFunc(api.GetSiswaByID)))
	mux.Handle("/api/siswa/token", api.GET(api.AuthMiddleware(http.HandlerFunc(api.GetSiswaByToken))))
	mux.Handle("/api/siswa/update", api.POST(api.AuthMiddleware(http.HandlerFunc(api.UpdateSiswa))))

	mux.Handle("/api/pendaftaran/all", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getAllPendaftaran))))
	mux.Handle("/api/pendaftaran", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getPendaftaranById))))
	mux.Handle("/api/pendaftaran/siswa", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getPendaftaranBySiswa))))
	mux.Handle("/api/pendaftaran/beasiswa", api.GET(api.AuthMiddleware(http.HandlerFunc(api.getPendaftaranByBeasiswa))))

	mux.Handle("/api/pendaftaran/create", api.POST(api.AuthMiddleware(http.HandlerFunc(api.createPendaftaran))))
	mux.Handle("/api/pendaftaran/update", api.POST(api.AuthMiddleware(http.HandlerFunc(api.updatePendaftaran))))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", api.Handler())
}
