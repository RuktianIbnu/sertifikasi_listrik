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
	Username   string `json:"username"`
	Password   string `json:"password"`
	Nama_admin string `json:"nama_admin"`
	IDLevel    int64  `json:"id_level"`
}

// Tarif ...
type Tarif struct {
	ID    int64   `json:"id_tarif"`
	Daya  int64   `json:"daya"`
	Tarif float64 `json:"tarifperkwh"`
}

// Tagihan ...
type Tagihan struct {
	ID            int64  `json:"id_tagihan"`
	Id_penggunaan int64  `json:"id_penggunaan"`
	Id_pelanggan  int64  `json:"id_pelanggan"`
	Bulan         string `json:"bulan"`
	Tahun         string `json:"tahun"`
	Jumlah_meter  int64  `json:"jumlah_meter"`
	Status        string `json:"status"`
}

// Penggunaan ...
type Penggunaan struct {
	ID           int64  `json:"id_penggunaan"`
	Id_pelanggan int64  `json:"id_pelanggan"`
	Bulan        string `json:"bulan"`
	Tahun        string `json:"tahun"`
	Meter_awal   int64  `json:"meter_awal"`
	Meter_akhir  int64  `json:"meter_akhir"`
}

// Pelanggan ...
type Pelanggan struct {
	ID             int64   `json:"id_pelanggan"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Nomor_kwh      int64   `json:"nomor_kwh"`
	Nama_pelanggan string  `json:"nama_pelanggan"`
	Alamat         string  `json:"alamat"`
	Tarif          float64 `json:"id_tarif"`
}

// Pembayaran ...
type Pembayaran struct {
	ID                 int64      `json:"id_pembayaran"`
	Id_tagihan         int64      `json:"id_tagihan"`
	Id_pelanggan       int64      `json:"id_pelanggan"`
	Tanggal_pembayaran *time.Time `json:"tanggal_pembayaran"`
	Bulan_bayar        string     `json:"bulan_bayar"`
	Biaya_admin        string     `json:"biaya_admin"`
	Total_bayar        float64    `json:"total_bayar"`
	Id_user            int64      `json:"id_user"`
}

// Level ...
type Level struct {
	ID         int64  `json:"id_level"`
	Nama_level string `json:"nama_level"`
}
