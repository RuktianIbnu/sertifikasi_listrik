package tagihan

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Tagihan) (int64, error)
	GetOneByID(id int64) ([]*model.Tagihan, error)
	// GetAllByID(id int64) (*model.Tagihan, error)
	UpdateOneByID(data *model.Tagihan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM tagihan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Tagihan) (int64, error) {
	query := `INSERT INTO tagihan(
		id_penggunaan, id_pelanggan, bulan, tahun, jumlah_meter, status) VALUES(?, ?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.Id_penggunaan,
		&data.Id_pelanggan,
		&data.Bulan,
		&data.Tahun,
		&data.Jumlah_meter,
		&data.Status,
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

func (m *repository) UpdateOneByID(data *model.Tagihan) (int64, error) {
	query := `UPDATE Tagihan set
	id_penggunaan = ?, id_pelanggan = ?, bulan = ?, tahun = ?, jumlah_meter = ?, status = ?
	WHERE id_tagihan = ?`

	res, err := m.DB.Exec(query,
		data.Id_penggunaan,
		data.Id_pelanggan,
		data.Bulan,
		data.Tahun,
		data.Jumlah_meter,
		data.Status,
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

func (m *repository) GetOneByID(id_tagihan int64) ([]*model.Tagihan, error) {
	var (
		list_data = make([]*model.Tagihan, 0)
	)

	query := `SELECT 
	id_tagihan, 
	id_penggunaan,
	id_pelanggan,
	bulan,
	tahun,
	jumlah_meter,
	status
	FROM tagihan  
	WHERE id_tagihan = ?`

	rows, err := m.DB.Query(query, id_tagihan)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Tagihan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Id_penggunaan,
			&data.Id_pelanggan,
			&data.Bulan,
			&data.Tahun,
			&data.Jumlah_meter,
			&data.Status,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

// func (m *repository) GetAllByID(id int64) (*model.Tagihan, error) {
// 	query := `SELECT
// 	id,
// 	COALESCE(kode_seksi, ''),
// 	COALESCE(nama_seksi, ''),
// 	id_parent_subdirektorat
// 	FROM ms_seksi
// 	WHERE id = ?`

// 	data := &model.Tagihan{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Kodesubdirektorat,
// 		&data.Namasubdirektorat,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error) {
	var (
		list = make([]*model.Tagihan, 0)
	)

	query := `SELECT 
	id_tagihan,
	id_penggunaan,
	id_pelanggan,
	bulan,
	tahun,
	jumlah_meter,
	status 
	FROM tagihan`

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
			data model.Tagihan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Id_penggunaan,
			&data.Id_pelanggan,
			&data.Bulan,
			&data.Tahun,
			&data.Jumlah_meter,
			&data.Status,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM tagihan WHERE id_tagihan = ?`

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
