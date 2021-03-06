package service

import "github.com/dgrijalva/jwt-go"

type JWTService interface {
	GenerateToken(username string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetRedisToken(username string) string
	DeleteRedisToken(username string) error
}
