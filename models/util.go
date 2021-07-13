package models

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

//ValidResponse structures valid api response into a json object

func ValidResponse(code int, body interface{}, message string) ResponseObject {
	var response ResponseObject
	response.Code = code
	response.Message = message
	response.Body = body

	return response
}

//HashPassword generates a secure password string
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash decryptes password to see validity
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GenerateRandomString returns a random string of the specified length
func GenerateRandomString(lenght int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, lenght)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
