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

	})

})
