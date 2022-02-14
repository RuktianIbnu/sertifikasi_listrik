package level

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Level) (int64, error)
	GetOneByID(id int64) ([]*model.Level, error)
	// GetAllByID(id int64) (*model.Level, error)
	UpdateOneByID(data *model.Level) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Level, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM level").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Level) (int64, error) {
	query := `INSERT INTO level(
		nama_level) VALUES(?)`

	res, err := m.DB.Exec(query,
		&data.Nama_level,
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

func (m *repository) UpdateOneByID(data *model.Level) (int64, error) {
	query := `UPDATE level set nama_level = ?
	WHERE id_level = ?`

	res, err := m.DB.Exec(query,
		&data.Nama_level,
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

func (m *repository) GetOneByID(id int64) ([]*model.Level, error) {
	var (
		list_data = make([]*model.Level, 0)
	)

	query := `SELECT 
	id_level, 
	nama_level
	FROM level  
	WHERE id_level = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Level
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_level,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

// func (m *repository) GetAllByID(id int64) (*model.Level, error) {
// 	query := `SELECT
// 	id,
// 	COALESCE(kode_seksi, ''),
// 	COALESCE(nama_seksi, ''),
// 	id_parent_subdirektorat
// 	FROM level
// 	WHERE id = ?`

// 	data := &model.Level{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Kodesubdirektorat,
// 		&data.Namasubdirektorat,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Level, int, error) {
	var (
		list = make([]*model.Level, 0)
	)

	query := `SELECT nama_level FROM level`

	if dqp.Search != "" {
		query += ` WHERE MATCH(nama_level) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Level
		)

		if err := rows.Scan(
			&data.ID,
			&data.Nama_level,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM level WHERE id_level = ?`

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
