package tagihan

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rt "sertifikasi_listrik/pkg/repository/tagihan"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Tagihan) (int64, error)
	GetOneByID(id int64) ([]*model.Tagihan, error)
	UpdateOneByID(data *model.Tagihan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error)
}

type usecase struct {
	tagihanRepo rt.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rt.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Tagihan) (int64, error) {
	return m.tagihanRepo.Create(data)
}

func (m *usecase) UpdateOneByID(data *model.Tagihan) (int64, error) {

	rowsAffected, err := m.tagihanRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Tagihan, error) {
	return m.tagihanRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error) {
	return m.tagihanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.tagihanRepo.DeleteOneByID(id)
}
