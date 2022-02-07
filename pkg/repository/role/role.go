package role

import (
	"epiket-api/pkg/helper"
	"epiket-api/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Roles) (int64, error)
	GetOneByID(id int64) ([]*model.Roles, error)
	// GetAllByID(id int64) (*model.Roles, error)
	UpdateOneByID(data *model.Roles) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Roles, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM ms_roles").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Roles) (int64, error) {
	query := `INSERT INTO ms_roles(
		name_role) VALUES(?)`

	res, err := m.DB.Exec(query,
		&data.Namarole,
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

func (m *repository) UpdateOneByID(data *model.Roles) (int64, error) {
	query := `UPDATE ms_roles set name_role = ?
	WHERE id = ?`

	res, err := m.DB.Exec(query,
		&data.Namarole,
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

func (m *repository) GetOneByID(id int64) ([]*model.Roles, error) {
	var (
		list_data = make([]*model.Roles, 0)
	)

	query := `SELECT 
	id, 
	name_role
	FROM ms_roles  
	WHERE id = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Roles
		)

		if err := rows.Scan(
			&data.ID,
			&data.Namarole,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

// func (m *repository) GetAllByID(id int64) (*model.Roles, error) {
// 	query := `SELECT
// 	id,
// 	COALESCE(kode_seksi, ''),
// 	COALESCE(nama_seksi, ''),
// 	id_parent_subdirektorat
// 	FROM ms_roles
// 	WHERE id = ?`

// 	data := &model.Roles{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Kodesubdirektorat,
// 		&data.Namasubdirektorat,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Roles, int, error) {
	var (
		list = make([]*model.Roles, 0)
	)

	query := `SELECT id, COALESCE(name_role, '') FROM ms_roles`

	if dqp.Search != "" {
		query += ` WHERE MATCH(name_role) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Roles
		)

		if err := rows.Scan(
			&data.ID,
			&data.Namarole,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM ms_roles WHERE id = ?`

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
