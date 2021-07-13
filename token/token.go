package token

import (
	//"newjwt/models"
	"time"

	"github.com/Waire214/newjwt/models"
	"github.com/dgrijalva/jwt-go"
)

const (
	jWTPrivateToken = "SecrteTokenSecrteToken"
	ip              = "192.168.0.107"
)

func GenrateTokenString(claims *models.JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
