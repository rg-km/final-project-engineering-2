package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS beasiswa (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama TEXT,
		jenis_beasiswa TEXT,
		jenjang_pendidikan TEXT,
		tanggal_mulai TEXT,
		tanggal_selesai TEXT,
		deskripsi TEXT,
    lama_program INTEGER);
		
		CREATE TABLE IF NOT EXISTS siswa (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama TEXT,
			password TEXT,
			email TEXT,
			jenjang_pendidikan TEXT,
			nik TEXT,
			tanggal_lahir TEXT,
			tempat_lahir TEXT,
			kota_domisili TEXT);
		
		CREATE TABLE IF NOT EXISTS pendaftaran (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_beasiswa INTEGER,
			id_siswa INTEGER,
			tanggal_daftar TEXT,
			status TEXT,
      UNIQUE(id_siswa, id_beasiswa)
    );
		
		CREATE TABLE IF NOT EXISTS mitra (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama TEXT,
			email TEXT,
			lokasi TEXT,
			no_telp TEXT,
			legalitas TEXT);
			
    INSERT INTO beasiswa (nama, jenis_beasiswa, jenjang_pendidikan, tanggal_mulai, tanggal_selesai, deskripsi, lama_program)
		VALUES ('beasiswa Pertama', 'Dalam Negeri', 'S1', '2020-01-01', '2020-01-01','Beasiswa ini adalah beasiswa pertama', 4),
		('beasiswa Kedua', 'Luar Negeri', 'S1', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kedua', 4),
    ('beasiswa Ketiga', 'Dalam Negeri', 'S1', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa ketiga', 4),
    ('easiswa Keempat', 'Dalam Negeri', 'D3', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa keempat', 3),
    ('easiswa Kelima', 'Luar Negeri', 'S2', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kelima', 2),
    ('Beasiswa Keenam', 'Luar Negeri', 'S1', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa keenam', 4),
    ('Beasiswa Ketujuh', 'Dalam Negeri', 'S2', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa ketujuh', 2),
    ('Beasiswa kedelapan','Dalam Negeri', 'D3', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kedelapan', 3);
		
		INSERT INTO siswa (nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir, kota_domisili)
		VALUES ('Siswa Pertama', '12345', 'ex@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Jakarta'),
		('Siswa Kedua', '12345', 'contoh@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-3', '12345', 'c@gmail.com', 'S1', '1234568', '2020-01-01', 'Jakarta', 'Bandung'),
    ('Siswa Ke-4', '12345', 'ch@gmail.com', 'S1', '1234589', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-5', '12345', 'coth@gmail.com', 'S1', '456789', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-6', '1294', 'coh@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Surabaya');`)

	if err != nil {
		panic(err)
	}
}
