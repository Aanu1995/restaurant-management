package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// verify password
func VerifyPassword(userPassword string, providedPassword string) (passwordIsValid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	passwordIsValid = err == nil
	return
}

// HashPassword
func Hashpassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(result)
}