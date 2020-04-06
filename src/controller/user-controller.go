package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserControl struct {
	UserService service.UserService
}

type UserController interface {
	FindAllUsers(ctx *gin.Context)
	FindByIdUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
}

func NewUserController(serv service.UserService) UserController  {
	return &UserControl{UserService:serv}
}

func (u *UserControl) FindAllUsers(ctx *gin.Context) {
	users, err := u.UserService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,users)
}

func (u *UserControl) FindByIdUser(ctx *gin.Context) {
	panic("implement me")
}

func (u *UserControl) UpdateUser(ctx *gin.Context) {
	panic("implement me")
}

func (u *UserControl) DeleteUser(ctx *gin.Context) {
	panic("implement me")
}

func (u *UserControl) SaveUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	userResult, err := u.UserService.Save(&user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, userResult)
}

