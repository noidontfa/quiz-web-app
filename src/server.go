package main

import (
	"./config"
	"./controller"
	"./middlewares"
	"./model"
	"./repository"
	service "./service/impl"
	"./utils"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"os"
)

var (
	cf = config.NewConfiguration()
	client *redis.Client
)

func upload(c *gin.Context) {
	//file, err := c.FormFile("file")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = c.SaveUploadedFile(file,"./src/public/" + file.Filename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//filepath := "http://localhost:8080/file/" + file.Filename

	var image model.Image
	if err := c.BindJSON(&image); err != nil {
		log.Fatal(err)
	}
	randFileName, _ := utils.Random_filename_16_char()
	randFileName += "-" + image.Filename
	dec, err := base64.StdEncoding.DecodeString(image.DataBase64)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("./src/public/" + randFileName)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(dec); err != nil {
		log.Fatal(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"filepath": "/file/" + randFileName})
}

func main() {
	config, err := cf.GetConfig()
	if err != nil {
		panic(err.Error())
	}
	repo := repository.NewMySqlRepository(config)
	repo.AutoMigration()
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	categoryService := service.NewCategoryService(repo)
	quizService := service.NewQuizService(repo)
	languageService := service.NewLanguageService(repo)
	timingService := service.NewTimingService(repo)
	roleService := service.NewRoleService(repo)
	userService := service.NewUserService(repo)
	questionService := service.NewQuestionService(repo)
	ratingService := service.NewRatingService(repo)
	historyService := service.NewHistoryService(repo)
	stateService := service.NewStateService(repo)
	choiceService := service.NewChoiceService(repo)
	loginService := service.NewLoginService(repo)
	jwtService := service.NewJWTService(client)
	categoryController := controller.NewCategoryController(categoryService)
	quizController := controller.NewQuizController(quizService)
	languageController := controller.NewLanguageController(languageService)
	timingController := controller.NewTimingController(timingService)
	roleController := controller.NewRoleController(roleService)
	userController := controller.NewUserController(userService)
	questionController := controller.NewQuestionController(questionService)
	ratingController := controller.NewRatingController(ratingService)
	historyController := controller.NewHistoryController(historyService)
	stateController := controller.NewStateController(stateService)
	choiceController := controller.NewChoiceController(choiceService)
	loginController := controller.NewLoginController(loginService, jwtService,userService)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("./src/template/*")
	router.GET("/select", func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	})
	router.POST("/upload", upload)
	router.StaticFS("/file", http.Dir("./src/public"))

	router.POST("/login", loginController.Login)
	router.POST("/logout", loginController.Logout)


	testAuthor := router.Group("/test", middlewares.AuthorizeJWT(client))
	{
		testAuthor.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Value": "Hello world",
			})
		})
	}

	api := router.Group("/api")
	{
		api.GET("/categories", categoryController.FindAllCategories)
		api.GET("/categories/:id", categoryController.FindByIdCategory)
		api.PUT("/categories/:id", categoryController.UpdateCategory)
		api.DELETE("/categories/:id", categoryController.DeleteCategory)
		api.POST("/categories", categoryController.SaveCategory)

		api.GET("/quizzes", quizController.FindAllQuizzes)
		api.GET("/quizzes/:id/my", quizController.FindByUserId)
		api.GET("/quizzes/:id", quizController.FindByIdQuiz)
		api.PUT("/quizzes/:id", quizController.UpdateQuiz)
		api.DELETE("/quizzes/:id", quizController.DeleteQuiz)
		api.POST("/quizzes", quizController.SaveQuiz)

		api.GET("/languages", languageController.FindAllLanguages)
		api.GET("/languages/:id", languageController.FindByIdLanguage)
		api.PUT("/languages/:id", languageController.UpdateLanguage)
		api.DELETE("/languages/:id", languageController.DeleteLanguage)
		api.POST("/languages", languageController.SaveLanguage)

		api.GET("/timings", timingController.FindAllTimings)
		api.GET("/timings/:id", timingController.FindByIdTiming)
		api.PUT("/timings/:id", timingController.UpdateTiming)
		api.DELETE("/timings/:id", timingController.DeleteTiming)
		api.POST("/timings", timingController.SaveTiming)

		api.GET("/roles", roleController.FindAllRoles)
		api.GET("/roles/:id", roleController.FindByIdRole)
		api.PUT("/roles/:id", roleController.UpdateRole)
		api.DELETE("/roles/:id", roleController.DeleteRole)
		api.POST("/roles", roleController.SaveRole)

		api.GET("/users", userController.FindAllUsers)
		api.GET("/users/:id", userController.FindByIdUser)
		api.PUT("/users/:id", userController.UpdateUser)
		api.GET("/user/info", userController.GetUserInfo)
		api.DELETE("/users/:id", userController.DeleteUser)
		api.POST("/users", userController.SaveUser)

		api.POST("/questions/:id", questionController.SaveQuestions)
		api.DELETE("/questions/:id", questionController.DeleteQuestions)

		api.DELETE("/choices/:id", choiceController.DeleteChoices)

		api.POST("/ratings/", ratingController.SaveRating)

		api.GET("/histories/", historyController.FindByIdHistory)
		api.GET("/histories/d/", historyController.FindByDateHistory)
		api.POST("/histories/", historyController.SaveHistory)

		api.GET("/states", stateController.FindAllState)
	}

	log.Fatal(router.Run(fmt.Sprintf("%s:%s", config.HttpServerHost, config.Port)))

}
