package role

import (
	"epiket-api/pkg/model"
	rr "epiket-api/pkg/repository/role"
	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Roles) error
	GetOneByID(id int64) ([]*model.Roles, error)
	UpdateOneByID(data *model.Roles) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Roles, int, error)
}

type usecase struct {
	roleRepo rr.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rr.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Roles) error {
	lastID, err := m.roleRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Roles) (int64, error) {

	rowsAffected, err := m.roleRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Roles, error) {
	return m.roleRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Roles, int, error) {
	return m.roleRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.roleRepo.DeleteOneByID(id)
}
