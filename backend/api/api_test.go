package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http/httptest"
	"os"

	"final-project-eng2-be/db"
	"final-project-eng2-be/repository"
)

var _ = Describe("Api testing", func() {
	Describe("login", func() {
		When("email and password are correct", func() {

			AfterEach(func() {
				os.Remove("./beasiswa.db")
			})

			It("should return email and token", func() {
				db, err := sql.Open("sqlite3", "./beasiswa.db")
				Expect(err).To(BeNil())
				database.Migrate(db)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(`{"email":"ex@gmail.com","password":"123456"}`)))
				r.Header.Set("Content-Type", "application/json")
				siswaRepo := repository.NewSiswaRepository(db)
				api := NewApi(*siswaRepo)
				api.login(w, r)

				res := w.Result()
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())

				var data map[string]string
				err = json.Unmarshal(body, &data)
				Expect(err).To(BeNil())

				Expect(data["email"]).To(Equal("ex@gmail.com"))
			})
		})
		When("email and password are incorrect", func() {
			It("should return error", func() {
				db, err := sql.Open("sqlite3", "./beasiswa.db")
				Expect(err).To(BeNil())
				database.Migrate(db)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer([]byte(`{"email":"ex@gmail.com","password":"1234567"}`)))
				r.Header.Set("Content-Type", "application/json")
				api := NewApi(*repository.NewSiswaRepository(db))
				api.login(w, r)

				res := w.Result()
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())

				var data map[string]string
				err = json.Unmarshal(body, &data)
				Expect(err).To(BeNil())

				Expect(data["error"]).To(Equal("invalid email or password"))
			})
		})
	})
	Describe("register", func() {
		When("register success", func() {
			It("should return email, name and token", func() {
				db, err := sql.Open("sqlite3", "./beasiswa.db")
				Expect(err).To(BeNil())
				database.Migrate(db)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer([]byte(`{"nama":"ex","email":"ex@gmail.com", "password":"123456", "jenjang_pendidikan":"S1", tempat_lahir":"Bandung", tanggal_lahir":"2020-01-01", nik":"123456789"}`)))
				r.Header.Set("Content-Type", "application/json")

				api := NewApi(*repository.NewSiswaRepository(db))
				api.register(w, r)

				res := w.Result()
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())

				var data map[string]string
				err = json.Unmarshal(body, &data)
				Expect(err).To(BeNil())

				Expect(data["email"]).To(Equal("ex@gmail.com"))
				Expect(data["name"]).To(Equal("ex"))
			})
		})
		When("register failed", func() {
			It("should return error", func() {
				db, err := sql.Open("sqlite3", "./beasiswa.db")
				Expect(err).To(BeNil())
				database.Migrate(db)

				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer([]byte(`{"nama":"ex","email":"ex@gmail.com", "password":"123456",  "jenjang_pendidikan":"S1", tempat_lahir":"Bandung", tanggal_lahir":"2020-01-01", nik":"123456789"}`)))
				r.Header.Set("Content-Type", "application/json")

				api := NewApi(*repository.NewSiswaRepository(db))
				api.register(w, r)

				res := w.Result()
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())

				var data map[string]string
				err = json.Unmarshal(body, &data)
				Expect(err).To(BeNil())

				Expect(data["error"]).To(Equal("Failed to register"))
			})
		})
	})
})
