package penggunaan

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Penggunaan) (int64, error)
	GetOneByID(id int64) (*model.Penggunaan, error)
	GetAllByID(id int64) ([]*model.Penggunaan, error)
	UpdateOneByID(data *model.Penggunaan) (int64, error)
	UpdateStatus(id int64) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error)
	getTotalCount() (totalEntries int)
}

type repository struct {
	DB *sqlx.DB
}

// NewRepository ...
func NewRepository() Repository {
	return &repository{
		DB: helper.GetConnection(),
	}
}

func (m *repository) getTotalCount() (totalEntries int) {
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM penggunaan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Penggunaan) (int64, error) {
	query := `INSERT INTO penggunaan(
		id_pelanggan, bulan, tahun, meter_awal, meter_akhir, status
	) VALUES(?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.IDPelanggan,
		&data.Bulan,
		&data.Tahun,
		&data.Meter_awal,
		&data.Meter_akhir,
	)

	if err != nil {
		return -1, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return lastID, nil
}

func (m *repository) UpdateStatus(id int64) (int64, error) {
	query := `UPDATE penggunaan set status = "Sudah Bayar"
	WHERE id_penggunaan = ?`

	res, err := m.DB.Exec(query,
		id,
	)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func (m *repository) UpdateOneByID(data *model.Penggunaan) (int64, error) {
	query := `UPDATE penggunaan set
	id_pelanggan = ?, bulan = ?, tahun = ?, meter_awal = ?, meter_akhir = ?
	WHERE id_penggunaan = ?`

	res, err := m.DB.Exec(query,
		data.IDPelanggan,
		data.Bulan,
		data.Tahun,
		data.Meter_awal,
		data.Meter_akhir,
		data.ID,
	)

	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}

func (m *repository) GetOneByID(id int64) (*model.Penggunaan, error) {
	query := `SELECT 
	id_penggunaan, 
	id_pelanggan, 
	bulan, 
	tahun, 
	meter_awal, 
	meter_akhir, status
	FROM penggunaan  
	WHERE id_penggunaan = ?`

	data := &model.Penggunaan{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.IDPelanggan,
		&data.Bulan,
		&data.Tahun,
		&data.Meter_awal,
		&data.Meter_akhir,
		&data.Status,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAllByID(id int64) ([]*model.Penggunaan, error) {
	var (
		list_data = make([]*model.Penggunaan, 0)
	)

	query := `SELECT 
	id_penggunaan, 
	id_pelanggan, 
	bulan, 
	tahun, 
	meter_awal, 
	meter_akhir,
	status
	FROM penggunaan 
	WHERE id_penggunaan = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			dataPenggunaan model.Penggunaan
		)

		if err := m.DB.QueryRow(query, id).Scan(
			&dataPenggunaan.ID,
			&dataPenggunaan.IDPelanggan,
			&dataPenggunaan.Bulan,
			&dataPenggunaan.Tahun,
			&dataPenggunaan.Meter_awal,
			&dataPenggunaan.Meter_akhir,
			&dataPenggunaan.Status,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &dataPenggunaan)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error) {
	var (
		list = make([]*model.Penggunaan, 0)
	)

	query := `SELECT 
	a.id_penggunaan, 
	a.id_pelanggan, 
	a.bulan, 
	a.tahun, 
	a.meter_awal, 
	a.meter_akhir,
	a.status,
	b.username, b.nomor_kwh, b.nama_pelanggan, b.alamat
	FROM penggunaan as a
	LEFT JOIN pelanggan as b on b.id_pelanggan = a.id_pelanggan `

	if dqp.Search != "" {
		query += ` WHERE a.status = "Belum Bayar"`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data          model.Penggunaan
			dataPelanggan model.Pelanggan
		)

		if err := rows.Scan(
			&data.ID,
			&data.IDPelanggan,
			&data.Bulan,
			&data.Tahun,
			&data.Meter_awal,
			&data.Meter_akhir,
			&data.Status,
			&dataPelanggan.Username,
			&dataPelanggan.Nomor_kwh,
			&dataPelanggan.Nama_pelanggan,
			&dataPelanggan.Alamat,
		); err != nil {
			return nil, -1, err
		}

		data.PelangganDetail = &model.Pelanggan{
			ID:             data.IDPelanggan,
			Username:       dataPelanggan.Username,
			Nama_pelanggan: dataPelanggan.Nama_pelanggan,
			Alamat:         dataPelanggan.Alamat,
			Nomor_kwh:      dataPelanggan.Nomor_kwh,
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM penggunaan WHERE id_penggunaan = ?`

	res, err := m.DB.Exec(query, id)
	if err != nil {
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}

	return rowsAffected, nil
}
