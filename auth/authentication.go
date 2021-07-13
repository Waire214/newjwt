package auth

import (
	"time"

	"github.com/Waire214/newjwt/misc"
	"github.com/Waire214/newjwt/models"
	"github.com/Waire214/newjwt/token"
)

//HandleUserRegistration handles user registration for single users
func HandleUserRegistration(regData models.RegistrationData) models.ResponseObject {
	conn := misc.GetDB()
	var userData models.RegistrationData
	conn.Where("email = ?", regData.Email).First(&userData)
	if userData.Email != "" {
		return models.ValidResponse(400, regData, "User Already exists")
	}
	userData.Email = regData.Email
	userData.FullName = regData.FullName
	// userData.Status = "pending"
	hashPassword, err := models.HashPassword(regData.Password)
	misc.CheckErr(err, true, "Error encrypting password")
	userData.Password = hashPassword

	if err := conn.Create(&regData).Error; err != nil {
		misc.LogError(err)
		return models.ValidResponse(400, regData, "Unable to register user")
	}

	return models.ValidResponse(200, userData, "Success")
}
func HandleUserLogin(authdetails models.LoginData) models.ResponseObject {
	conn := misc.GetDB()
	var authuser models.RegistrationData
	conn.Where("email = ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		return models.ValidResponse(400, authuser, "Username or Password is incorrect")
	}
	passwordMatch := models.CheckPasswordHash(authdetails.Password, authuser.Password)
	if !passwordMatch {
		return models.ValidResponse(400, authuser, "Username or Password is incorrect")
	}
	var claims = &models.JwtClaims{}
	claims.Fullname = authuser.FullName
	claims.Roles = authuser.Role
	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	tokenString, err := token.GenrateTokenString(claims, expirationTime)
	if err != nil {
		models.ValidResponse(400, tokenString, "error in generating token")
	}

	return models.ValidResponse(200, tokenString, "Success")

}
