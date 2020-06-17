package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QControl struct {
	quizService service.QuizService
}

type QuizController interface {
	FindAllQuizzes(ctx *gin.Context)
	FindByIdQuiz(ctx *gin.Context)
	UpdateQuiz(ctx *gin.Context)
	DeleteQuiz(ctx *gin.Context)
	SaveQuiz(ctx *gin.Context)
	FindByUserId(ctx *gin.Context)
}

func NewQuizController(serv service.QuizService) QuizController {
	return &QControl{quizService: serv}
}

func (q *QControl) FindAllQuizzes(ctx *gin.Context) {
	quizzes, err := q.quizService.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, quizzes)
}

func (q *QControl) FindByIdQuiz(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	quiz, err1 := q.quizService.FindById(uint(id))
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, quiz)
}

func (q *QControl) UpdateQuiz(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var quiz models.Quiz
	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	quizResult, err := q.quizService.Update(uint(id), &quiz)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, quizResult)

}

func (q *QControl) DeleteQuiz(ctx *gin.Context) {
	id, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusBadRequest, err1.Error())
		return
	}
	err := q.quizService.Delete(uint(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
}

func (q *QControl) SaveQuiz(ctx *gin.Context) {
	var quiz models.Quiz
	if err := ctx.ShouldBindJSON(&quiz); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	quizResult, err := q.quizService.Save(&quiz)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, quizResult)
}

func (q *QControl) FindByUserId(ctx *gin.Context) {
	quizId, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	quizzes, err1 := q.quizService.FindByUserId(uint(quizId))
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, quizzes)
}
