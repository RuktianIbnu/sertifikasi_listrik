package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenMetadata ...
type TokenMetadata struct {
	Id_user    int64
	Username   string
	Nama_admin string
	Id_level   int64
}

// CreateToken ...
func CreateToken(idUser int64, username, nama_admin string, id_level int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["id_user"] = idUser
	claims["username"] = username
	claims["nama_admin"] = nama_admin
	claims["id_level"] = id_level
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyToken(headerToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid ...
func TokenValid(headerToken string) error {
	token, err := verifyToken(headerToken)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(headerToken string) (*TokenMetadata, error) {
	token, err := verifyToken(headerToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id_user, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id_user"]), 10, 64)
		if err != nil {
			// fmt.Println(err.Error())
			return nil, err
		}

		username, ok := claims["username"].(string)
		if !ok {
			return nil, err
		}

		nama_admin, ok := claims["nama_admin"].(string)
		if !ok {
			return nil, err
		}

		id_level, ok := claims["id_level"].(int64)
		if !ok {
			return nil, err
		}

		return &TokenMetadata{
			Id_user:    id_user,
			Username:   username,
			Nama_admin: nama_admin,
			Id_level:   id_level,
		}, nil
	}
	return nil, err
}
