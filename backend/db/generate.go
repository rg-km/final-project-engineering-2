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
		program_pendidikan TEXT,
		tanggal_mulai TEXT,
		tanggal_selesai TEXT,
		deskripsi TEXT,
    	lama_program TEXT,
		url_gambar);

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

    INSERT INTO beasiswa (nama, jenis_beasiswa, jenjang_pendidikan, program_pendidikan , tanggal_mulai, tanggal_selesai, deskripsi, lama_program, url_gambar) VALUES
	('Beasiswa Pertama', 'Dalam Negeri', 'S1', 'Bisnis & Perdagangan', '2020-01-01', '2020-01-01','Beasiswa ini adalah beasiswa pertama', '4 Tahun', 'https://upload.wikimedia.org/wikipedia/id/7/73/Lambang_Universitas_Negeri_Malang.jpg'),
    ('Beasiswa Ketiga', 'Dalam Negeri', 'S1', 'Arsitektur/Lanskap', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa ketiga', '4 Tahun', 'https://1.bp.blogspot.com/-9NtlLXYPf8Y/YD42i-uzQVI/AAAAAAAAJXY/mX7vXVoSJtQfM4Za-xg-YtuZo64wKIM7wCLcBGAsYHQ/w320-h319/UNIVERSITAS%2BMATARAM.png'),
	('Beasiswa Kedua', 'Luar Negeri', 'S1', 'Rekayasa',  '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kedua', '2 Tahun', 'https://upload.wikimedia.org/wikipedia/en/thumb/6/6e/University_of_Waterloo_seal.svg/800px-University_of_Waterloo_seal.svg.png'),
    ('Beasiswa Keempat', 'Dalam Negeri', 'D3', 'Tenkik Informatikan', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa keempat', '3 Tahun', 'https://upload.wikimedia.org/wikipedia/id/9/95/Logo_Institut_Teknologi_Bandung.png'),
    ('Beasiswa kedelapan','Dalam Negeri', 'D3', 'Tenkik Informatikan', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kedelapan', '3 Tahun','https://upload.wikimedia.org/wikipedia/id/c/c4/Badge_ITS.png'),
    ('Beasiswa Ketujuh', 'Dalam Negeri', 'S2', 'Seni Rupa & Terapan', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa ketujuh', '2 Tahun', 'https://upload.wikimedia.org/wikipedia/commons/5/51/Logo_of_Universitas_Negeri_Semarang.jpg'),
    ('Beasiswa Kelima', 'Luar Negeri', 'S2', 'Ilmu Sosial', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa kelima', '2 Tahun', 'https://upload.wikimedia.org/wikipedia/en/thumb/b/b7/Stanford_University_seal_2003.svg/1200px-Stanford_University_seal_2003.svg.png'),
    ('Beasiswa Keenam', 'Luar Negeri', 'S1', 'Ilmu Komputer', '2020-01-01', '2020-01-01', 'Beasiswa ini adalah beasiswa keenam', '2 Tahun', 'https://upload.wikimedia.org/wikipedia/en/thumb/b/b7/Stanford_University_seal_2003.svg/1200px-Stanford_University_seal_2003.svg.png');

	INSERT INTO siswa (nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir, kota_domisili) VALUES
	('Siswa Pertama', '12345', 'ex@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Jakarta'),
	('Siswa Kedua', '12345', 'contoh@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-3', '12345', 'c@gmail.com', 'S1', '1234568', '2020-01-01', 'Jakarta', 'Bandung'),
    ('Siswa Ke-4', '12345', 'ch@gmail.com', 'S1', '1234589', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-5', '12345', 'coth@gmail.com', 'S1', '456789', '2020-01-01', 'Jakarta', 'Surabaya'),
    ('Siswa Ke-6', '1294', 'coh@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta', 'Surabaya');`)
	if err != nil {
		panic(err)
	}
}
