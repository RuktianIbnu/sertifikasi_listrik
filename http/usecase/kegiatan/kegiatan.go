package kegiatan

import (
	"epiket-api/pkg/model"
	rk "epiket-api/pkg/repository/kegiatan"
	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Kegiatan) error
	GetOneByID(id int64) ([]*model.Kegiatan, error)
	UpdateOneByID(data *model.Kegiatan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Kegiatan, int, error)
}

type usecase struct {
	kegiatanRepo rk.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rk.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Kegiatan) error {
	lastID, err := m.kegiatanRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Kegiatan) (int64, error) {

	rowsAffected, err := m.kegiatanRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Kegiatan, error) {
	return m.kegiatanRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Kegiatan, int, error) {
	return m.kegiatanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.kegiatanRepo.DeleteOneByID(id)
}
