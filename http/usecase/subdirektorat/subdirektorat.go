package subdirektorat

import (
	"epiket-api/pkg/model"
	rsub "epiket-api/pkg/repository/subdirektorat"
	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Subdirektorat) error
	GetOneByID(id int64) ([]*model.Subdirektorat, error)
	UpdateOneByID(data *model.Subdirektorat) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Subdirektorat, int, error)
}

type usecase struct {
	subditRepo rsub.Repository
	// SubditRepo rsub.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rsub.NewRepository(),
		// rsub.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Subdirektorat) error {
	lastID, err := m.subditRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Subdirektorat) (int64, error) {

	rowsAffected, err := m.subditRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Subdirektorat, error) {
	return m.subditRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Subdirektorat, int, error) {
	return m.subditRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.subditRepo.DeleteOneByID(id)
}
