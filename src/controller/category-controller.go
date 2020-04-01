package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type Control struct {
	CategoryService service.CategoryService
}


type CategoryController interface {
	GetCategories(ctx *gin.Context)
	PostCategories(ctx *gin.Context)
}

func NewCategoryController(serv service.CategoryService) CategoryController {
	return &Control{
		CategoryService: serv,
	}
}

func (c *Control) GetCategories(ctx *gin.Context) {
	c.CategoryService.FindAll()
}

func (c *Control) PostCategories(ctx *gin.Context) {
	var category models.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		log.Fatal(err.Error())
		ctx.String(http.StatusNoContent,err.Error())
	}
	ctx.JSON(http.StatusOK,c.CategoryService.Save(category))
}


