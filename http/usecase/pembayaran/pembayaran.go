package pembayaran

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rplgn "sertifikasi_listrik/pkg/repository/pelanggan"
	rp "sertifikasi_listrik/pkg/repository/pembayaran"
	rt "sertifikasi_listrik/pkg/repository/tagihan"
	ru "sertifikasi_listrik/pkg/repository/user"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Pembayaran) error
	GetOneByID(id int64) (*model.Pembayaran, error)
	UpdateOneByID(data *model.Pembayaran) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error)
}

type usecase struct {
	pembayaranRepo rp.Repository
	tagihanRepo    rt.Repository
	pelangganRepo  rplgn.Repository
	userRepo       ru.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
		rt.NewRepository(),
		rplgn.NewRepository(),
		ru.NewRepository(),
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

func (m *usecase) GetOneByID(id int64) (*model.Pembayaran, error) {
	data_pembayaran, err := m.pembayaranRepo.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	data_tagihan, err := m.tagihanRepo.GetOneByID(data_pembayaran.IDTagihan)
	if err != nil {
		return nil, err
	}

	data, err := m.tagihanRepo.GetOneByID(data_pembayaran.IDTagihan)
	if err != nil {
		return nil, err
	}

	pelanggan_detail, err := m.pelangganRepo.GetOneByID(data.IDPelanggan)
	if err != nil {
		return nil, err
	}

	data_pembayaran.PelangganDetail = pelanggan_detail
	data_pembayaran.TagihanDetail = data_tagihan

	return data_pembayaran, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error) {
	return m.pembayaranRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.pembayaranRepo.DeleteOneByID(id)
}
