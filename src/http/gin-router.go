package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

var (
	categoryController = controller.NewCategoryController()
)

func register(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/category", categoryController.GetCategories)
	}
}
