package controller

import (
	models "../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QuestionControl struct {
	QuestionService service.QuestionService
}

type QuestionController interface {
	DeleteQuestions(ctx *gin.Context)
	SaveQuestions(ctx *gin.Context)
}

func NewQuestionController(serv service.QuestionService) QuestionController {
	return &QuestionControl{QuestionService: serv}
}

func (q *QuestionControl) DeleteQuestions(ctx *gin.Context) {
	questionId, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	err := q.QuestionService.Delete(uint(questionId))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.String(http.StatusOK, "Deleted")
}

func (q *QuestionControl) SaveQuestions(ctx *gin.Context) {
	quizId, err1 := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err1 != nil {
		ctx.String(http.StatusInternalServerError, err1.Error())
		return
	}
	var questions []models.Question
	if err := ctx.ShouldBindJSON(&questions); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	questionResult, err := q.QuestionService.Save(uint(quizId), questions)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, questionResult)
}
