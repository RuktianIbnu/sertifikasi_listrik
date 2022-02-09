package penggunaan

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rp "sertifikasi_listrik/pkg/repository/penggunaan"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Penggunaan) error
	GetOneByID(id int64) ([]*model.Penggunaan, error)
	UpdateOneByID(data *model.Penggunaan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error)
}

type usecase struct {
	penggunaanRepo rp.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
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

func (m *usecase) UpdateOneByID(data *model.Penggunaan) (int64, error) {

	rowsAffected, err := m.penggunaanRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Penggunaan, error) {
	return m.penggunaanRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Penggunaan, int, error) {
	return m.penggunaanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.penggunaanRepo.DeleteOneByID(id)
}
