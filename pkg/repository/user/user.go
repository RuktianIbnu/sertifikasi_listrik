package user

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	//Register(nip string, nama string, Nohp string, IDsubdirektorat int64, IDseksi int64, Levelpengguna int64, Password string) (int64, error)
	CheckPelangganIsExist(username string) (exist bool)
	// Create(data *model.User) (int64, error)
	CheckUserIsActive(username string) (active bool)
	// UpdateOneByID(data *model.User) (rowsAffected int64, err error)
	// UpdatePasswordOneByID(ID int64, newpassword string) (rowsAffected int64, err error)
	GetUserMetadataByIdUser(id_user string) (*model.User, error)
	// GetOneByID(id int64) (*model.User, error)
	// GetOneByIDPegawai(id int64) (*model.User, error)
	// GetAll() ([]*model.User, error)
	// DeleteOneByID(userExist, ids int64, active bool) (int64, error)
	// ResetPasswordByID(userID int64, newpassword string) (int64, error)
	// CheckUserStatus(userID int64) (status string, err error)
	// UpdateStatusActiveOneByID(ID int64, active bool) (rowsAffected int64, err error)
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

// func (m *repository) Register(nip string, nama string, Nohp string, IDsubdirektorat int64, IDseksi int64, Levelpengguna int64, Password string) (int64, error) {
// 	tx, err := m.DB.Begin()
// 	if err != nil {
// 		return -1, err
// 	}

// 	res, err := tx.Exec(`INSERT INTO ms_users(nip, nama, no_hp, id_subdirektorat, id_seksi, level_pengguna, password) VALUES(?, ?, ?,?,? ,?, ?)`, nip, nama, Nohp, IDsubdirektorat, IDseksi, Levelpengguna, Password)
// 	if err != nil {
// 		tx.Rollback()
// 		return -1, err
// 	}
// 	lastIDUser, _ := res.LastInsertId()

// 	return lastIDUser, tx.Commit()
// }

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

// func (m *repository) UpdateStatusActiveOneByID(ID int64, active bool) (rowsAffected int64, err error) {
// 	query := `UPDATE user set  active=?, updated_at=now(), updated_by=?
// 	WHERE id_pegawai = ?`

// 	res, err := m.DB.Exec(query,
// 		active,
// 		ID,
// 		ID,
// 	)

// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

// func (m *repository) UpdatePasswordOneByID(ID int64, newpassword string) (rowsAffected int64, err error) {
// 	query := `UPDATE user set  password=?, updated_at=now(), updated_by=?
// 	WHERE id_pegawai = ?`
// 	//fmt.Println(newpassword,ID,ID,)
// 	res, err := m.DB.Exec(query,
// 		newpassword,
// 		ID,
// 		ID,
// 	)

// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

// func (m *repository) UpdateOneByID(data *model.User) (rowsAffected int64, err error) {
// 	query := `UPDATE user set  active=?, role=?, updated_at=now(), updated_by=?
// 	WHERE id_pegawai = ?`

// 	res, err := m.DB.Exec(query,
// 		&data.Active,
// 		&data.Role,
// 		&data.ActionBy.CreatedBy,
// 		&data.IDPegawai,
// 	)

// 	//fmt.Println("user", err)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ = res.RowsAffected()

// 	return
// }

// func (m *repository) DeleteOneByID(userExist, ids int64, active bool) (int64, error) {
// 	query := `UPDATE user set
// 	 active=?,  deleted_at=now(), deleted_by=?
// 	WHERE id_pegawai = ?`

// 	res, err := m.DB.Exec(query,
// 		active,
// 		userExist,
// 		ids,
// 	)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ := res.RowsAffected()

// 	return rowsAffected, nil
// // }

// func (m *repository) CheckUserStatus(userID int64) (status string, err error) {
// 	var (
// 		statusData string
// 	)
// 	query := `SELECT
// 		 coalesce(role,'')
// 		FROM user WHERE id=? `

// 	if err := m.DB.QueryRow(query, userID).Scan(&statusData); err != nil {
// 		return "", err
// 	}
// 	//fmt.Println(statusData)
// 	return statusData, nil
// }

// func (m *repository) GetOneByID(id int64) (*model.User, error) {
// 	query := `SELECT
// 	id, email, active, id_pegawai, role
// 	FROM user
// 	WHERE id = ?`

// 	data := &model.User{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Email,
// 		&data.Active,
// 		&data.IDPegawai,
// 		&data.Role,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func (m *repository) GetOneByIDPegawai(id int64) (*model.User, error) {
// 	query := `SELECT
// 	id, email, active, id_pegawai, role, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
// 	FROM user
// 	WHERE id_pegawai = ?`

// 	data := &model.User{}

// 	if err := m.DB.QueryRow(query, id).Scan(
// 		&data.ID,
// 		&data.Email,
// 		&data.Active,
// 		&data.IDPegawai,
// 		&data.Role,
// 		&data.ActionBy.CreatedAt,
// 		&data.ActionBy.CreatedBy,
// 		&data.ActionBy.UpdatedAt,
// 		&data.ActionBy.UpdatedBy,
// 		&data.ActionBy.DeletedAt,
// 		&data.ActionBy.DeletedBy,
// 	); err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func (m *repository) GetAll() ([]*model.User, error) {
// 	query := `SELECT
// 	id, email, active, id_pegawai, role, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
// 	FROM user`

// 	var (
// 		list = make([]*model.User, 0)
// 	)

// 	rows, err := m.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var (
// 			data model.User
// 		)

// 		if err := rows.Scan(
// 			&data.ID,
// 			&data.Email,
// 			&data.Active,
// 			&data.IDPegawai,
// 			&data.Role,
// 			&data.ActionBy.CreatedAt,
// 			&data.ActionBy.CreatedBy,
// 			&data.ActionBy.UpdatedAt,
// 			&data.ActionBy.UpdatedBy,
// 			&data.ActionBy.DeletedAt,
// 			&data.ActionBy.DeletedBy,
// 		); err != nil {
// 			return nil, err
// 		}

// 		list = append(list, &data)
// 	}

// 	return list, nil
// }

func (m *repository) CheckUserIsActive(nip string) (active bool) {
	query := `SELECT
	coalesce(nip, '')
	FROM ms_users
	WHERE nip = ? AND aktif = true`

	var e string

	if err := m.DB.QueryRow(query, nip).Scan(
		&e,
	); err != nil {
		return false
	}

	if e != "" {
		active = true
	}

	return
}

func (m *repository) GetUserMetadataByIdUser(id_user string) (*model.User, error) {
	query := `SELECT
	id_user,
	username,
	nama_admin,
	id_level
	FROM user 
	WHERE id_user = ?`

	data := &model.User{}

	if err := m.DB.QueryRow(query, id_user).Scan(
		&data.ID,
		&data.Username,
		&data.Nama_admin,
		&data.IDLevel,
	); err != nil {
		return nil, err
	}

	return data, nil
}

// func (m *repository) ResetPasswordByID(userID int64, newpassword string) (int64, error) {
// 	query := `UPDATE user set
// 	password=?, updated_at=now(), updated_by=?
// 	WHERE id_pegawai = ?`
// 	//fmt.Println(newpassword, userID)
// 	res, err := m.DB.Exec(query,
// 		newpassword,
// 		userID,
// 		userID,
// 	)

// 	//fmt.Println("user", err)
// 	if err != nil {
// 		return -1, err
// 	}

// 	rowsAffected, _ := res.RowsAffected()

// 	return rowsAffected, nil
// }
