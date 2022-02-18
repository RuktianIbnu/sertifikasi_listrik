package tagihan

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rplgn "sertifikasi_listrik/pkg/repository/pelanggan"
	rpgn "sertifikasi_listrik/pkg/repository/penggunaan"
	rt "sertifikasi_listrik/pkg/repository/tagihan"
	rtrf "sertifikasi_listrik/pkg/repository/tarif"
)

// Usecase ...
type Usecase interface {
	Create(data *model.Tagihan) (int64, error)
	GetOneByID(id int64) (*model.Tagihan, error)
	UpdateOneByID(data *model.Tagihan) (int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error)
}

type usecase struct {
	tagihanRepo    rt.Repository
	penggunaanRepo rpgn.Repository
	pelangganRepo  rplgn.Repository
	tarifRepo      rtrf.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rt.NewRepository(),
		rpgn.NewRepository(),
		rplgn.NewRepository(),
		rtrf.NewRepository(),
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

func (m *usecase) GetOneByID(id int64) (*model.Tagihan, error) {
	data_tagihan, err := m.tagihanRepo.GetOneByID(id)
	if err != nil {
		return nil, err
	}

	penggunaan_detail, err := m.penggunaanRepo.GetOneByID(data_tagihan.IDPenggunaan)
	if err != nil {
		return nil, err
	}

	penggunaan, err := m.penggunaanRepo.GetOneByID(data_tagihan.IDPenggunaan)
	if err != nil {
		return nil, err
	}

	pelanggan_detail, err := m.pelangganRepo.GetOneByID(penggunaan.IDPelanggan)
	if err != nil {
		return nil, err
	}

	detailTarif, err := m.tarifRepo.GetOneByID(int64(pelanggan_detail.IDTarif))
	if err != nil {
		return nil, err
	}

	data_tagihan.PelangganDetail = pelanggan_detail
	data_tagihan.PenggunaanDetail = penggunaan_detail
	data_tagihan.PelangganDetail.TarifDetail = detailTarif

	return data_tagihan, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Tagihan, int, error) {
	return m.tagihanRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.tagihanRepo.DeleteOneByID(id)
}
