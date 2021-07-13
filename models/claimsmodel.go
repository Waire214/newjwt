package models

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Create claims model
type JwtClaims struct {
	Fullname string `json:"fullname,omitempty"`
	Email    string `json:"email,omitempty"`
	Roles    string `json:"roles,omitempty"`
	jwt.StandardClaims
}

const ip = "192.168.0.107"

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	//pass in the time(now), do you really want to validate the claims?
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil //nil = no error
	}
	return fmt.Errorf("token is invalid")
}
func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
