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
		deskripsi TEXT);
		
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
			
		INSERT INTO beasiswa (nama, jenis_beasiswa, jenjang_pendidikan, tanggal_mulai, tanggal_selesai, deskripsi)
		VALUES ('Beasiswa Pertama', 'Dalam Negeri', 'S1', '2020-01-01', '2020-01-01','Beasiswa ini adalah beasiswa pertama'),
		('Beasiswa Kedua', 'Luar Negeri', 'S1', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kedua'),
    ('Beasiswa Ketiga', 'Dalam Negeri', 'S2', '2019-02-20', '2020-05-20', 'Beasiswa ini adalah beasiswa ketiga'),
    ('Beasiswa Keempat', 'Luar Negeri', 'S1', '2020-05-02', '2020-06-02', 'Beasiswa ini adalah beasiswa keempat'),
    ('Beasiswa ke-5', 'Dalam Negeri', 'S2', '2020-03-01', '2020-07-05', 'Beasiswa ini adalah beasiswa ke-5'),
    ('Beasiswa ke-6', 'Dalam Negeri', 'S1', '2020-03-01', '2020-07-05', 'Beasiswa ini adalah beasiswa ke-6'),
    ('Beasiswa ke-7', 'Luar Negeri', 'S2', '2020-04-01', '2020-08-01', 'Beasiswa ini adalah beasiswa ke-7'),
    ('Beasiswa ke-8', 'Dalam Negeri', 'S2', '2019-08-13', '2019-10-13', 'Beasiswa ini adalah beasiswa ke-8');
		
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
