package pembayaran

import (
	"sertifikasi_listrik/pkg/helper"
	"sertifikasi_listrik/pkg/model"

	"github.com/jmoiron/sqlx"
)

// Repository ...
type Repository interface {
	Create(data *model.Pembayaran) (int64, error)
	GetAllByID(id int64) ([]*model.Pembayaran, error)
	GetOneByID(id int64) (*model.Pembayaran, error)
	UpdateOneByID(data *model.Pembayaran) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error)
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
	if err := m.DB.QueryRow("SELECT COUNT(*) FROM pembayaran").Scan(&totalEntries); err != nil {
		return -1
	}

	return totalEntries
}

func (m *repository) Create(data *model.Pembayaran) (int64, error) {
	query := `INSERT INTO pembayaran(
		id_tagihan, id_pelanggan, tanggal_pembayaran, bulan_bayar, biaya_admin, total_bayar, id_user) VALUES(?, ?, ?, ?, ?, ?, ?)`

	res, err := m.DB.Exec(query,
		&data.IDTagihan,
		&data.IDPelanggan,
		&data.Tanggal_pembayaran,
		&data.Bulan_bayar,
		&data.Biaya_admin,
		&data.Total_bayar,
		&data.IDUser,
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

func (m *repository) UpdateOneByID(data *model.Pembayaran) (int64, error) {
	query := `UPDATE pembayaran set
	id_tagihan = ?, id_pelanggan = ?, tanggal_pembayaran = ?, bulan_bayar = ?, biaya_admin = ?, total_bayar = ?, id_user = ?
	WHERE id_pembayaran = ?`

	res, err := m.DB.Exec(query,
		data.IDTagihan,
		data.IDPelanggan,
		data.Tanggal_pembayaran,
		data.Bulan_bayar,
		data.Biaya_admin,
		data.Total_bayar,
		data.IDUser,
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

func (m *repository) GetAllByID(id int64) ([]*model.Pembayaran, error) {
	var (
		list_data = make([]*model.Pembayaran, 0)
	)

	query := `SELECT 
	id_pembayaran, id_tagihan, id_pelanggan, tanggal_pembayaran, bulan_bayar, biaya_admin, total_bayar, id_user
	FROM pembayaran  
	WHERE id_pembayaran = ?`

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data model.Pembayaran
		)

		if err := rows.Scan(
			&data.ID,
			&data.IDTagihan,
			&data.IDPelanggan,
			&data.Tanggal_pembayaran,
			&data.Bulan_bayar,
			&data.Biaya_admin,
			&data.Total_bayar,
			&data.IDUser,
		); err != nil {
			return nil, err
		}

		list_data = append(list_data, &data)
	}

	return list_data, nil
}

func (m *repository) GetOneByID(id int64) (*model.Pembayaran, error) {
	query := `SELECT 
	id_pembayaran, id_tagihan, id_pelanggan, tanggal_pembayaran, bulan_bayar, biaya_admin, total_bayar, id_user
	FROM pembayaran  
	WHERE id_pembayaran = ?`

	data := &model.Pembayaran{}

	if err := m.DB.QueryRow(query, id).Scan(
		&data.ID,
		&data.IDTagihan,
		&data.IDPelanggan,
		&data.Tanggal_pembayaran,
		&data.Bulan_bayar,
		&data.Biaya_admin,
		&data.Total_bayar,
		&data.IDUser,
	); err != nil {
		return nil, err
	}

	return data, nil
}

func (m *repository) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error) {
	var (
		list = make([]*model.Pembayaran, 0)
	)

	query := `SELECT 
	a.id_pembayaran, 
	a.id_tagihan, 
	a.id_pelanggan, 
	a.tanggal_pembayaran, 
	a.bulan_bayar, 
	a.biaya_admin, 
	a.total_bayar, 
	a.id_user,
	b.status,
	d.nama_pelanggan
	FROM pembayaran as a
	LEFT JOIN tagihan as b on b.id_tagihan = a.id_tagihan
	LEFT JOIN penggunaan as c on c.id_penggunaan = b.id_penggunaan
	LEFT JOIN pelanggan as d on d.id_pelanggan = c.id_pelanggan`

	if dqp.Search != "" {
		query += ` WHERE MATCH(tanggal_pembayaran, bulan_bayar) AGAINST(:search IN NATURAL LANGUAGE MODE)`
	}
	query += ` LIMIT :limit OFFSET :offset`

	rows, err := m.DB.NamedQuery(m.DB.Rebind(query), dqp.Params)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			data          model.Pembayaran
			dataPelanggan model.Pelanggan
			dataTagihan   model.Tagihan
		)

		if err := rows.Scan(
			&data.ID,
			&data.IDTagihan,
			&data.IDPelanggan,
			&data.Tanggal_pembayaran,
			&data.Bulan_bayar,
			&data.Biaya_admin,
			&data.Total_bayar,
			&data.IDUser,
			&dataTagihan.Status,
			&dataPelanggan.Nama_pelanggan,
		); err != nil {
			return nil, -1, err
		}

		data.TagihanDetail = &model.Tagihan{
			Status: dataTagihan.Status,
		}
		data.PelangganDetail = &model.Pelanggan{
			Nama_pelanggan: dataPelanggan.Nama_pelanggan,
		}
		list = append(list, &data)
	}

	return list, m.getTotalCount(), nil
}

func (m *repository) DeleteOneByID(id int64) (int64, error) {
	query := `DELETE FROM pembayaran WHERE id_pembayaran = ?`

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
