package pembayaran

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rp "sertifikasi_listrik/pkg/repository/pembayaran"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Pembayaran) error
	GetOneByID(id int64) ([]*model.Pembayaran, error)
	UpdateOneByID(data *model.Pembayaran) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error)
}

type usecase struct {
	pembayaranRepo rp.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
	}
}

func (m *usecase) Create(data *model.Pembayaran) error {
	lastID, err := m.pembayaranRepo.Create(data)
	if err != nil {
		return err
	}
	data.ID = lastID

	return nil
}

func (m *usecase) UpdateOneByID(data *model.Pembayaran) (int64, error) {

	rowsAffected, err := m.pembayaranRepo.UpdateOneByID(data)

	if rowsAffected <= 0 {
		return rowsAffected, errors.New("no rows affected or data not found")
	}

	return rowsAffected, err
}

func (m *usecase) GetOneByID(id int64) ([]*model.Pembayaran, error) {
	return m.pembayaranRepo.GetOneByID(id)
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error) {
	return m.pembayaranRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.pembayaranRepo.DeleteOneByID(id)
}
