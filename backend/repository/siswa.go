package repository

import "database/sql"

type SiswaRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) *SiswaRepository {
	return &SiswaRepository{db: db}
}

func (r *SiswaRepository) Login(email string, password string) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("SELECT * FROM siswa WHERE email = ? AND password = ?", email, password).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

