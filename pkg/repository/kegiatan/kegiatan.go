package kegiatan

import (
	"epiket-api/pkg/helper"
	"epiket-api/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Kegiatan) (int64, error)
	GetOneByID(id int64) ([]*model.Kegiatan, error)
	// GetAllByID(id int64) (*model.Kegiatan, error)
	UpdateOneByID(data *model.Kegiatan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Kegiatan, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_kegiatan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Kegiatan) (int64, error) {
	query := `INSERT INTO ms_kegiatan(
		kode_kegiatan, nama_kegiatan) VALUES(?, ?)`

	res, err := m.DB.Exec(query,
		&data.Kodekegiatan,
		&data.Namakegiatan,
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

func (m *repository) UpdateOneByID(data *model.Kegiatan) (int64, error) {
	query := `UPDATE ms_kegiatan set
	kode_kegiatan = ?, nama_kegiatan = ?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Kodekegiatan,
		&data.Namakegiatan,
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

func (m *repository) GetOneByID(id int64) ([]*model.Kegiatan, error) {
	var (
		list_data = make([]*model.Kegiatan, 0)
	)

	query := `SELECT 
	id, 
	kode_kegiatan, 
	COALESCE(nama_kegiatan, '')
	FROM ms_kegiatan  
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Kegiatan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Kodekegiatan,
			&data.Namakegiatan,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

// func (m *repository) GetAllByID(id int64) (*model.Kegiatan, error) {
// 	query := `SELECT
// 	id,
// 	COALESCE(kode_seksi, ''),
// 	COALESCE(nama_seksi, ''),
// 	id_parent_subdirektorat
// 	FROM ms_kegiatan
// 	WHERE id = ?`

// 	data := &model.Kegiatan{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Kodesubdirektorat,
// 		&data.Namasubdirektorat,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Kegiatan, int, error) {
	var (
		list = make([]*model.Kegiatan, 0)
	)

	query := `SELECT id, kode_kegiatan, COALESCE(nama_kegiatan, '') FROM ms_kegiatan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(kode_kegiatan, nama_kegiatan) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Kegiatan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Kodekegiatan,
			&data.Namakegiatan,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM ms_kegiatan WHERE id = ?`

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
