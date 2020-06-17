package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Control struct {
	CategoryService service.CategoryService
}

type CategoryController interface {
	FindAllCategories(ctx *gin.Context)
	FindByIdCategory(ctx *gin.Context)
	UpdateCategory(ctx *gin.Context)
	DeleteCategory(ctx *gin.Context)
	SaveCategory(ctx *gin.Context)
}

func NewCategoryController(serv service.CategoryService) CategoryController {
	return &Control{
		CategoryService: serv,
	}
}

func (c *Control) FindAllCategories(ctx *gin.Context) {
	categories, err := c.CategoryService.FindAll()
	if err != nil {
		//log.Fatal(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *Control) FindByIdCategory(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		//log.Fatal(err1.Error())
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	category, err := c.CategoryService.FindById(uint(id))
	if err != nil {
		//log.Fatal(err.Error())
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *Control) UpdateCategory(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.String(http.StatusNoContent, err.Error())
		return
	}

	categoryResult, err := c.CategoryService.Update(uint(id), &category)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categoryResult)
}

func (c *Control) DeleteCategory(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	err := c.CategoryService.Delete(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
}

func (c *Control) SaveCategory(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	categoryResult, err := c.CategoryService.Save(&category)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, categoryResult)
}
