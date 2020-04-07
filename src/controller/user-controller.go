package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	user, err := u.UserService.FindById(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,user)
}

func (u *UserControl) UpdateUser(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	userResult, err := u.UserService.Update(uint(id),&user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, userResult)
}

func (u *UserControl) DeleteUser(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"),0,0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	if err := u.UserService.Delete(uint(id)); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
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

