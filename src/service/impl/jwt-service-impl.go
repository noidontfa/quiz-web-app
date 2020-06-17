package impl

import (
	models "../../model"
	"../../service"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"time"
)

type JWTSevc struct {
	secretKey string
	client *redis.Client
}

func NewJWTService(cl *redis.Client) service.JWTService {
	return &JWTSevc{
		secretKey: "ThinhTest",
		client:cl,
	}
}

func (J *JWTSevc) GenerateToken(username string) string {
	claim := &models.JWTCustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(J.secretKey))
	if err != nil {
		panic(err)
	} else {
		err = J.SaveToken(username,t,claim.StandardClaims.ExpiresAt)
		if err != nil {
			return ""
		}
	}
	return t
}

func (J *JWTSevc) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(J.secretKey), nil
	})
}

func (J *JWTSevc)SaveToken(username string,tokenString string,exp int64) error  {
	at := time.Unix(exp,0)

	now := time.Now()
	ctx := context.Background()
	errAccess := J.client.Set(ctx, username,tokenString,at.Sub(now)).Err()
	return errAccess
}

func (J *JWTSevc) GetRedisToken(username string) string {
	ctx := context.Background()
	token,_ := J.client.Get(ctx,username).Result()
	return token
}