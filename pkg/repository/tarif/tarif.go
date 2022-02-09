package tarif

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Tarif) (int64, error)
	GetOneByID(id int64) ([]*model.Tarif, error)
	// GetAllByID(id int64) (*model.Tarif, error)
	UpdateOneByID(data *model.Tarif) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Tarif, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM tarif").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Tarif) (int64, error) {
	query := `INSERT INTO tarif(
		daya, tarifperkwh) VALUES(?)`

	res, err := m.DB.Exec(query,
		&data.Daya,
		&data.Tarif,
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

func (m *repository) UpdateOneByID(data *model.Tarif) (int64, error) {
	query := `UPDATE tarif set daya = ?, tarifperkwh = ?
	WHERE id_tarif = ?`

	res, err := m.DB.Exec(query,
		data.Daya,
		data.Tarif,
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

func (m *repository) GetOneByID(id int64) ([]*model.Tarif, error) {
	var (
		list_data = make([]*model.Tarif, 0)
	)

	query := `SELECT 
	id_tarif, 
	daya,
	tarifperkwh
	FROM tarif  
	WHERE id_tarif = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Tarif
		)

		if err := rows.Scan(
			&data.ID,
			&data.Daya,
			&data.Tarif,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

// func (m *repository) GetAllByID(id int64) (*model.Tarif, error) {
// 	query := `SELECT
// 	id,
// 	COALESCE(kode_seksi, ''),
// 	COALESCE(nama_seksi, ''),
// 	id_parent_subdirektorat
// 	FROM tarif
// 	WHERE id = ?`

// 	data := &model.Tarif{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Kodesubdirektorat,
// 		&data.Namasubdirektorat,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Tarif, int, error) {
	var (
		list = make([]*model.Tarif, 0)
	)

	query := `SELECT id_tarif, daya, tarifperkwh FROM tarif`

	if dqp.Search != "" {
		query += ` WHERE MATCH(id_tarif, daya) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Tarif
		)

		if err := rows.Scan(
			&data.ID,
			&data.Daya,
			&data.Tarif,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM tarif WHERE id_tarif = ?`

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
