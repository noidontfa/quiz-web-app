package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleControl struct {
	RoleService service.RoleService
}

type RoleController interface {
	FindAllRoles(ctx *gin.Context)
	FindByIdRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	SaveRole(ctx *gin.Context)
}

func NewRoleController(sevc service.RoleService) RoleController {
	return &RoleControl{RoleService:sevc}
}

func (r *RoleControl) FindAllRoles(ctx *gin.Context) {
	roles, err := r.RoleService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,roles)
}

func (r *RoleControl) FindByIdRole(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	role, err1 := r.RoleService.FindById(uint(id))
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	ctx.JSON(http.StatusOK,role)
}

func (r *RoleControl) UpdateRole(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	roleResult, err1 := r.RoleService.Update(uint(id),&role)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	ctx.JSON(http.StatusOK,roleResult)
}

func (r *RoleControl) DeleteRole(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"),0,0)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	err1 := r.RoleService.Delete(uint(id))
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	ctx.String(http.StatusOK,"Deleted")
}

func (r *RoleControl) SaveRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	roleResult, err := r.RoleService.Save(&role)
	if err != nil {
		ctx.String(http.StatusInternalServerError,err.Error())
		return
	}
	ctx.JSON(http.StatusOK,roleResult)
}



