package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginControl struct {
	LoginService service.LoginService
	JWTService service.JWTService
}



type LoginController interface {
	Login(ctx *gin.Context)
}

func NewLoginController(sevc service.LoginService, jwtSevc service.JWTService) LoginController {
	return &LoginControl{
		LoginService:sevc,
		JWTService:jwtSevc,
	}
}

func (l *LoginControl) Login(ctx *gin.Context) {
	var credentials models.User
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,err.Error())
		return
	}
	isAuthenticated := l.LoginService.Login(credentials.Username,credentials.Password)
	if !isAuthenticated {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	token := l.JWTService.GenerateToken(credentials.Username)
	ctx.JSON(http.StatusOK,gin.H{
		"token": token,
	})
}