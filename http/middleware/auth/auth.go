package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"sertifikasi_listrik/pkg/helper/jwt"
	resp "sertifikasi_listrik/pkg/helper/response"

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

		fmt.Println(metadata)
		// token valid and forward to original request

		// c.Set("id_user", metadata.Id_user)
		// c.Set("username", metadata.Username)
		// c.Set("nama_admin", metadata.Nama_admin)
		// c.Set("id_level", metadata.Id_level)

		c.Next()

		// original request goes here
	}
}
