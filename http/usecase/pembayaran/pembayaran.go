package pembayaran

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	rplgn "sertifikasi_listrik/pkg/repository/pelanggan"
	rp "sertifikasi_listrik/pkg/repository/pembayaran"
	rpgn "sertifikasi_listrik/pkg/repository/penggunaan"
	rt "sertifikasi_listrik/pkg/repository/tagihan"
	rtrf "sertifikasi_listrik/pkg/repository/tarif"
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
	penggunaanRepo rpgn.Repository
	tarifRepo      rtrf.Repository
}

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		rp.NewRepository(),
		rt.NewRepository(),
		rplgn.NewRepository(),
		ru.NewRepository(),
		rpgn.NewRepository(),
		rtrf.NewRepository(),
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

	//////////////////////get tagihan////////////////////////
	data_tagihan, err := m.tagihanRepo.GetOneByID(data_pembayaran.IDTagihan)
	if err != nil {
		return nil, err
	}

	data, err := m.tagihanRepo.GetOneByID(data_pembayaran.IDTagihan)
	if err != nil {
		return nil, err
	}
	/////////////////////////////////////////////////////////
	///////////////////////////get pelanggan////////////////
	pelanggan_detail, err := m.pelangganRepo.GetOneByID(data.IDPelanggan)
	if err != nil {
		return nil, err
	}
	/////////////////////////////////////////////////////////
	///////////////////////////get penggunaan////////////////
	penggunaan_detail, err := m.penggunaanRepo.GetOneByID(data.IDPenggunaan)
	if err != nil {
		return nil, err
	}
	/////////////////////////////////////////////////////////
	///////////////////////////get tarif/////////////////////
	detailTarif, err := m.tarifRepo.GetOneByID(int64(pelanggan_detail.IDTarif))
	if err != nil {
		return nil, err
	}
	/////////////////////////////////////////////////////////

	pelanggan_detail.TarifDetail = detailTarif
	data_pembayaran.PelangganDetail = pelanggan_detail
	data_pembayaran.TagihanDetail = data_tagihan
	data_pembayaran.PenggunaanDetail = penggunaan_detail

	return data_pembayaran, nil
}

func (m *usecase) GetAll(dqp *model.DefaultQueryParam) ([]*model.Pembayaran, int, error) {
	return m.pembayaranRepo.GetAll(dqp)
}

func (m *usecase) DeleteOneByID(id int64) (int64, error) {
	return m.pembayaranRepo.DeleteOneByID(id)
}
