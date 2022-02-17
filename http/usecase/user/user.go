package user

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	ru "sertifikasi_listrik/pkg/repository/user"
)

// Usecase ...
type Usecase interface {
	Create(data *model.User) error
	GetOneByID(id int64) (*model.User, error)
	UpdateOneByID(data *model.User) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.User, int, error)
}

type usecase struct {
	userRepo ru.Repository
	// userRepo rsub.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ru.NewRepository(),
		// rsub.NewRepository(),
	}
}

func (m *usecase) Create(data *model.User) error {
	lastID, err := m.userRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.User) (int64, error) {

	rowsAffected, err := m.userRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.User, error) {
	return m.userRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.User, int, error) {
	return m.userRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.userRepo.DeleteOneByID(id)
}
