package pelanggan

import (
	"errors"
	"net/http"
	uk "sertifikasi_listrik/http/usecase/pelanggan"
	qry "sertifikasi_listrik/pkg/helper/query"
	resp "sertifikasi_listrik/pkg/helper/response"
	"sertifikasi_listrik/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler interface {
	Create(c *gin.Context)
	GetOneByID(c *gin.Context)
	UpdateOneByID(c *gin.Context)
	DeleteOneByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type handler struct {
	pelangganUc uk.Usecase
}

// NewHandler ...
func NewHandler() Handler {
	return &handler{
		uk.NewUsecase(),
	}
}

func (m *handler) Create(c *gin.Context) {
	var (
		data model.Pelanggan
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	// lastID, err := m.pelangganUc.Create(&data)
	if err := m.pelangganUc.Create(&data); err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) UpdateOneByID(c *gin.Context) {
	var (
		data   model.Pelanggan
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(resp.Format(http.StatusBadRequest, err))
		return
	}

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	data.ID = ids

	_, err := m.pelangganUc.UpdateOneByID(&data)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) GetOneByID(c *gin.Context) {
	var (
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	data, err := m.pelangganUc.GetOneByID(ids)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, data))
}

func (m *handler) GetAll(c *gin.Context) {
	var (
		dq = qry.Q{
			Ctx:     c,
			Sorting: []string{"kode", "tipe", "deskripsi"},
		}
	)

	dqp, err := dq.DefaultQueryParam()
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}
	list, totalEntries, err := m.pelangganUc.GetAll(dqp)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil, list, totalEntries, dqp.Page, dqp.Limit))
}

func (m *handler) DeleteOneByID(c *gin.Context) {
	var (
		ids, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	)

	if ids <= 0 {
		c.JSON(resp.Format(http.StatusBadRequest, errors.New("Provide a valid ID")))
		return
	}

	_, err := m.pelangganUc.DeleteOneByID(ids)
	if err != nil {
		c.JSON(resp.Format(http.StatusInternalServerError, err))
		return
	}

	c.JSON(resp.Format(http.StatusOK, nil))
}
