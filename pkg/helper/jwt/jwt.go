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
	IDUser          int64
	NIP             string
	Nama            string
	Nohp            string
	IDsubdirektorat int64
	IDseksi         int64
	Aktif           bool
	Levelpengguna   int64
	Photo           string
}

// CreateToken ...
func CreateToken(idUser int64, nip, nama, nohp, photo string, idSubdirektorat, idSeksi, level_pengguna int64, aktif bool) (string, error) {
	claims := jwt.MapClaims{}
	claims["id_user"] = idUser
	claims["nip"] = nip
	claims["nama"] = nama
	claims["no_hp"] = nohp
	claims["photo"] = photo
	claims["id_subdirektorat"] = idSubdirektorat
	claims["id_seksi"] = idSeksi
	claims["aktif"] = aktif
	claims["level_pengguna"] = level_pengguna
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
		idUser, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id_user"]), 10, 64)
		if err != nil {
			return nil, err
		}

		nip, ok := claims["nip"].(string)
		if !ok {
			return nil, err
		}

		nama, ok := claims["nama"].(string)
		if !ok {
			return nil, err
		}

		nohp, ok := claims["no_hp"].(string)
		if !ok {
			return nil, err
		}

		photo, ok := claims["photo"].(string)
		if !ok {
			return nil, err
		}

		idsubdirektorat, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id_subdirektorat"]), 10, 64)
		if err != nil {
			return nil, err
		}

		idseksi, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["id_seksi"]), 10, 64)
		if err != nil {
			return nil, err
		}

		aktif, err := strconv.ParseBool(fmt.Sprintf("%.t", claims["aktif"]))
		if err != nil {
			return nil, err
		}

		level_pengguna, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["level_pengguna"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &TokenMetadata{
			IDUser:          idUser,
			NIP:             nip,
			Nama:            nama,
			Nohp:            nohp,
			Photo:           photo,
			IDsubdirektorat: idsubdirektorat,
			IDseksi:         idseksi,
			Aktif:           aktif,
			Levelpengguna:   level_pengguna,
		}, nil
	}
	return nil, err
}
