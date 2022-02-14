package login

import (
	"net/http"
	gu "sertifikasi_listrik/http/usecase/login"
	resp "sertifikasi_listrik/pkg/helper/response"

	"sertifikasi_listrik/pkg/model"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type handler struct {
	globalUsecase gu.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		gu.NewUsecase(),
	}
}

func (m *handler) Register(c *gin.Context) {
	type register struct {
		Username   string `json:"username"`
		Password   string `json:"password" binding:"required"`
		Nama_admin string `json:"nama_admin" binding:"required"`
		Id_level   int64  `json:"id_level" binding:"required"`
	}

	var (
		registerData register
	)

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	dataResult := &model.User{}

	lastID, err := m.globalUsecase.Register(registerData.Username, registerData.Password, registerData.Nama_admin, registerData.Id_level)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	dataResult.ID = lastID
	c.JSON(resp.Format(http.StatusOK, nil, gin.H{"registered": true}))
}

func (m *handler) Login(c *gin.Context) {
	type login struct {
		Nip      string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}

	var (
		loginData login
	)

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	token, userMetadata, err := m.globalUsecase.Login(loginData.Nip, loginData.Password)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, gin.H{"token": token, "user_metadata": userMetadata}))
}
