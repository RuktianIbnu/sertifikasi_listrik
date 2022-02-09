package query

import (
	"errors"
	"sertifikasi_listrik/pkg/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Q ...
type Q struct {
	Ctx     *gin.Context
	Sorting []string
}

// DefaultQueryParam ...
func (m *Q) DefaultQueryParam() (res *model.DefaultQueryParam, err error) {
	page, _ := strconv.Atoi(m.Ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(m.Ctx.DefaultQuery("limit", "10"))
	search := m.Ctx.Query("search")
	sortMap := m.Ctx.QueryMap("sort")
	filterMap := make(map[string]interface{}, 0)
	offset := (page - 1) * limit

	for i, v := range sortMap {
		filterMap[i] = v
	}
	filterMap["limit"] = limit
	filterMap["offset"] = offset
	filterMap["search"] = search

	result := &model.DefaultQueryParam{}

	result.Search = search
	result.Page = page
	result.Limit = limit
	result.Sorting = sortMap
	result.Offset = offset
	result.Params = filterMap

	for i := range result.Sorting {
		match := false
		for _, w := range m.Sorting {
			if i == w {
				match = true
				break
			}
		}

		if !match {
			return nil, errors.New("unaccepted parameter sort")
		}
	}

	return result, nil
}
