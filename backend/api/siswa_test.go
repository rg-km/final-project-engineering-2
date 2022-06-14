package api

import (
	"database/sql"
	"encoding/json"
	"final-project-eng2-be/db"
	"final-project-eng2-be/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http/httptest"
	"os"
)

var _ = Describe("Siswa testing", func() {
	AfterEach(func() {
		os.Remove("./beasiswa.db")
	})

	Describe("Siswa Data", func() {

		It("Should Return all Siswa", func() {
			db, err := sql.Open("sqlite3", "./beasiswa.db")
			Expect(err).To(BeNil())
			database.Migrate(db)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/siswa/all", nil)
			siswaRepo := repository.NewSiswaRepository(db)
			api := NewApi(*siswaRepo)
			api.GetAllSiswa(w, r)

			res := w.Result()
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())

			var data []Siswa

			err = json.Unmarshal(body, &data)
			Expect(err).To(BeNil())

			Expect(len(data)).To(Equal(2))

		})

		It("Should Return Empty list", func() {
			db, err := sql.Open("sqlite3", "./beasiswa.db")
			Expect(err).To(BeNil())

			db.Exec(`CREATE TABLE IF NOT EXISTS beasiswa (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				nama TEXT,
				jenis_beasiswa TEXT,
				jenjang_pendidikan TEXT,
				tanggal_mulai TEXT,
				tanggal_selesai TEXT);
				
				CREATE TABLE IF NOT EXISTS siswa (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					nama TEXT,
					password TEXT,
					email TEXT,
					jenjang_pendidikan TEXT,
					nik TEXT,
					tanggal_lahir TEXT,
					tempat_lahir TEXT);
				
				CREATE TABLE IF NOT EXISTS pendaftaran (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					id_beasiswa INTEGER,
					id_siswa INTEGER,
					tanggal_daftar TEXT,
					status TEXT);
				
				CREATE TABLE IF NOT EXISTS mitra (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					nama TEXT,
					email TEXT,
					lokasi TEXT,
					no_telp TEXT,
					legalitas TEXT);`)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/api/siswa/all", nil)
			siswaRepo := repository.NewSiswaRepository(db)
			api := NewApi(*siswaRepo)
			api.GetAllSiswa(w, r)

			res := w.Result()
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			Expect(err).To(BeNil())

			var data []Siswa

			err = json.Unmarshal(body, &data)
			Expect(err).To(BeNil())

			Expect(len(data)).To(Equal(0))

		})

	})

})
