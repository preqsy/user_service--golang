package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encryptedPassword)
}
