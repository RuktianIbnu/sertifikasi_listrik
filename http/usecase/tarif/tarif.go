package tarif

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rt "sertifikasi_listrik/pkg/repository/tarif"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Tarif) (int64, error)
	GetOneByID(id int64) (*model.Tarif, error)
	UpdateOneByID(data *model.Tarif) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Tarif, int, error)
}

type usecase struct {
	tarifRepo rt.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rt.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Tarif) (int64, error) {
	return m.tarifRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.Tarif) (int64, error) {

	rowsAffected, err := m.tarifRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.Tarif, error) {

	data_seksi, err := m.tarifRepo.GetAllByID(id)
	if err != nil {
		return nil, err
	}

	return data_seksi, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Tarif, int, error) {
	return m.tarifRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.tarifRepo.DeleteOneByID(id)
}
