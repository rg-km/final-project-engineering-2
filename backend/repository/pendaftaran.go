package repository

import (
	"database/sql"
	"sync"
)

type PendaftaranRepository struct {
	mu *sync.Mutex
	db *sql.DB
}

func NewPendaftaranRepository(db *sql.DB) *PendaftaranRepository {
	return &PendaftaranRepository{
		db: db,
		mu: &sync.Mutex{},
	}
}

func (r *PendaftaranRepository) GetPendaftaranAll() ([]Pendaftaran, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []Pendaftaran

	sqlStatement := "SELECT * FROM pendaftaran"

	rows, err := r.db.Query(sqlStatement)

	if err != nil {
		return []Pendaftaran{}, err
	}
	for rows.Next() {
		var p Pendaftaran
		rows.Scan(&p.Id, &p.IdBeasiswa, &p.IdSiswa, &p.TanggalDaftar, &p.Status)
		result = append(result, p)
	}

	return result, nil
}

func (r *PendaftaranRepository) GetPendaftaranById(id int) (Pendaftaran, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var p Pendaftaran
	err := r.db.QueryRow("SELECT * FROM pendaftaran WHERE id = ?", id).Scan(&p.Id, &p.IdBeasiswa, &p.IdSiswa, &p.TanggalDaftar, &p.Status)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (r *PendaftaranRepository) GetBySiswa(idSiswa int) ([]Pendaftaran, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []Pendaftaran

	sqlStatement := "SELECT * FROM pendaftaran WHERE id_siswa = ?"

	rows, err := r.db.Query(sqlStatement, idSiswa)
	if err != nil {
		return []Pendaftaran{}, err
	}

	for rows.Next() {
		var p Pendaftaran
		rows.Scan(&p.Id, &p.IdBeasiswa, &p.IdSiswa, &p.TanggalDaftar, &p.Status)
		result = append(result, p)
	}
	return result, nil
}

func (r *PendaftaranRepository) GetByBeasiswa(idBeasiswa int) ([]Pendaftaran, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []Pendaftaran

	sqlStatement := "SELECT * FROM pendaftaran WHERE id_beasiswa = ?"

	rows, err := r.db.Query(sqlStatement, idBeasiswa)
	if err != nil {
		return []Pendaftaran{}, err
	}

	for rows.Next() {
		var p Pendaftaran
		rows.Scan(&p.Id, &p.IdBeasiswa, &p.IdSiswa, &p.TanggalDaftar, &p.Status)
		result = append(result, p)
	}
	return result, nil
}

func (r *PendaftaranRepository) CreatePendaftaran(idBeasiswa int, idSiswa int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sqlStatement := "INSERT INTO pendaftaran (id_beasiswa, id_siswa, tanggal_daftar, status) VALUES (?, ?, date('now'), ?)"

	_, err := r.db.Exec(sqlStatement, idBeasiswa, idSiswa, status)
	if err != nil {
		return err
	}

	return nil
}

