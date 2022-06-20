package repository

import (
	"database/sql"
	"sync"
)

type PendaftaranResponse struct {
	ID              int    `json:"id"`
	IdSiswa         int    `json:"id_siswa"`
	IdBeasiswa      int    `json:"id_beasiswa"`
	NamaSiswa       string `json:"nama_siswa"`
	EmailSiswa      string `json:"email_siswa"`
	NamaBeasiswa    string `json:"nama_beasiswa"`
	JenjangBeasiswa string `json:"jenjang_beasiswa"`
	TanggalDaftar   string `json:"tanggal_daftar"`
	Status          string `json:"status"`
}

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

func (r *PendaftaranRepository) GetPendaftaranAll() ([]PendaftaranResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []PendaftaranResponse = []PendaftaranResponse{}

	sqlStatement :=
		`SELECT
    p.id as id_pendaftaran,
    s.id as id_siswa,
    b.id as id_beasiswa,
    s.nama AS nama_siswa,
    s.email AS email_siswa,
    b.nama AS nama_beasiswa,
    b.jenjang_pendidikan AS jenjang_pendidikan,
    p.tanggal_daftar AS tanggal_daftar,
	p.status as status
  FROM pendaftaran p
  INNER JOIN siswa s on p.id_siswa = s.id
  INNER JOIN beasiswa b on p.id_beasiswa = b.id`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return []PendaftaranResponse{}, err
	}

	for rows.Next() {
		var p PendaftaranResponse
		rows.Scan(&p.ID, &p.IdSiswa, &p.IdBeasiswa, &p.NamaSiswa, &p.EmailSiswa, &p.NamaBeasiswa, &p.JenjangBeasiswa, &p.TanggalDaftar, &p.Status)
		result = append(result, p)
	}
	return result, nil

}

func (r *PendaftaranRepository) GetPendaftaranById(id int) (PendaftaranResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var p PendaftaranResponse
	err := r.db.QueryRow(
		`SELECT
    p.id as id_pendaftaran,
    s.id as id_siswa,
    b.id as id_beasiswa,
    s.nama AS nama_siswa,
    s.email AS email_siswa,
    b.nama AS nama_beasiswa,
    b.jenjang_pendidikan AS jenjang_pendidikan,
    p.tanggal_daftar AS tanggal_daftar,
	p.status as status
  FROM pendaftaran p
  INNER JOIN siswa s on p.id_siswa = s.id
  INNER JOIN beasiswa b on p.id_beasiswa = b.id
  WHERE p.id = ?`, id).Scan(&p.ID, &p.IdSiswa, &p.IdBeasiswa, &p.NamaSiswa, &p.EmailSiswa, &p.NamaBeasiswa, &p.JenjangBeasiswa, &p.TanggalDaftar, &p.Status)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (r *PendaftaranRepository) GetBySiswa(idSiswa int) ([]PendaftaranResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []PendaftaranResponse = []PendaftaranResponse{}

	sqlStatement :=
		`SELECT
    p.id as id_pendaftaran,
    s.id as id_siswa,
    b.id as id_beasiswa,
    s.nama AS nama_siswa,
    s.email AS email_siswa,
    b.nama AS nama_beasiswa,
    b.jenjang_pendidikan AS jenjang_pendidikan,
    p.tanggal_daftar AS tanggal_daftar,
	p.status as status
  FROM pendaftaran p
  INNER JOIN siswa s on p.id_siswa = s.id
  INNER JOIN beasiswa b on p.id_beasiswa = b.id
  WHERE p.id_siswa = ?`

	rows, err := r.db.Query(sqlStatement, idSiswa)
	if err != nil {
		return []PendaftaranResponse{}, err
	}

	for rows.Next() {
		var p PendaftaranResponse
		rows.Scan(&p.ID, &p.IdSiswa, &p.IdBeasiswa, &p.NamaSiswa, &p.EmailSiswa, &p.NamaBeasiswa, &p.JenjangBeasiswa, &p.TanggalDaftar, &p.Status)
		result = append(result, p)
	}
	return result, nil
}

func (r *PendaftaranRepository) GetByBeasiswa(idBeasiswa int) ([]PendaftaranResponse, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []PendaftaranResponse = []PendaftaranResponse{}

	sqlStatement :=
		`SELECT
    p.id as id_pendaftaran,
    s.id as id_siswa,
    b.id as id_beasiswa,
    s.nama AS nama_siswa,
    s.email AS email_siswa,
    b.nama AS nama_beasiswa,
    b.jenjang_pendidikan AS jenjang_pendidikan,
    p.tanggal_daftar AS tanggal_daftar,
	p.status as status
  FROM pendaftaran p
  INNER JOIN siswa s on p.id_siswa = s.id
  INNER JOIN beasiswa b on p.id_beasiswa = b.id
  WHERE p.id_beasiswa = ?`

	rows, err := r.db.Query(sqlStatement, idBeasiswa)
	if err != nil {
		return []PendaftaranResponse{}, err
	}

	for rows.Next() {
		var p PendaftaranResponse
		rows.Scan(&p.ID, &p.IdSiswa, &p.IdBeasiswa, &p.NamaSiswa, &p.EmailSiswa, &p.NamaBeasiswa, &p.JenjangBeasiswa, &p.TanggalDaftar, &p.Status)
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

func (r *PendaftaranRepository) UpdatePendaftaran(id int, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sqlStatement := "UPDATE pendaftaran SET status = ? WHERE id = ?"

	_, err := r.db.Exec(sqlStatement, status, id)
	if err != nil {
		return err
	}

	return nil
}
