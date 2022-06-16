package repository

import (
	"database/sql"
	"sync"
)

type SiswaRepository struct {
	mu *sync.Mutex
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) *SiswaRepository {
	return &SiswaRepository{
		db: db,
		mu: &sync.Mutex{},
	}
}

func (r *SiswaRepository) Login(email string, password string) (Siswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var s Siswa
	err := r.db.QueryRow("SELECT * FROM siswa WHERE email = ? AND password = ?", email, password).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) Register(nama string, password string, email string, jenjangPendidikan string, nik string, tanggalLahir string, tempatLahir string) (Siswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var s Siswa
	err := r.db.QueryRow("INSERT INTO siswa (nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir", nama, password, email, jenjangPendidikan, nik, tanggalLahir, tempatLahir).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) GetAll() ([]Siswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []Siswa

	sqlStatement := "SELECT * FROM siswa"

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		return []Siswa{}, err
	}

	for rows.Next() {
		var s Siswa
		rows.Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
		result = append(result, s)
	}

	return result, nil
}
func (r *SiswaRepository) GetSiswaByID(id int) (*Siswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var s Siswa

	sqlStatement := "SELECT * FROM siswa WHERE id = ?"

	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
