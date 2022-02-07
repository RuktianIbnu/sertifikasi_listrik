package seksi

import (
	"epiket-api/pkg/model"
	rs "epiket-api/pkg/repository/seksi"

	rsub "epiket-api/pkg/repository/subdirektorat"
	"errors"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Seksi) (int64, error)
	GetOneByID(id int64) (*model.Seksi, error)
	UpdateOneByID(data *model.Seksi) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Seksi, int, error)
}

type usecase struct {
	seksiRepo  rs.Repository
	subditRepo rsub.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rs.NewRepository(),
		rsub.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Seksi) (int64, error) {
	return m.seksiRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.Seksi) (int64, error) {

	rowsAffected, err := m.seksiRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) (*model.Seksi, error) {

	data_seksi, err := m.seksiRepo.GetAllByID(id)
	if err != nil {
		return nil, err
	}

	detail_subdit, err := m.subditRepo.GetOneByID(data_seksi.IDsubdirektorat)
	if err != nil {
		return nil, err
	}

	data_seksi.ParentDetail = detail_subdit

	return data_seksi, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Seksi, int, error) {
	return m.seksiRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.seksiRepo.DeleteOneByID(id)
}
