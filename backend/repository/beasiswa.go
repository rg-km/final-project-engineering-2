package repository

import (
	"database/sql"
	"sync"
)

type BeasiswaRepository struct {
	mu *sync.Mutex
	db *sql.DB
}

func NewBeasiswaRepository(db *sql.DB) *BeasiswaRepository {
	return &BeasiswaRepository{db: db, mu: &sync.Mutex{}}
}

func (r *BeasiswaRepository) GetAll() ([]Beasiswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var beasiswa []Beasiswa

	rows, err := r.db.Query("SELECT * FROM beasiswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Beasiswa
		err := rows.Scan(&b.Id, &b.Nama, &b.JenisBeasiswa, &b.JenjangPendidikan, &b.ProgramPendidikan, &b.TanggalMulai, &b.TanggalSelesai, &b.Deskripsi, &b.LamaProgram, &b.UrlGambar)
		if err != nil {
			return nil, err
		}
		beasiswa = append(beasiswa, b)
	}
	return beasiswa, nil
}

func (r *BeasiswaRepository) GetById(id int64) (Beasiswa, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var b Beasiswa
	err := r.db.QueryRow("SELECT * FROM beasiswa WHERE id = ?", id).Scan(&b.Id, &b.Nama, &b.JenisBeasiswa, &b.JenjangPendidikan, &b.ProgramPendidikan, &b.TanggalMulai, &b.TanggalSelesai, &b.Deskripsi, &b.LamaProgram, &b.UrlGambar)
	if err != nil {
		return b, err
	}
	return b, nil
}
