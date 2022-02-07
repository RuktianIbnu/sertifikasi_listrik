package seksi

import (
	"epiket-api/pkg/helper"
	"epiket-api/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Seksi) (int64, error)
	GetOneByID(id int64) ([]*model.Seksi, error)
	GetAllByID(id int64) (*model.Seksi, error)
	UpdateOneByID(data *model.Seksi) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Seksi, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_seksi").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Seksi) (int64, error) {
	query := `INSERT INTO ms_seksi(
		kode_seksi, nama_seksi, id_parent_subdirektorat
	) VALUES(?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.Kodeseksi,
		&data.Namaseksi,
		&data.IDsubdirektorat,
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

func (m *repository) UpdateOneByID(data *model.Seksi) (int64, error) {
	query := `UPDATE ms_seksi set
	kode_seksi = ?, nama_seksi = ?, id_parent_subdirektorat = ?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		data.Kodeseksi,
		data.Namaseksi,
		data.IDsubdirektorat,
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

func (m *repository) GetAllByID(id int64) (*model.Seksi, error) {
	query := `SELECT 
	id, 
	COALESCE(kode_seksi, ''), 
	COALESCE(nama_seksi, ''),
	id_parent_subdirektorat
	FROM ms_seksi  
	WHERE id = ?`

	data := &model.Seksi{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Kodeseksi,
		&data.Namaseksi,
		&data.IDsubdirektorat,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetOneByID(id int64) ([]*model.Seksi, error) {
	var (
		list_data = make([]*model.Seksi, 0)
	)

	query := `SELECT 
	id, 
	COALESCE(kode_seksi, ''), 
	COALESCE(nama_seksi, ''),
	id_parent_subdirektorat
	FROM ms_seksi  
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Seksi
		)

		if err := m.DB.QueryRow(query, id).Scan(
			&data.ID,
			&data.Kodeseksi,
			&data.Namaseksi,
			&data.IDsubdirektorat,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Seksi, int, error) {
	var (
		list = make([]*model.Seksi, 0)
	)

	query := `SELECT id, kode_seksi, COALESCE(nama_seksi, ''), id_parent_subdirektorat FROM ms_seksi`

	if dqp.Search != "" {
		query += ` WHERE MATCH(tipe, deskripsi) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Seksi
		)

		if err := rows.Scan(
			&data.ID,
			&data.Kodeseksi,
			&data.Namaseksi,
			&data.IDsubdirektorat,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM ms_seksi WHERE id = ?`

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
