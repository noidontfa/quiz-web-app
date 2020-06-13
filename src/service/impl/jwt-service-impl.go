package impl

import (
	"../../service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTSevc struct {
	secretKey string
}

type jwtCustomClaims struct {
	username string
	jwt.StandardClaims
}


func NewJWTService () service.JWTService {
	return &JWTSevc{secretKey: "ThinhTest"}
}

func (J *JWTSevc) GenerateToken(username string) string {
	claim := &jwtCustomClaims{
		username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer: username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(J.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (J *JWTSevc) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(J.secretKey),nil
	})
}




