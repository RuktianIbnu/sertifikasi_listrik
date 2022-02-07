package model

import "time"

// DefaultQueryParam ...
type DefaultQueryParam struct {
	Search  string
	Page    int
	Limit   int
	Offset  int
	Sorting map[string]string
	Params  map[string]interface{}
}

// User ...
type User struct {
	ID         int64  `json:"id_user"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Nama_admin string `json:"nama_admin"`
	IDLevel    int64  `json:"id_level,omitempty"`
}

// Tarif ...
type Tarif struct {
	ID    int64   `json:"id_tarif"`
	Daya  int64   `json:"daya,omitempty"`
	Tarif float64 `json:"tarifperkwh,omitempty"`
}

// Tagihan ...
type Tagihan struct {
	ID            int64  `json:"id_tagihan"`
	Id_penggunaan int64  `json:"id_penggunaan,omitempty"`
	Id_pelanggan  int64  `json:"id_pelanggan,omitempty"`
	Bulan         string `json:"bulan"`
	Tahun         string `json:"tahun"`
	Jumlah_meter  int64  `json:"jumlah_meter"`
	Status        string `json:"status"`
}

// Ruang ...
type Ruang struct {
	ID        int64  `json:"id"`
	Namaseksi string `json:"nama_ruang,omitempty"`
	Deskripsi string `json:"deskripsi,omitempty"`
}

// Roles ...
type Roles struct {
	ID       int64  `json:"id"`
	Namarole string `json:"name_role,omitempty"`
}

// Kegiatan ...
type Kegiatan struct {
	ID           int64  `json:"id"`
	Kodekegiatan string `json:"kode_kegiatan,omitempty"`
	Namakegiatan string `json:"nama_kegiatan,omitempty"`
}

// Pendataan Tamu ...
type PendataanTamu struct {
	ID                 int64      `json:"id"`
	Namatamu           string     `json:"nama_tamu,omitempty"`
	Departement        string     `json:"departement,omitempty"`
	Jumlah             int64      `json:"jumlah"`
	Lokasi             string     `json:"lokasi,omitempty"`
	Detaillokasi       string     `json:"detail_lokasi,omitempty"`
	Tanggalmulai       *time.Time `json:"tanggal_mulai"`
	Tanggalselesai     *time.Time `json:"tanggal_selesai"`
	Starttime          *time.Time `json:"start_time"`
	Endtime            *time.Time `json:"end_time"`
	Kategori           string     `json:"kategori,omitempty"`
	Lainlain           string     `json:"lain-lain,omitempty"`
	Deskripsi          string     `json:"deskripsi,omitempty"`
	Efek               string     `json:"efek,omitempty"`
	Resiko             string     `json:"resiko,omitempty"`
	IDpetugas          int64      `json:"id_petugas,omitempty"`
	IDpetugas2         int64      `json:"id_petugas2,omitempty"`
	Status             string     `json:"status,omitempty"`
	Photopemberitahuan string     `json:"photo_pemberitahuan,omitempty"`
	Typepemberitahuan  string     `json:"type_pemberitahuan,omitempty"`
	Photoperintah      string     `json:"photo_perintah,omitempty"`
	Typeperintah       string     `json:"type_perintah,omitempty"`
	Photokegiatan      string     `json:"photo_kegiatan,omitempty"`
	Petugas            *User      `json:"user"`
	Petugas2           *User      `json:"user"`
}
