package main

import (
	"./config"
	"./controller"
	"./repository"
	service "./service/impl"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	cf = config.NewConfiguration()
)

func main() {
	config, err := cf.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	repo := repository.NewMySqlRepository(config)
	repo.AutoMigration()

	categoryService := service.NewCategoryService(repo)
	categoryController := controller.NewCategoryController(categoryService)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/category",categoryController.GetCategories)
		api.POST("/category",categoryController.PostCategories)
	}

	log.Fatal(router.Run(fmt.Sprintf("%s:%s",config.HttpServerHost,config.Port)))


}

