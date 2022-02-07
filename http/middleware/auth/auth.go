package auth

import (
	"errors"
	"net/http"
	"strings"

	"epiket-api/pkg/helper/jwt"
	resp "epiket-api/pkg/helper/response"

	"github.com/gin-gonic/gin"
)

// Middleware ...
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authorization = strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", -1)
		)

		if authorization == "" {
			c.JSON(resp.Format(http.StatusBadRequest, errors.New("please provide authorization")))
			c.Abort()
			return
		}

		if err := jwt.TokenValid(authorization); err != nil {
			c.JSON(resp.Format(http.StatusUnauthorized, errors.New("invalid token")))
			c.Abort()
			return
		}

		metadata, err := jwt.ExtractTokenMetadata(authorization)
		if err != nil {
			c.JSON(resp.Format(http.StatusInternalServerError, err)) //errors.New("unable to extract token metadata")))
			c.Abort()
			return
		}

		// token valid and forward to original request

		c.Set("id", metadata.IDUser)
		c.Set("nip", metadata.NIP)
		c.Set("nama", metadata.Nama)
		c.Set("no_hp", metadata.Nohp)
		c.Set("photo", metadata.Photo)
		c.Set("id_subdirektorat", metadata.IDsubdirektorat)
		c.Set("id_seksi", metadata.IDseksi)
		c.Set("level_pengguna", metadata.Levelpengguna)
		c.Set("aktif", metadata.Aktif)

		c.Next()

		// original request goes here
	}
}
