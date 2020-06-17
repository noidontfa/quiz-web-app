package model

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTCustomClaims struct {
	Username string
	jwt.StandardClaims
}
