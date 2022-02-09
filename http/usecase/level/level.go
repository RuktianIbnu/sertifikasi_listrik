package level

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rl "sertifikasi_listrik/pkg/repository/level"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Level) error
	GetOneByID(id int64) ([]*model.Level, error)
	UpdateOneByID(data *model.Level) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Level, int, error)
}

type usecase struct {
	levelRepo rl.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rl.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Level) error {
	lastID, err := m.levelRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Level) (int64, error) {

	rowsAffected, err := m.levelRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Level, error) {
	return m.levelRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Level, int, error) {
	return m.levelRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.levelRepo.DeleteOneByID(id)
}
