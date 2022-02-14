package login

import (
	"errors"
	"fmt"
	"sertifikasi_listrik/pkg/helper/bcrypt"
	"sertifikasi_listrik/pkg/helper/jwt"
	"sertifikasi_listrik/pkg/model"
	ur "sertifikasi_listrik/pkg/repository/user"
)

// Usecase ...
type Usecase interface {
	Login(nip, password string) (string, *model.User, error)
	Register(username, password, nama_admin string, IDLevel int64) (int64, error)
}

type usecase struct {
	userRepo ur.Repository
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
	}
}

func (m *usecase) Register(username, password, nama_admin string, IDLevel int64) (int64, error) {
	if username := m.userRepo.CheckPelangganIsExist(username); username {
		return 500, errors.New("username already registered")
	}

	hashedPwd, err := bcrypt.Hash(password)
	if err != nil {
		return 500, errors.New("gagal encrypt password")
	}

	lastIDUser, err := m.userRepo.Register(username, hashedPwd, nama_admin, IDLevel)
	if err != nil {
		return 500, errors.New("gagal registrasi")
	}

	return lastIDUser, err
}

func (m *usecase) Login(username, password string) (string, *model.User, error) {
	userMetadata, err := m.userRepo.GetUserMetadataByIdUser(username)
	if err != nil {
		return "", nil, errors.New("Username not registered")
	}

	if !bcrypt.Compare(password, userMetadata.Password) {
		return "", nil, errors.New("incorrect nip or password")
	}

	token, err := jwt.CreateToken(userMetadata.ID, userMetadata.Username, userMetadata.Nama_admin, userMetadata.IDLevel)
	if err != nil {
		return "", nil, fmt.Errorf("failed generate temporary token %s", err.Error())
	}

	//set empty password
	userMetadata.Password = ""

	return token, userMetadata, nil
}
