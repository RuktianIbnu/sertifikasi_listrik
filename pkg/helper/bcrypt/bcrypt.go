package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash ...
func Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 5)
	return string(bytes), err
}

// Compare ...
func Compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// fmt.Println(err.Error())
	return err == nil
}
