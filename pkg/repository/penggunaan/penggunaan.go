package penggunaan

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Penggunaan) (int64, error)
	GetOneByID(id int64) ([]*model.Penggunaan, error)
	GetAllByID(id int64) (*model.Penggunaan, error)
	UpdateOneByID(data *model.Penggunaan) (int64, error)
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
		id_pelanggan, bulan, tahun, meter_awal, meter_akhir
	) VALUES(?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.Id_pelanggan,
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

func (m *repository) UpdateOneByID(data *model.Penggunaan) (int64, error) {
	query := `UPDATE penggunaan set
	id_pelanggan = ?, bulan = ?, tahun = ?, meter_awal = ?, meter_akhir = ?
	WHERE id_penggunaan = ?`

	res, err := m.DB.Exec(query,
		data.Id_pelanggan,
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

func (m *repository) GetAllByID(id int64) (*model.Penggunaan, error) {
	query := `SELECT 
	id_penggunaan, 
	id_pelanggan, 
	bulan, 
	tahun, 
	meter_awal, 
	meter_akhir
	FROM penggunaan  
	WHERE id = ?`

	data := &model.Penggunaan{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Id_pelanggan,
		&data.Bulan,
		&data.Tahun,
		&data.Meter_awal,
		&data.Meter_akhir,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetOneByID(id int64) ([]*model.Penggunaan, error) {
	var (
		list_data = make([]*model.Penggunaan, 0)
	)

	query := `SELECT 
	id_penggunaan, 
	id_pelanggan, 
	bulan, 
	tahun, 
	meter_awal, 
	meter_akhir
	FROM penggunaan  
	WHERE id_penggunaan = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Penggunaan
		)

		if err := m.DB.QueryRow(query, id).Scan(
			&data.ID,
			&data.Id_pelanggan,
			&data.Bulan,
			&data.Tahun,
			&data.Meter_awal,
			&data.Meter_akhir,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error) {
	var (
		list = make([]*model.Penggunaan, 0)
	)

	query := `SELECT 
	id_penggunaan, 
	id_pelanggan, 
	bulan, 
	tahun, 
	meter_awal, 
	meter_akhir
	FROM penggunaan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(bulan, tahun) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Penggunaan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Id_pelanggan,
			&data.Bulan,
			&data.Tahun,
			&data.Meter_awal,
			&data.Meter_akhir,
		); err != nil {
			return nil, -1, err
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