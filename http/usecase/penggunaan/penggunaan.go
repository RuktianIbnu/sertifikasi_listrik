package penggunaan

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rplgn "sertifikasi_listrik/pkg/repository/pelanggan"
	rp "sertifikasi_listrik/pkg/repository/penggunaan"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Penggunaan) error
	GetOneByID(id int64) (*model.Penggunaan, error)
	UpdateOneByID(data *model.Penggunaan) (int64, error)
	UpdateStatus(id int64) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error)
}

type usecase struct {
	penggunaanRepo rp.Repository
	pelangganRepo  rplgn.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
		rplgn.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Penggunaan) error {
	lastID, err := m.penggunaanRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateStatus(id int64) (int64, error) {

	rowsAffected, err := m.penggunaanRepo.UpdateStatus(id)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) UpdateOneByID(data *model.Penggunaan) (int64, error) {

	rowsAffected, err := m.penggunaanRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.Penggunaan, error) {
	data_penggunaan, err := m.penggunaanRepo.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	pelanggan_detail, err := m.pelangganRepo.GetOneByID(data_penggunaan.IDPelanggan)
	if err != nil {
		return nil, err
	}

	data_penggunaan.PelangganDetail = pelanggan_detail

	return data_penggunaan, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error) {
	return m.penggunaanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.penggunaanRepo.DeleteOneByID(id)
}
