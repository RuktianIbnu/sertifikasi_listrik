package login

import (
	"epiket-api/pkg/helper/bcrypt"
	"epiket-api/pkg/helper/jwt"
	"epiket-api/pkg/model"
	useksi "epiket-api/pkg/repository/seksi"
	usubdit "epiket-api/pkg/repository/subdirektorat"
	ur "epiket-api/pkg/repository/user"
	"errors"
	"fmt"
)

// Usecase ...
type Usecase interface {
	Login(nip, password string) (string, *model.User, error)
	Register(nip string, nama string, Nohp string, IDsubdirektorat int64, IDseksi int64, Levelpengguna int64, Password string) (int64, error)
}

type usecase struct {
	userRepo   ur.Repository
	seksiRepo  useksi.Repository
	subditRepo usubdit.Repository
}

const (
	// ActionLogin ...
	ActionLogin = "login"

	// RoleRespondent ...
	RoleRespondent = "respondent"
)

// NewUsecase ...
func NewUsecase() Usecase {
	return &usecase{
		ur.NewRepository(),
		useksi.NewRepository(),
		usubdit.NewRepository(),
	}
}

func (m *usecase) Register(nip string, nama string, Nohp string, IDsubdirektorat int64, IDseksi int64, Levelpengguna int64, Password string) (int64, error) {
	if nipExist := m.userRepo.CheckNIPExist(nip); nipExist {
		return 500, errors.New("nip already registered")
	}

	hashedPwd, err := bcrypt.Hash(Password)
	if err != nil {
		return 500, errors.New("gagal crypt password")
	}

	lastIDUser, err := m.userRepo.Register(nip, nama, Nohp, IDsubdirektorat, IDseksi, Levelpengguna, hashedPwd)
	if err != nil {
		return 500, errors.New("gagal registrasi")
	}

	return lastIDUser, err
}

func (m *usecase) Login(nip, password string) (string, *model.User, error) {
	userMetadata, err := m.userRepo.GetUserMetadataByNip(nip)
	if err != nil {
		return "", nil, errors.New("nip not registered")
	}

	userIsActive := m.userRepo.CheckUserIsActive(nip)
	if !userIsActive {
		return "", nil, errors.New("please activate your account")
	}

	if !bcrypt.Compare(password, userMetadata.Password) {
		return "", nil, errors.New("incorrect nip or password")
	}

	token, err := jwt.CreateToken(userMetadata.ID, userMetadata.NIP, userMetadata.Nama, userMetadata.Nohp, userMetadata.Photo,
		userMetadata.IDsubdirektorat, userMetadata.IDseksi, userMetadata.Levelpengguna, userMetadata.Aktif)
	if err != nil {
		return "", nil, fmt.Errorf("failed generate temporary token %s", err.Error())
	}

	detail_subdit, err := m.subditRepo.GetOneByID(userMetadata.IDsubdirektorat)
	if err != nil {
		return "", nil, err
	}

	detail_seksi, err := m.seksiRepo.GetOneByID(userMetadata.IDseksi)
	if err != nil {
		return "", nil, err
	}

	//set empty password
	userMetadata.Password = ""

	//set subdit n seksi detail
	userMetadata.SubdirektoratDetail = detail_subdit
	userMetadata.SeksiDetail = detail_seksi

	return token, userMetadata, nil
}
