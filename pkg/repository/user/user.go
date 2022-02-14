package user

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	CheckPelangganIsExist(username string) (exist bool)
	Create(data *model.User) (int64, error)
	UpdateOneByID(data *model.User) (int64, error)
	GetUserMetadataByIdUser(username string) (*model.User, error)
	GetOneByID(id int64) ([]*model.User, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.User, int, error)
	DeleteOneByID(id int64) (int64, error)
	getTotalCount() (totalEntries int)

	Register(username string, password string, nama_admin string, id_level int64) (int64, error)
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

func (m *repository) Register(username string, password string, nama_admin string, id_level int64) (int64, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return -1, err
	}

	res, err := tx.Exec(`INSERT INTO user(username, password, nama_admin, id_level) VALUES(?, ?, ?, ?)`, username, password, nama_admin, id_level)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	lastIDUser, _ := res.LastInsertId()

	return lastIDUser, tx.Commit()
}

func (m *repository) getTotalCount() (totalEntries int) {
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM user").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.User) (int64, error) {
	query := `INSERT INTO user(
		username, password, nama_admin, id_level) VALUES(?,?,?,?)`

	res, err := m.DB.Exec(query,
		&data.Username,
		&data.Password,
		&data.Nama_admin,
		&data.IDLevel,
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

func (m *repository) CheckPelangganIsExist(username string) (exist bool) {
	query := `SELECT 
	coalesce(username, '') 
	FROM ms_users 
	WHERE username = ?`

	var e string

	if err := m.DB.QueryRow(query, username).Scan(
		&e,
	); err != nil {
		return false
	}

	if e != "" {
		exist = true
	}

	return
}

func (m *repository) UpdateOneByID(data *model.User) (int64, error) {
	query := `UPDATE user set  username=?, password=?, nama_admin=?, id_level=?
	WHERE id_pegawai = ?`

	res, err := m.DB.Exec(query,
		&data.Username,
		&data.Password,
		&data.Nama_admin,
		&data.IDLevel,
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

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM tarif WHERE id_user = ?`

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

func (m *repository) GetOneByID(id int64) ([]*model.User, error) {
	var (
		list_data = make([]*model.User, 0)
	)

	query := `SELECT
	id_user, username, nama_admin, id_level
	FROM user
	WHERE id_user = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.User
		)

		if err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Nama_admin,
			&data.IDLevel,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.User, int, error) {
	var (
		list = make([]*model.User, 0)
	)

	query := `SELECT id_user, username, nama_admin, id_level FROM user`

	if dqp.Search != "" {
		query += ` WHERE MATCH(username, nama_admin) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.User
		)

		if err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Nama_admin,
			&data.IDLevel,
		); err != nil {
			return nil, -1, err
		}

		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) GetUserMetadataByIdUser(username string) (*model.User, error) {
	query := `SELECT
	id_user,
	username,
	nama_admin,
	id_level
	FROM user 
	WHERE username = ?`

	data := &model.User{}

	if err := m.DB.QueryRow(query, username).Scan(
		&data.ID,
		&data.Username,
		&data.Nama_admin,
		&data.IDLevel,
	); err != nil {
		return nil, err
	}

	return data, nil
}
