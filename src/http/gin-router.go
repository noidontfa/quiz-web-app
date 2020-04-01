package router

import (
	"github.com/gin-gonic/gin"
	"../controller"
)

var (
	categoryController = controller.NewCategoryController()
)


func register(router *gin.Engine)  {
	api := router.Group("/api")
	{
		api.GET("/category",categoryController.GetCategories)
	}
}