package pelanggan

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Pelanggan) (int64, error)
	GetOneByID(id int64) ([]*model.Pelanggan, error)
	GetAllByID(id int64) (*model.Pelanggan, error)
	UpdateOneByID(data *model.Pelanggan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Pelanggan, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM pelanggan").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Pelanggan) (int64, error) {
	query := `INSERT INTO pelanggan(
		username, password, nomor_kwh, nama_pelanggan, alamat, id_tarif
	) VALUES(?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.Username,
		&data.Password,
		&data.Nomor_kwh,
		&data.Nama_pelanggan,
		&data.Alamat,
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

func (m *repository) UpdateOneByID(data *model.Pelanggan) (int64, error) {
	query := `UPDATE pelanggan set
	username = ?, password = ?, nomor_kwh = ?, nama_pelanggan = ?, alamat = ?, id_tarif = ?
	WHERE id_pelanggan = ?`

	res, err := m.DB.Exec(query,
		data.Username,
		data.Password,
		data.Nomor_kwh,
		data.Nama_pelanggan,
		data.Alamat,
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

func (m *repository) GetAllByID(id int64) (*model.Pelanggan, error) {
	query := `SELECT 
	id_pelanggan, 
	username, 
	password, 
	nomor_kwh, 
	nama_pelanggan, 
	alamat, 
	id_tarif
	FROM pelanggan  
	WHERE id_pelanggan = ?`

	data := &model.Pelanggan{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.Username,
		&data.Password,
		&data.Nomor_kwh,
		&data.Nama_pelanggan,
		&data.Alamat,
		&data.Tarif,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetOneByID(id int64) ([]*model.Pelanggan, error) {
	var (
		list_data = make([]*model.Pelanggan, 0)
	)

	query := `SELECT 
	id_pelanggan, 
	username, 
	password, 
	nomor_kwh, 
	nama_pelanggan, 
	alamat, 
	id_tarif
	FROM pelanggan  
	WHERE id_pelanggan = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Pelanggan
		)

		if err := m.DB.QueryRow(query, id).Scan(
			&data.ID,
			&data.Username,
			&data.Password,
			&data.Nomor_kwh,
			&data.Nama_pelanggan,
			&data.Alamat,
			&data.Tarif,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pelanggan, int, error) {
	var (
		list = make([]*model.Pelanggan, 0)
	)

	query := `SELECT 
	id_pelanggan, 
	username, 
	password, 
	nomor_kwh, 
	nama_pelanggan, 
	alamat, 
	id_tarif
	FROM pelanggan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(username, nomor_kwh) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Pelanggan
		)

		if err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Password,
			&data.Nomor_kwh,
			&data.Nama_pelanggan,
			&data.Alamat,
			&data.Tarif,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM pelanggan WHERE id_pelanggan = ?`

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
