package pelanggan

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rp "sertifikasi_listrik/pkg/repository/pelanggan"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Pelanggan) error
	GetOneByID(id int64) (*model.Pelanggan, error)
	UpdateOneByID(data *model.Pelanggan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Pelanggan, int, error)
}

type usecase struct {
	pelangganRepo rp.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Pelanggan) error {
	lastID, err := m.pelangganRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Pelanggan) (int64, error) {

	rowsAffected, err := m.pelangganRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.Pelanggan, error) {
	return m.pelangganRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pelanggan, int, error) {
	return m.pelangganRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.pelangganRepo.DeleteOneByID(id)
}
