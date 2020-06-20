package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginControl struct {
	LoginService 	service.LoginService
	JWTService   	service.JWTService
	UserService		service.UserService
}

type LoginController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewLoginController(sevc service.LoginService, jwtSevc service.JWTService, userSevc service.UserService) LoginController {
	return &LoginControl{
		LoginService: sevc,
		JWTService:   jwtSevc,
		UserService:	userSevc,
	}
}

func (l *LoginControl) Login(ctx *gin.Context) {
	var credentials models.User
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	isAuthenticated := l.LoginService.Login(credentials.Username, credentials.Password)
	if !isAuthenticated {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	token := l.JWTService.GetRedisToken(credentials.Username)
	if token == "" {
		token = l.JWTService.GenerateToken(credentials.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (l *LoginControl) Logout(ctx *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]

	user, err := l.UserService.FindByToken(tokenString)
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}

	err = l.JWTService.DeleteRedisToken(user.Username)
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.String(http.StatusOK,"Done")
}



