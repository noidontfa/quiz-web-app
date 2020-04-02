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

	categoryService 	:= service.NewCategoryService(repo)
	quizService 		:= service.NewQuizService(repo)
	categoryController 	:= controller.NewCategoryController(categoryService)
	quizController		:= controller.NewQuizController(quizService)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/categories",categoryController.FindAllCategories)
		api.GET("/categories/:id",categoryController.FindByIdCategory)
		api.PUT("/categories/:id",categoryController.UpdateCategory)
		api.DELETE("/categories/:id",categoryController.DeleteCategory)
		api.POST("/categories",categoryController.SaveCategory)

		api.GET("/quizzes",quizController.FindAllQuizzes)
		api.GET("/quizzes/:id",quizController.FindByIdQuiz)
		api.PUT("/quizzes/:id",quizController.UpdateQuiz)
		api.DELETE("/quizzes/:id",quizController.DeleteQuiz)
		api.POST("/quizzes",quizController.SaveQuiz)
	}

	log.Fatal(router.Run(fmt.Sprintf("%s:%s",config.HttpServerHost,config.Port)))


}

